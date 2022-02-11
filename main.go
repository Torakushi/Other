package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"main/robot_handler_other"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// check for interuption
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

		<-ch

		fmt.Println("OMG WE RECEIVED A DEMAND TO STOP !!!!!")
		// Stop all process
		cancel()
	}()

	rh := robot_handler_other.NewRobotHandler(10) // Sem
	//rh := robot_handler.NewRobotHandler(10) // for i ...

	// Server Part
	http.HandleFunc("/mean", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mean is: %f", rh.GetMean())
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	rh.Process(ctx)
	fmt.Println("Finish the process ... Bye !!!")
}

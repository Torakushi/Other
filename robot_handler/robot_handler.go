package robot_handler

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RobotHandler struct {
	ch     chan int
	robots []robot
}

func NewRobotHandler(maxGoroutine int) *RobotHandler {
	return &RobotHandler{
		ch:     make(chan int, 5000),
		robots: make([]robot, maxGoroutine),
	}
}

type robot struct {
	count int
	total int
}

func (rh *RobotHandler) Process(ctx context.Context) {
	fmt.Println("Launch process ! ")
	// Send to channel
	go rh.produce(ctx)

	var wg sync.WaitGroup
	for i := 0; i < len(rh.robots); i++ {
		wg.Add(1)
		go func(r *robot, i int) {
			defer wg.Done()
			for v := range rh.ch {
				r.count++
				r.total += v
			}
			fmt.Printf("robot %d has count= %d and total= %d\n", i, r.count, r.total)
		}(&rh.robots[i], i)
	}
	wg.Wait()
}

func (rh *RobotHandler) produce(ctx context.Context) {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Sending in channel !")
			for i := 0; i < 5000; i++ {
				rh.ch <- rand.Intn(200)
			}
		case <-ctx.Done():
			close(rh.ch)
			fmt.Println("Stop produce as user asked to stop!")
			return
		}

	}
}

func (rh *RobotHandler) GetMean() float64 {
	var count, total int
	for _, r := range rh.robots {
		count += r.count
		total += r.total
	}

	if count == 0 {
		return 0
	}

	return float64(total) / float64(count)
}

package robot_handler_other

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

type RobotHandler struct {
	ch            chan int
	sem           *semaphore.Weighted
	total         int64
	count         int64
	maxGoroutines int64
}

func NewRobotHandler(maxGoroutines int64) *RobotHandler {
	return &RobotHandler{
		ch:            make(chan int, 5000),
		sem:           semaphore.NewWeighted(maxGoroutines),
		maxGoroutines: maxGoroutines,
	}
}

func (rh *RobotHandler) Process(ctx context.Context) {
	fmt.Println("Launch process ! ")
	// Send to channel
	go rh.produce(ctx)

	var id int
	for rh.isRunning(ctx) {
		rh.sem.Acquire(ctx, 1)
		fmt.Printf("Launch goroutine %d\n", id)
		id++

		go func() {
			defer rh.sem.Release(1)
			for v := range rh.ch {
				atomic.AddInt64(&rh.count, 1)
				atomic.AddInt64(&rh.total, int64(v))
			}
		}()
	}

	rh.sem.Acquire(ctx, rh.maxGoroutines)
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

func (rh *RobotHandler) isRunning(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		return true
	}
}

func (rh *RobotHandler) GetMean() float64 {
	count := float64(atomic.LoadInt64(&rh.count))
	if count == 0 {
		return 0
	}
	return float64(atomic.LoadInt64(&rh.total)) / count
}

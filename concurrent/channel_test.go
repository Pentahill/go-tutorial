package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func getNameTask(idx int) string {
	time.Sleep(time.Microsecond * 200)
	return fmt.Sprintf("get name from %d ", idx)
}

func getName() string {
	numberWorker := 10
	ret := make(chan string, numberWorker)
	for i := 1; i < numberWorker; i++ {
		go func(idx int) {
			ret <- getNameTask(idx)
		}(i)
	}

	return <-ret
}

func getAllName() string {
	numberWorker := 10
	ret := make(chan string, numberWorker)
	for i := 1; i < numberWorker; i++ {
		go func(idx int) {
			ret <- getNameTask(idx)
		}(i)
	}

	var nameString string
	for i := 1; i < numberWorker; i++ {
		nameString += <-ret
	}

	return nameString
}

func TestGetName(t *testing.T) {
	fmt.Printf("g number is %d\n", runtime.NumGoroutine())
	// ret := getName()
	ret := getAllName()
	fmt.Printf("result is %s\n", ret)

	time.Sleep(time.Second * 4)
	fmt.Printf("g number is %d\n", runtime.NumGoroutine())
}

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i < 5; i++ {
		go func(idx int, ctx context.Context) {
			defer func() {
				fmt.Printf("************** task %d is finished ****************\n", idx)
			}()

			for {
				if isCanceled(ctx) {
					fmt.Printf("task %d is canceled\n", idx)
					break
				}

				fmt.Println("task is running")
			}

		}(i, ctx)
	}

	cancel()
	time.Sleep(time.Second * 2)
}

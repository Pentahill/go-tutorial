package main

import (
	"fmt"
	"testing"
	"time"
)

func blockedService() chan int {
	ret := make(chan int)

	go func() {
		fmt.Println("run service")
		ret <- 2
		time.Sleep(time.Second * 2)
		fmt.Println("end service")
	}()

	fmt.Println("return channel")

	return ret
}

func TestChannel(t *testing.T) {
	t.Log(<-blockedService())
	t.Log("end test")

	time.Sleep(time.Second * 5)
}

func isCanceled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})

	for i := 1; i < 5; i++ {
		go func(idx int, ch chan struct{}) {
			defer func() {
				fmt.Println("task is finished")
			}()

			for {
				if isCanceled(cancelChan) {
					fmt.Printf("task %d is canceled\n", idx)
					break
				}

				fmt.Println("task is running")
			}

		}(i, cancelChan)
	}

	time.Sleep(time.Millisecond * 200)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 2)
}

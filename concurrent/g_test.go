package main

import (
	"sync"
	"testing"
	"time"
)

func TestG(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			t.Log(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	time.Sleep(time.Second * 2)
}

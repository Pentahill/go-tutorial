package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{New: func() interface{} {
		fmt.Println("create a new ")
		return 100
	}}

	v := pool.Get().(int)
	fmt.Println(v)
	v1 := pool.Get().(int)
	fmt.Println(v1)
}

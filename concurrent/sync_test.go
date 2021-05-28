package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type SingleInstance struct {
}

var singleInstance *SingleInstance
var once sync.Once

func getInstane() *SingleInstance {
	once.Do(func() {
		fmt.Println("create instance")
		singleInstance = new(SingleInstance)
	})

	return singleInstance
}

func TestOnce(t *testing.T) {
	for i := 1; i < 10; i++ {
		instacne := getInstane()
		fmt.Println(unsafe.Pointer(instacne))
	}
}

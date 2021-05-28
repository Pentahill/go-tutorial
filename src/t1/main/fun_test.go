package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(fun func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		re := fun(op)
		fmt.Println("time consume: ", time.Since(start).Seconds())

		return re
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)

	return op
}

func TestFunc(t *testing.T) {
	a, _ := returnMultiValue()
	t.Log(a)

	timeSpentFunc := timeSpent(slowFun)
	t.Log(timeSpentFunc(10))
}

func Clear() {
	fmt.Println("Clear resource")
}

func TestDefer(t *testing.T) {
	defer Clear()

	fmt.Println("start")
	panic("err")
}

package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("ou")
		case i%2 == 1:
			t.Log("Odd")
		}
	}
}

func TestArray(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}

func TestSlice(t *testing.T) {
	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := make([]int, 3, 4)
	t.Log(len(s1), cap(s1))

	s1 = append(s1, 5)
	t.Log(s1[3])

	str_slice := []string{"1", "2", "3", "4", "5", "6"}
	x0 := str_slice[3:6]
	t.Log(x0, len(str_slice), cap(str_slice), len(x0), cap(x0))
}

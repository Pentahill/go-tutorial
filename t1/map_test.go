package main

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 4, 2: 6}
	t.Log(m1[1])

	m2 := map[int]int{}
	t.Log(len(m2))

	if v, ok := m2[10]; ok {
		t.Log(v, ok)
	} else {
		t.Log("not exist")
	}

	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapFunc(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	m1[4] = func(op int) int { return op * 2 }

	t.Log(m1[1](2))
}

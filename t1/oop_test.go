package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employ struct {
	id   int
	age  int
	name string
}

func (e Employ) String() string {
	fmt.Println("name address", unsafe.Pointer(&e.name))
	return fmt.Sprintf("employ id %x, age %x, name %s", e.id, e.age, e.name)
}

func TestOOP(t *testing.T) {
	e := Employ{id: 1, age: 18, name: "liu"}
	t.Log("name address", unsafe.Pointer(&e.name))
	t.Log(e.String())
}

func emptyInterfaceHandler(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	}
}

func TestEmpytInterface(t *testing.T) {
	emptyInterfaceHandler(1)
}

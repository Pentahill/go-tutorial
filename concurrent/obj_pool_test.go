package main

import (
	"errors"
	"testing"
	"time"
)

type ReusableObject struct {
}

type ObjectPool struct {
	bufChan chan *ReusableObject
}

func NewObjectPool(numOfObj int) *ObjectPool {
	objectPool := ObjectPool{}
	objectPool.bufChan = make(chan *ReusableObject, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objectPool.bufChan <- &ReusableObject{}
	}

	return &objectPool
}

func (objPool *ObjectPool) GetObject(timeout time.Duration) (*ReusableObject, error) {
	select {
	case obj := <-objPool.bufChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (objPool *ObjectPool) ReleaseObject(obj *ReusableObject) error {
	select {
	case objPool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	objectPool := NewObjectPool(10)
	for i := 0; i < 11; i++ {
		if _, error := objectPool.GetObject(time.Second * 2); error != nil {
			t.Log("no enough object")
		} else {
			t.Log("get object ", i)
		}
	}

}

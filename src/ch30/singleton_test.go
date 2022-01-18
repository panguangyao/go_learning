package ch30

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonOjb() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(i)
		go func() {
			obj := GetSingletonOjb()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			//fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

package ch23

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine01(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}

func TestGoRoutine02(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

package ch29

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel01(cancelChan chan struct{}) {
	cancelChan <- struct{}{} //cancel one chan
}

func cancel02(cancelChan chan struct{}) {
	close(cancelChan) //broadcast cancel all chan
}

func TestCancelChan(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled.")
		}(i, cancelChan)
	}
	//cancel01(cancelChan)
	cancel02(cancelChan)
	time.Sleep(time.Second * 1)
}

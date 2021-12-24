package ch27

import (
	"fmt"
	"sync"
	"testing"
)

//chan close

func Source(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i := 0; i < 10; i++ {
		//	ch <- i
		//}
		//wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//ch <- 11 panic chan is closed
		wg.Done()
	}()
}

func Sink(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i := 0; i < 11; i++ {
		//	data := <-ch
		//	fmt.Println(data)
		//}
		//wg.Done()
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChan(t *testing.T) {
	//var wg *sync.WaitGroup wg maybe nil/NPE
	//var wg sync.WaitGroup
	//ch := make(chan int)
	//wg.Add(1)
	//Source(ch, &wg)
	//wg.Add(1)
	//Sink(ch, &wg)
	//wg.Wait()
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	Source(ch, &wg)
	wg.Add(1)
	Sink(ch, &wg)
	wg.Add(1)
	Sink(ch, &wg)
	wg.Wait()
}

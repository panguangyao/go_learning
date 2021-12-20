package ch13

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
1、函数可以有多个返回值
2、所有参数都是值传递，slice、map、channel会有传引用的错觉
3、函数可以作为变量的值
4、函数可以作为参数和返回值
*/

func TestFunc(t *testing.T) {
	a, b := returnMultipleValues()
	t.Log(a, b)
}

func returnMultipleValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//函数一等公民
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc02(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

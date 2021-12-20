package ch16

import (
	"fmt"
	"testing"
	"time"
)

//接口变量？

//自定义类型
type IntConv func(op int) int

func TestCustomType(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func timeSpent(inner IntConv) IntConv {
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

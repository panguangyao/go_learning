package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	t.Log("a:", a, ",b:", b)
}

func TestExchange2(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log("a:", a, ",b:", b)
}

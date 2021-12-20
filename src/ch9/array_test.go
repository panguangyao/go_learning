package ch9

import "testing"

func TestArray(t *testing.T) {
	var a [3]int
	t.Log(a)
	a[0] = 1
	t.Log(a)
	arr1 := [2]int{1, 2}
	arr2 := [...]int{3, 4}
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for i, v := range arr {
		t.Log(i, v)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[:3]
	t.Log(arrSec)
	t.Log(arr[3:])
}

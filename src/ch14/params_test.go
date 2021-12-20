package ch14

import "testing"

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4, 5, 6))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

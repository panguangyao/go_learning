package loop

import "testing"

func TestWhile(t *testing.T) {
	n := 0

	for n < 5 {
		t.Log(n)
		n++
	}
}

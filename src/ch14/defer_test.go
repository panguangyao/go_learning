package ch14

import "testing"

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources.")
	}()
	t.Log("started")
	panic("Fatal error")
}

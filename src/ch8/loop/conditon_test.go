package loop

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	if v,e = someFunc(); e == nil {

	}

}

func someFunc() bool {

	return true
}

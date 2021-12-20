package ch12

import (
	"strconv"
	"strings"
	"testing"
)

//string是不可变的byte切片
func TestStringInit(t *testing.T) {
	var s string
	t.Log(s)
	s = "Hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(c, len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for i, v := range parts {
		t.Log(i, v)
	}

	t.Log(strings.Join(parts, "-"))

}

func TestStringConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}

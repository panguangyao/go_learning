package ch16

import (
	"testing"
)

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

//接口定义
type Programmer interface {
	WriteHelloWorld() string
}

//接口实现
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
}

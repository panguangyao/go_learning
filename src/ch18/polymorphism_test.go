package ch18

import "testing"

//接口定义
type Programmer interface {
	WriteHelloWorld() string
}

//接口实现
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
}

//接口实现
type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.println(\"Hello World!\")"
}

func TestPolymorphism(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
}

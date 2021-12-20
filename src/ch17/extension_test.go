package ch17

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Println("Wang!")
}

func (d *Dog) SpeakTo(host string) {
	//d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Go")
}

package ch18

import (
	"fmt"
	"testing"
)

func DoSth(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if i, ok := p.(string); ok {
	//	fmt.Println("string", i)
	//	return
	//}
	//
	//fmt.Println("Unknown Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown Type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSth(10)
	DoSth("10")
}

package ch15

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Log(e, e1, e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

//行为（方法）定义
//第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String01() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String02() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	e1 := Employee{"0", "Bob", 20}
	t.Log(e.String01())
	t.Log(e1.String01())
	t.Log(e.String02())
	t.Log(e1.String02())
}

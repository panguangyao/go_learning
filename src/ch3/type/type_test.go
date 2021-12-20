package _type

import "testing"

//数据类型不支持隐式转换 别名一样不支持 需要显示转换
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(a)
	t.Log(a, b, c)
}

//指针不支持运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//指针不支持运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	t.Log(s == "")
}

package try

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func Test01(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func Test02(t *testing.T) {
	t.Log(Readable, Writable, Executable)
}

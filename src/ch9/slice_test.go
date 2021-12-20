package ch9

import "testing"

//array可比较 slice不可以
//array定长 slice可扩容
//slice array []没有指定长度
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))

}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 8; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
	t.Log("\n")
	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		t.Log(len(s1), cap(s1))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(year, len(year), cap(year))

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2), year)

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer), Q2, year)

	summer[0] = "Unknown"
	t.Log(summer, len(summer), cap(summer), Q2, year)
}

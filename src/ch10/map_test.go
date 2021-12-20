package ch10

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1, len(m1))

	m2 := map[string]int{}
	t.Log(m2)
	m2["one"] = 1
	t.Log(m2)

	m3 := make(map[string]int, 3) //init capacity
	t.Log(m3, len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {

	m1 := map[int]int{}
	t.Log(m1[0], m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("key 3 is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m1 {
		t.Log(k, v)
	}
	m2 := map[string]int{}
	for k, v := range m2 {
		t.Log("----------")
		t.Log(k, v)
	}

}

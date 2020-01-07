package ch4

import "testing"

func TestConnectedComponents(t *testing.T) {
	adj := [][]int{
		[]int{5,1,2,6},
		[]int{0},
		[]int{0},
		[]int{5,4},
		[]int{5,3,6},
		[]int{0,3,4},
		[]int{0,4},
		[]int{8},
		[]int{7},
		[]int{10,11,12},
		[]int{9},
		[]int{9,12},
		[]int{11,9},
	}

	cc := NewConnectedComponents(adj)
	if !cc.Connected(1, 4) {
		t.Fatalf("1 and 4 should be connected")
	}
	if cc.Connected(2, 12) {
		t.Fatalf("2 and 12 should not be connected")
	}
	if cc.Count() != 3 {
		t.Fatalf("there should be 3 connected components")
	}
}
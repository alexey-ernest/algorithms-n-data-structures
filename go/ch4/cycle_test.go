package ch4

import "testing"

func TestCyclic(t *testing.T) {
	adj := [][]int{
		[]int{2,1,5},
		[]int{0,2},
		[]int{0,1,3,4},
		[]int{2,4,5},
		[]int{2,3},
		[]int{0,3},
	}

	if IsAcyclic(adj) {
		t.Fatalf("the graph is cyclic")
	}
}

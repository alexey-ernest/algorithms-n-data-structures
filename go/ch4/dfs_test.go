package ch4

import "testing"

func TestDFSSingleSource(t *testing.T) {
	adj := [][]int{
		[]int{2,1,5},
		[]int{0,2},
		[]int{0,1,3,4},
		[]int{2,4,5},
		[]int{2,3},
		[]int{0,3},
	}

	dfs := NewDepthFirstPaths(adj, 0)
	if dfs.HasPathTo(5) != true {
		t.Fatalf("it should has a path from 0 to 5")
	}

	path := dfs.PathTo(5)
	exp := []int{0,2,3,5}
	for i := range path {
		if path[i] != exp[i] {
			t.Fatalf("wrong path to 5: %v", path)
		}
	}
}
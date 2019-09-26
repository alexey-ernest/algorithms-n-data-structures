package ch1

import (
	//"fmt"
	"testing"
)

func TestUnionFindQuickFind(t *testing.T) {
	uf := NewUnionFindQuickFind(10)
	input := [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
		{7, 6}, // already connected
	}

	for _, pair := range input {
		p, q := pair[0], pair[1]
		if uf.connected(p, q) {
			continue
		}

		uf.union(p, q)
		//fmt.Printf("%d %d\n", p, q)
	}

	if uf.count() != 2 {
		t.Errorf("number of components %d != 2", uf.count())
	}
}

func TestUnionFindQuickUnion(t *testing.T) {
	uf := NewUnionFindQuickUnion(10)
	input := [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
		{7, 6}, // already connected
	}

	for _, pair := range input {
		p, q := pair[0], pair[1]
		if uf.connected(p, q) {
			continue
		}

		uf.union(p, q)
		//fmt.Printf("%d %d\n", p, q)
	}

	if uf.count() != 2 {
		t.Errorf("number of components %d != 2", uf.count())
	}
}

func TestUnionFindWeightedQuickUnion(t *testing.T) {
	uf := NewUnionFindWeightedQuickUnion(10)
	input := [][2]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
		{7, 6}, // already connected
	}

	for _, pair := range input {
		p, q := pair[0], pair[1]
		if uf.connected(p, q) {
			continue
		}

		uf.union(p, q)
		//fmt.Printf("%d %d\n", p, q)
	}

	if uf.count() != 2 {
		t.Errorf("number of components %d != 2", uf.count())
	}
}
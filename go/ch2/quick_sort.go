package ch2

import (
	//"math/rand"
	//"time"
)

type QuickSort struct {}

func (s *QuickSort) Sort(data Sortable) {
	// shuffle the array to avoid worst case scenario
	//rand.Seed(time.Now().UnixNano())
	// for i := data.Len() - 1; i > 0; i -= 1 {
	// 	j := rand.Intn(i + 1)
	// 	data.Swap(i, j)
	// }

	sort_quick(data, 0, data.Len() - 1)
}

func sort_quick(data Sortable, lo, hi int) {
	if lo >= hi {
		return
	}

	j := partition(data, lo, hi)
	sort_quick(data, lo, j-1)
	sort_quick(data, j+1, hi)
}

func partition(data Sortable, lo, hi int) int {
	i, j := lo, hi+1
	v := lo
	for {
		// important to avoid infinite loop
		i, j = i+1, j-1

		for ; data.Less(i, v) && i < hi; i += 1 {}
		for ; data.Less(v, j) && j > lo; j -= 1 {}
		if i >= j {
			break
		}
		data.Swap(i, j)
	}
	data.Swap(v, j)
	return j
}
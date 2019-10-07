package ch2

import (
	//"math/rand"
	//"time"
)

type QuickSort3 struct {}

func (s *QuickSort3) Sort(data Sortable) {
	// shuffle the array to avoid worst case scenario
	//rand.Seed(time.Now().UnixNano())
	// for i := data.Len() - 1; i > 0; i -= 1 {
	// 	j := rand.Intn(i + 1)
	// 	data.Swap(i, j)
	// }

	sort_quick3(data, 0, data.Len() - 1)
}

func sort_quick3(data Sortable, lo, hi int) {
	if lo >= hi {
		return
	}

	lt, i, gt := lo, lo+1, hi
	v := lo
	for i <= gt {
		c := data.Compare(i, v)
		if c < 0 {
			data.Swap(lt, i)
			i += 1
			lt += 1
			v = lt
		} else if c > 0 {
			data.Swap(i, gt)
			gt -= 1
		} else {
			i += 1
		}
	}

	sort_quick3(data, lo, lt-1)	
	sort_quick3(data, gt+1, hi)
}
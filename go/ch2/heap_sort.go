package ch2

type HeapSort struct {}

func (s *HeapSort) Sort(data Sortable) {
	// builing a heap: 2N at most
	for i := data.Len()/2; i >= 1; i-- {
		sink(i, data.Len()-1, data)
	}

	// sorting
	for i := data.Len() - 1; i > 0; i-- {
		// setting max element into the its position
		data.Swap(1, i)

		// restoring the heap order
		sink(1, i-1, data)
	}
}

func sink(k int, n int, data Sortable) {
	for 2*k <= n {

		// ch is the largest of two children
		ch := 2*k
		if ch < n && data.Less(ch, ch+1) {
			ch++
		}
		
		if data.Less(k, ch) {
			// restoring max heap order
			data.Swap(ch, k)
			k = ch
		} else {
			break
		}
	}
}
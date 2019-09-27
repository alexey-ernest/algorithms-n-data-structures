package ch2

type InsertionSortSentinel struct {}

func (s *InsertionSortSentinel) Sort(data Sortable) {
	n := data.Len()

	// first searching for the smallest item to avoid inner loop index test
	min := 0
	for i := 1; i < n; i+= 1 {
		if data.Less(i, min) {
			min = i
		}
	}
	data.Swap(0, min)

	for i := 1; i < n; i += 1 {
		for j := i; ; j -= 1 {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
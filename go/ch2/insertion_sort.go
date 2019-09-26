package ch2

type InsertionSort struct {}

func (s *InsertionSort) Sort(data Sortable) {
	n := data.Len()
	for i := 1; i < n; i += 1 {
		for j := i; j > 0; j -= 1 {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
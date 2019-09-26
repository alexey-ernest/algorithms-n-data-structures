package ch2

type SelectionSort struct {}

func (s *SelectionSort) Sort(data Sortable) {
	n := data.Len()
	for i := 0; i < n; i += 1 {
		min := i
		for j := i; j < n; j += 1 {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
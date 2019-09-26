package ch2

type ShellSort struct {}

func (s *ShellSort) Sort(data Sortable) {
	n := data.Len()
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// h-sort the array with offset 1 (important!)
		for i := h; i < n; i += 1 {
			for j := i; j >= h; j -= h {
				if data.Less(j, j-h) {
					data.Swap(j, j-h)
				} else {
					break
				}
			}
		}

		// decrease h
		h = h/3
	}
}
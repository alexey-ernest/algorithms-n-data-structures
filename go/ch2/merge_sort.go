package ch2

type MergeSort struct {}

func (s *MergeSort) Sort(a []int) {
	sort(a, 0, len(a) - 1)
}

func sort(a []int, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high - low)/2
	sort(a, low, mid)
	sort(a, mid + 1, high)
	merge(a, low, mid, high)
}

func merge(a []int, low, mid, high int) {
	i, j := low, mid + 1
	
	temp := make([]int, high - low + 1)
	copy(temp, a[low:high+1])
	
	for k := low; k <= high; k += 1 {
		if i > mid {
			// left part exhausted
			a[k] = temp[j-low]
			j += 1
		} else if j > high {
			// right part exhausted
			a[k] = temp[i-low]
			i += 1
		} else if temp[i-low] < temp[j-low] {
			a[k] = temp[i-low]
			i += 1
		} else {
			a[k] = temp[j-low]
			j += 1
		}
	}
}
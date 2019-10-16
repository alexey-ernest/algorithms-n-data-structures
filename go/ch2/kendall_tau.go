package ch2

// merge-sort based n*logn algorithm
func KendallTauDistance(a, b []int) int {
	check_inputs(a, b)	

	// building value2index map to check relative order 
	// in first array by using values from the second one
	value2index := make(map[int]int)
	for i, v := range a {
		value2index[v] = i
	}

	// counting number of inversions in b by sorting it
	invcount := sort_ktd(b, value2index, 0, len(a) - 1)
	return invcount
}

func sort_ktd(a []int, ind map[int]int, l, r int) int {
	if l >= r {
		return 0
	}

	mid := l + (r-l)/2
	linv := sort_ktd(a, ind, l, mid)
	rinv := sort_ktd(a, ind, mid+1, r)
	
	invcount := merge_ktd(a, ind, l, mid, r) + linv + rinv
	return invcount
}

func merge_ktd(a []int, ind map[int]int, l, mid, r int) int {
	i, j := l, mid+1
	ln := r-l+1
	
	cp := make([]int, ln)
	copy(cp, a[l:r+1])

	invcount := 0
	for k := 0; k < ln; k += 1 {
		if i > mid {
			a[l+k] = cp[j-l]
			j++
		} else if j > r {
			a[l+k] = cp[i-l]
			i++
		} else if ind[cp[i-l]] < ind[cp[j-l]] {
			a[l+k] = cp[i-l]
			i++
		} else {
			invcount += mid - i + 1
			a[l+k] = cp[j-l]
			j++
		}
	}

	return invcount
}

func check_inputs(a, b []int) {
	if len(a) != len(b) {
		panic("a and b have different lengths")
	}

	// checking that a has all unique values
	vals := make(map[int]bool)
	for _, v := range a {
		if val, ok := vals[v]; ok && val {
			panic("a should have only unique values")
		}
		
		vals[v] = true
	}

	for _, v := range b {
		val, ok := vals[v]
		if !ok {
			panic("value of b not found in a")
		}
		if !val {
			panic("b should have only unique values")
		}
		vals[v] = false
	}
}

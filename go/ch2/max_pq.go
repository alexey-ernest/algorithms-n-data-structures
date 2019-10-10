package ch2

type maxPriorityQueue struct {
	data []int
	n int
}

func NewMaxPQ(size int) maxPriorityQueue {
	return maxPriorityQueue{
		data: make([]int, size + 1),
	}
}

func (pq *maxPriorityQueue) Size() int {
	return pq.n
}

func (pq *maxPriorityQueue) IsEmpty() bool {
	return pq.n == 0
}

func (pq *maxPriorityQueue) Insert(key int) {
	if pq.n + 1 == cap(pq.data) {
		panic("max capacity reached")
	}

	pq.n++
	pq.data[pq.n] = key

	// logN
	pq.swim(pq.n)
}

func (pq *maxPriorityQueue) DelMax() int {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	max := pq.data[1]
	pq.data[1] = pq.data[pq.n]
	pq.n--

	// logN
	pq.sink(1)

	return max
}

func (pq *maxPriorityQueue) swim(k int) {
	for k > 1 {
		p := k/2
		if pq.data[p] < pq.data[k] {
			// swap
			pq.data[p], pq.data[k] = pq.data[k], pq.data[p]
			k = p
		} else {
			break
		}
	}
}

func (pq *maxPriorityQueue) sink(k int) {
	for 2*k <= pq.n {
		maxch := 2*k
		if maxch < pq.n && pq.data[maxch+1] > pq.data[maxch] {
			maxch++
		}

		if pq.data[maxch] > pq.data[k] {
			// swap
			pq.data[maxch], pq.data[k] = pq.data[k], pq.data[maxch]
			k = maxch
		} else {
			break
		}
	}
}
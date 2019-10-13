package ch2

type indexMinPQ struct {
	keys []int
	index2key []int
	key2index []int
	n int
}

func NewIndexMinPQ(size int) indexMinPQ {
	return indexMinPQ {
		keys: make([]int, size + 1),
		index2key: make([]int, size + 1),
		key2index: make([]int, size + 1),
	}
}

func (pq *indexMinPQ) Size() int {
	return pq.n
}

func (pq *indexMinPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *indexMinPQ) Insert(i, key int) {
	pq.checkIndex(i)
	if pq.index2key[i] > 0 {
		panic("index already used")
	}
	if pq.n + 1 == cap(pq.keys) {
		panic("pq is full")
	}

	pq.n++
	pq.keys[pq.n] = key
	pq.index2key[i] = pq.n
	pq.key2index[pq.n] = i

	// restore order
	pq.swim(i)
}

func (pq *indexMinPQ) Change(i, key int) {
	pq.checkIndex(i)
	idx := pq.index2key[i]
	if idx == 0 {
		panic("a key does not exist")
	}

	// updating the key
	k := pq.keys[idx]
	pq.keys[idx] = key

	// restore order
	if key > k {
		pq.sink(i)
	} else if key < k {
		pq.swim(i)
	}
}

func (pq *indexMinPQ) Contains(i int) bool {
	pq.checkIndex(i)
	return pq.index2key[i] > 0
}

func (pq *indexMinPQ) Delete(i int) {
	pq.checkIndex(i)
	idx := pq.index2key[i]
	if idx == 0 {
		panic("a key does not exist")
	}

	pq.keys[idx] = pq.keys[pq.n]
	pq.index2key[i] = 0

	lastkeyindex := pq.key2index[pq.n]
	pq.index2key[lastkeyindex] = idx
	pq.key2index[idx] = lastkeyindex
	pq.n--

	// restore order
	pq.sink(lastkeyindex)
}

func (pq *indexMinPQ) Min() int {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	return pq.keys[1]
}

func (pq *indexMinPQ) MinIndex() int {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	return pq.key2index[1]
}

// removes minimal element and returns it's index
func (pq *indexMinPQ) DelMin() int {
	minindex := pq.MinIndex()
	pq.Delete(minindex)
	return minindex
}


// helpers

func (pq *indexMinPQ) checkIndex(i int) {
	if i + 1 >= cap(pq.keys) || i < 0 {
		panic("invalid index")
	}
}

func (pq *indexMinPQ) swim(i int) {
	k := pq.index2key[i]
	for k > 1 && pq.keys[k] < pq.keys[k/2] {
		// swap keys
		pq.keys[k], pq.keys[k/2] = pq.keys[k/2], pq.keys[k]

		// swap indexes
		pq.index2key[pq.key2index[k]], pq.index2key[pq.key2index[k/2]] = pq.index2key[pq.key2index[k/2]], pq.index2key[pq.key2index[k]]
		pq.key2index[k], pq.key2index[k/2] = pq.key2index[k/2], pq.key2index[k]

		k = k/2
	}
}

func (pq *indexMinPQ) sink(i int) {
	k := pq.index2key[i]
	for 2*k <= pq.n {
		c := 2*k
		if c < pq.n && pq.keys[c+1] < pq.keys[c] {
			c++
		}
		if pq.keys[k] > pq.keys[c] {
			// swap keys
			pq.keys[k], pq.keys[c] = pq.keys[c], pq.keys[k]

			// swap indexes
			pq.index2key[pq.key2index[k]], pq.index2key[pq.key2index[c]] = pq.index2key[pq.key2index[c]], pq.index2key[pq.key2index[k]]
			pq.key2index[k], pq.key2index[c] = pq.key2index[c], pq.key2index[k]

			k = c
		} else {
			break
		}
	}
}

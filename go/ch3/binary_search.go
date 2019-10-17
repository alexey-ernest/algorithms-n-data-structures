package ch3

import "fmt"

type binarySearchST struct {
	keys []string
	values []string
	n int
}

func NewBinarySearchST() binarySearchST {
	return binarySearchST {
		keys: make([]string, 0),
		values: make([]string, 0),
		n: 0,
	}
}

func (st *binarySearchST) IsEmpty() bool {
	return st.n == 0
}

func (st *binarySearchST) Size() int {
	return st.n
}

func (st *binarySearchST) Get(key string) string {
	i := st.rank(key)
	if i < st.n && st.keys[i] == key {
		return st.values[i]
	}

	panic(fmt.Sprintf("st does not contain key %s", key))
}

func (st *binarySearchST) Contains(key string) bool {
	i := st.rank(key)
	if i < st.n && st.keys[i] == key {
		return true
	}

	return false
}

func (st *binarySearchST) Put(key, value string) {
	i := st.rank(key)
	if i < st.n && st.keys[i] == key {
		st.values[i] = value
		return
	}

	if i == st.n {
		st.keys = append(st.keys, key)
		st.values = append(st.values, value)
	} else {
		st.keys = append(st.keys, key)
		st.values = append(st.values, value)
		for j := st.n; j > i; j-=1 {
			st.keys[j] = st.keys[j-1]
			st.values[j] = st.values[j-1]
		}
		st.keys[i] = key
		st.values[i] = value
	}
	st.n++
}

func (st *binarySearchST) Delete(key string) {
	i := st.rank(key)
	if i == st.n || st.keys[i] != key {
		panic(fmt.Sprintf("st does not contain key %s", key))
	}

	for j := i; j < st.n-1; j+=1 {
		st.keys[j] = st.keys[j+1]
		st.values[j] = st.values[j+1]
	}

	st.keys = st.keys[:st.n-1]
	st.values = st.values[:st.n-1]
	st.n--
}

func (st *binarySearchST) rank(key string) int {
	l, r := 0, st.n-1
	for l <= r {
		k := l + (r - l) / 2
		if st.keys[k] == key {
			return k
		} else if st.keys[k] > key {
			r = k-1
		} else {
			l = k+1
		}
	}

	return l
}

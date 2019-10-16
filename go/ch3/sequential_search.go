package ch3

type node_ss struct {
	key string
	value string
	next *node_ss
}

type SequentialSearchST struct {
	first *node_ss
	n int
}

func (st *SequentialSearchST) Size() int {
	return st.n
}

func (st *SequentialSearchST) IsEmpty() bool {
	return st.n > 0
}

func (st *SequentialSearchST) Put(key, value string) {
	n := st.first
	for n != nil && n.key != key {
		n = n.next
	}
	if n != nil {
		n.value = value
	} else {
		n := &node_ss{
			key: key,
			value: value,
			next: st.first,
		}
		st.first = n
		st.n++
	}
}

func (st *SequentialSearchST) Get(key string) string {
	n := st.first
	for n != nil && n.key != key {
		n = n.next
	}
	if n == nil {
		return nil
	}

	return n.value
}

func (st *SequentialSearchST) Contains(key string) bool {
	return st.Get(key) != nil
}

func (st *SequentialSearchST) Delete(key string) {
	n := st.first
	var prev *node_ss
	for n != nil && n.key != key {
		prev = n
		n = n.next
	}
	if n == nil {
		panic("ST does not contain key %s", key)
	}

	// removing the node
	if prev != nil {
		prev.next = n.next
	} else {
		st.first = n.next
	}
	st.n--
}
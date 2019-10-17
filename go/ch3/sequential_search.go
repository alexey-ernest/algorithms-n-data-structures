package ch3

import "fmt"

type node_ss struct {
	key string
	value string
	next *node_ss
}

type sequentialSearchST struct {
	first *node_ss
	n int
}

func NewSequentialSearchST() sequentialSearchST {
	return sequentialSearchST {
		first: nil,
		n: 0,
	}
}

func (st *sequentialSearchST) Size() int {
	return st.n
}

func (st *sequentialSearchST) IsEmpty() bool {
	return st.n == 0
}

func (st *sequentialSearchST) Put(key, value string) {
	var n *node_ss
	for n = st.first; n != nil && n.key != key; n = n.next {
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

func (st *sequentialSearchST) Get(key string) string {
	var n *node_ss
	for n = st.first; n != nil && n.key != key; n = n.next {
	}
	if n == nil {
		panic(fmt.Sprintf("st does not contain key %s", key))
	}

	return n.value
}

func (st *sequentialSearchST) Contains(key string) bool {
	var n *node_ss
	for n = st.first; n != nil && n.key != key; n = n.next {
	}
	return n != nil
}

func (st *sequentialSearchST) Delete(key string) {
	n := st.first
	var prev *node_ss
	for n != nil && n.key != key {
		prev = n
		n = n.next
	}
	if n == nil {
		panic(fmt.Sprintf("ST does not contain key %s", key))
	}

	// removing the node
	if prev != nil {
		prev.next = n.next
	} else {
		st.first = n.next
	}
	st.n--
}
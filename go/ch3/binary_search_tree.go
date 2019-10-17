package ch3

import "fmt"

type nodeBST struct {
	key string
	value string
	left *nodeBST
	right *nodeBST
	n int
}

type binarySearchTreeST struct {
	root *nodeBST
}

func NewBST() binarySearchTreeST {
	return binarySearchTreeST{}
}

func (t *binarySearchTreeST) Size() int {
	return t.size(t.root)
}

func (t *binarySearchTreeST) size(n *nodeBST) int {
	if n == nil {
		return 0
	}
	return n.n
}

func (t *binarySearchTreeST) IsEmpty() bool {
	return t.Size() == 0
}

func (t *binarySearchTreeST) Get(key string) string {
	return t.get(t.root, key)
}

func (t *binarySearchTreeST) get(n *nodeBST, key string) string {
	if n == nil {
		panic(fmt.Sprintf("the key %s was not found", key))
	}

	if n.key == key {
		return n.value
	}

	if n.key > key {
		return t.get(n.left, key)
	} else {
		return t.get(n.right, key)
	}
}

func (t *binarySearchTreeST) Contains(key string) bool {
	return t.contains(t.root, key)
}

func (t *binarySearchTreeST) contains(n *nodeBST, key string) bool {
	if n == nil {
		return false
	}

	if n.key == key {
		return true
	}

	if n.key > key {
		return t.contains(n.left, key)
	} else {
		return t.contains(n.right, key)
	}
}

func (t *binarySearchTreeST) Put(key string, value string) {
	t.root = t.put(t.root, key, value)
}

func (t *binarySearchTreeST) put(n *nodeBST, key string, value string) *nodeBST {
	if n == nil {
		// creating new node
		n = &nodeBST{
			key: key,
			value: value,
			n: 1,
		}
		return n
	}

	if n.key == key {
		n.value = value
		return n
	}

	if n.key > key {
		n.left = t.put(n.left, key, value)
	} else {
		n.right = t.put(n.right, key, value)
	}

	// updating n to maintain actual count
	n.n = t.size(n.left) + t.size(n.right) + 1

	return n
}

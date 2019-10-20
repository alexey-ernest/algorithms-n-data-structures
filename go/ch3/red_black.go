package ch3

import "fmt"

// A self-balancing Binary Search Tree with 2*lgN worst case garantees for
// search, put, delete, min, max, select, rank, floor, ceiling operations.
// Average runtine for search-based operations estimated as 1*lgN

type nodeRedBlack struct {
	key string
	value int
	left *nodeRedBlack
	right *nodeRedBlack
	n int
	isRed bool
}

type redBlackBST struct {
	root *nodeRedBlack
}

func NewRedBlackBST() redBlackBST {
	return redBlackBST{}
}

func (t *redBlackBST) Size() int {
	return t.size(t.root)
}

func (t *redBlackBST) size(n *nodeRedBlack) int {
	if n == nil {
		return 0
	}

	return n.n
}

func (t *redBlackBST) IsEmpty() bool {
	return t.size(t.root) == 0
}

func (t *redBlackBST) Contains(key string) bool {
	return t.get(t.root, key) != nil
}

func (t *redBlackBST) Get(key string) int {
	if t.IsEmpty() {
		panic("Red Black BST is empty")
	}

	x := t.get(t.root, key)
	if x == nil {
		panic(fmt.Sprintf("key %q does not exist", key))
	}

	return x.value
}

func (t *redBlackBST) get(n *nodeRedBlack, key string) *nodeRedBlack {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	}

	if n.key > key {
		return t.get(n.left, key)
	} else {
		return t.get(n.right, key)
	}
}

func (t *redBlackBST) isRed(n *nodeRedBlack) bool {
	if n == nil {
		return false
	}

	return n.isRed
}

func (t *redBlackBST) flipColors(n *nodeRedBlack) {
	if n == nil {
		return
	}

	// making children black
	if n.left != nil {
		n.left.isRed = false
	}
	if n.right != nil {
		n.right.isRed = false
	}

	// making node red
	n.isRed = true
}

func (t *redBlackBST) rotateLeft(n *nodeRedBlack) *nodeRedBlack {
	x := n.right
	n.right = x.left
	x.left = n

	x.isRed = n.isRed
	n.isRed = true

	n.n = t.size(n.left) + 1 + t.size(n.right)
	x.n = t.size(x.left) + 1 + t.size(x.right)
	
	return x
}

func (t *redBlackBST) rotateRight(n *nodeRedBlack) *nodeRedBlack {
	x := n.left
	n.left = x.right
	x.right = n

	x.isRed = n.isRed
	n.isRed = true

	n.n = t.size(n.left) + 1 + t.size(n.right)
	x.n = t.size(x.left) + 1 + t.size(x.right)

	return x
}

func (t *redBlackBST) Put(key string, value int) {
	t.root = t.put(t.root, key, value)

	// keeping root black
	t.root.isRed = false
}

func (t *redBlackBST) put(n *nodeRedBlack, key string, value int) *nodeRedBlack {
	if n == nil {
		// search miss, creating a new node with a red link as a part of 3- or 4-node
		return &nodeRedBlack{
			key: key,
			value: value,
			n: 1,
			isRed: true,
		}
	}

	if n.key == key {
		// search hit, updating the value
		n.value = value
		return n
	}

	if n.key > key {
		n.left = t.put(n.left, key, value)
	} else {
		n.right = t.put(n.right, key, value)
	}

	// balancing the tree
	if t.isRed(n.right) && !t.isRed(n.left) {
		// fixing right leaning red link case
		// this can lead to the next case in upper level
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		// making 4-node
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		// fliping 4-node
		t.flipColors(n)
	}

	n.n = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *redBlackBST) Height() int {
	if t.IsEmpty() {
		return 0
	}

	return t.height(t.root)
}

func (t *redBlackBST) height(n *nodeRedBlack) int {
	if n == nil {
		return 0
	}

	lheight := t.height(n.left)
	rheight := t.height(n.right)
	
	height := lheight
	if rheight > lheight {
		height = rheight
	}

	if !n.isRed {
		// black height
		height += 1
	}

	return height
}

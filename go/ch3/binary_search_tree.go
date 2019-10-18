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

func (t *binarySearchTreeST) Min() string {
	if t.Size() == 0 {
		panic("bst is empty")
	}

	return t.min(t.root).key
}

func (t *binarySearchTreeST) min(n *nodeBST) *nodeBST {
	if n.left == nil {
		return n
	}

	return t.min(n.left)
}

func (t *binarySearchTreeST) Max() string {
	if t.Size() == 0 {
		panic("bst is empty")
	}

	return t.max(t.root).key
}

func (t *binarySearchTreeST) max(n *nodeBST) *nodeBST {
	if n.right == nil {
		return n
	}

	return t.max(n.right)
}

func (t *binarySearchTreeST) Floor(key string) string {
	if t.Size() == 0 {
		panic("bst is empty")
	}

	if t.Min() > key {
		panic(fmt.Sprintf("there are no keys <= %q", key))
	}

	floor := t.floor(t.root, key)
	return floor.key
}

func (t *binarySearchTreeST) floor(n *nodeBST, key string) *nodeBST {
	if n == nil {
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key > key {
		// exploring left subtree, floor must be there
		return t.floor(n.left, key)
	}

	// checking whether we have a key in the right subtree, otherwise returning current root
	r := t.floor(n.right, key)
	if r != nil {
		// search hit
		return r
	}

	// if there is no key in the right, then returning current root which is a smaller key
	return n
}

func (t *binarySearchTreeST) Ceiling(key string) string {
	if t.Size() == 0 {
		panic("bst is empty")
	}

	if t.Max() < key {
		panic(fmt.Sprintf("there are no keys >= %q", key))
	}

	ceiling := t.ceiling(t.root, key)
	return ceiling.key
}

func (t *binarySearchTreeST) ceiling(n *nodeBST, key string) *nodeBST {
	if n == nil {
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key < key {
		// it should be in the right subtree
		return t.ceiling(n.right, key)
	}

	l := t.ceiling(n.left, key)
	if l != nil {
		// search hit
		return l
	}

	// if there is no key in the left, than next key bigger than left is a current root
	return n
}

func (t *binarySearchTreeST) Select(k int) string {
	if k < 0 || k >= t.Size() {
		panic("index out of range")
	}

	return t.selectKey(t.root, k)
}

func (t *binarySearchTreeST) selectKey(n *nodeBST, k int) string {
	lc := t.size(n.left)
	if lc == k {
		return n.key
	}

	if lc > k {
		return t.selectKey(n.left, k)
	} else {
		return t.selectKey(n.right, k - lc - 1)
	}
}

func (t *binarySearchTreeST) Rank(key string) int {
	if t.Size() == 0 {
		panic("bst is empty")
	}

	return t.rank(t.root, key)
}

func (t *binarySearchTreeST) rank(n *nodeBST, key string) int {
	if n == nil {
		return 0
	}

	lc := t.size(n.left)
	if n.key == key {
		return lc
	}

	if n.key > key {
		return t.rank(n.left, key)
	} else {
		return lc + 1 + t.rank(n.right, key)
	}
}

func (t *binarySearchTreeST) DeleteMin() {
	t.root = t.deleteMin(t.root)
}

func (t *binarySearchTreeST) deleteMin(n *nodeBST) *nodeBST {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}
	
	n.left = t.deleteMin(n.left)
	n.n = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *binarySearchTreeST) Delete(key string) {
	t.root = t.delete(t.root, key)
}

func (t *binarySearchTreeST) delete(n *nodeBST, key string) *nodeBST {
	if n == nil {
		return nil
	}

	if n.key == key {
		// replacing by successor (we can do similar with precedessor)
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		new := t.min(n.right)
		new.right = t.deleteMin(n.right)
		new.left = n.left		
		n = new
	} else if n.key > key {
		n.left = t.delete(n.left, key)
	} else {
		n.right = t.delete(n.right, key)
	}

	// updating counter
	n.n = t.size(n.left) + 1 + t.size(n.right)
	return n
}


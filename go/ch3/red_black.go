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

func (t *redBlackBST) panicIfEmpty() {
	if t.IsEmpty() {
		panic("Red Black BST is empty")
	}
}

func (t *redBlackBST) Get(key string) int {
	t.panicIfEmpty()

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

	return height + 1
}

func (t *redBlackBST) IsRedBlack() bool {
	balanced, _ := t.isBalanced(t.root)
	return balanced && t.is23(t.root)
}

func (t *redBlackBST) is23(n *nodeRedBlack) bool {
	if n == nil {
		return true
	}

	if t.isRed(n.right) {
		// it should has only left leaning red links
		return false
	}

	if t.isRed(n) && t.isRed(n.left) {
		// no node should be connected by two red links
		return false
	}

	return t.is23(n.left) && t.is23(n.right)
}

func (t *redBlackBST) isBalanced(n *nodeRedBlack) (bool, int) {
	if n == nil {
		// null node is black by default
		return true, 1
	}

	lb, l := t.isBalanced(n.left)
	rb, r := t.isBalanced(n.right)

	b := l
	if r > l {
		b = r
	}

	if !t.isRed(n) {
		b += 1
	}

	return lb && rb && l == r, b
}

func (t *redBlackBST) Min() string {
	t.panicIfEmpty()
	return t.min(t.root).key
}

func (t *redBlackBST) min(n *nodeRedBlack) *nodeRedBlack {
	if n.left == nil {
		return n
	}

	return t.min(n.left)
}

func (t *redBlackBST) Max() string {
	t.panicIfEmpty()
	return t.max(t.root).key
}

func (t *redBlackBST) max(n *nodeRedBlack) *nodeRedBlack {
	if n.right == nil {
		return n
	}

	return t.max(n.right)
}

func (t *redBlackBST) Floor(key string) string {
	t.panicIfEmpty()

	floor := t.floor(t.root, key)
	if floor == nil {
		panic(fmt.Sprintf("there are no keys <= %q", key))
	}

	return floor.key
}

func (t *redBlackBST) floor(n *nodeRedBlack, key string) *nodeRedBlack {
	if n == nil {
		// search miss
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key > key {
		// floor must be in the left sub-tree
		return t.floor(n.left, key)
	}

	// key could be in the right sub-tree, if not, using current root
	floor := t.floor(n.right, key)
	if floor != nil {
		return floor
	}

	return n
}

func (t *redBlackBST) Ceiling(key string) string {
	t.panicIfEmpty()

	ceiling := t.ceiling(t.root, key)
	if ceiling == nil {
		panic(fmt.Sprintf("there are no keys >= %q", key))
	}

	return ceiling.key
}

func (t *redBlackBST) ceiling(n *nodeRedBlack, key string) *nodeRedBlack {
	if n == nil {
		// search miss
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key < key {
		// ceiling must be in the right sub-tree
		return t.ceiling(n.right, key)
	}

	// the key could be in the left sub-tree, if not, using current root
	ceiling := t.ceiling(n.left, key)
	if ceiling != nil {
		return ceiling
	}

	return n
}

func (t *redBlackBST) Select(k int) string {
	if k < 0 || k >= t.Size() {
		panic("index out of range")
	}

	return t.selectNode(t.root, k).key
}

func (t *redBlackBST) selectNode(n *nodeRedBlack, k int) *nodeRedBlack {
	if t.size(n.left) == k {
		return n
	}

	if t.size(n.left) > k {
		return t.selectNode(n.left, k)
	}

	k = k - t.size(n.left) - 1
	return t.selectNode(n.right, k)
}

func (t *redBlackBST) Rank(key string) int {
	t.panicIfEmpty()
	return t.rank(t.root, key)
}

func (t *redBlackBST) rank(n *nodeRedBlack, key string) int {
	if n == nil {
		return 0
	}

	if n.key == key {
		return t.size(n.left)
	}

	if n.key > key {
		return t.rank(n.left, key)
	}

	return t.size(n.left) + 1 + t.rank(n.right, key)
}

func (t *redBlackBST) Keys(lo, hi string) []string {
	if lo < t.Min() || hi > t.Max() {
		panic("keys out of range")
	}

	return t.keys(t.root, lo, hi)
}

func (t *redBlackBST) keys(n *nodeRedBlack, lo, hi string) []string {
	if n == nil {
		return nil
	}

	if n.key < lo {
		return t.keys(n.right, lo, hi)
	} else if n.key > hi {
		return t.keys(n.left, lo, hi)
	}

	l := t.keys(n.left, lo, hi)
	r := t.keys(n.right, lo, hi)
	
	var keys []string
	if l != nil {
		keys = append(l, n.key)
	} else {
		keys = []string{n.key}
	}

	if r != nil {
		keys = append(keys, r...)
	}

	return keys
}

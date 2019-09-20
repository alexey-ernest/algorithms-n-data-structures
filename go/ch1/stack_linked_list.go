package ch1

type byteNode struct {
	item byte
	next *byteNode
}

type StackLinkedList struct {
	first *byteNode
	n int
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.first == nil;
}

func (s *StackLinkedList) Size() int {
	return s.n;
}

func (s *StackLinkedList) Push(c byte) {
	n := byteNode {
		item: c,
		next: s.first,
	}

	s.first = &n
	s.n += 1;
}

func (s *StackLinkedList) Pop() byte {
	n := s.first
	s.first = n.next
	s.n -= 1;
	return n.item
}

package ch1

type byteNode struct {
	item byte
	next *byteNode
}

type StackLinkedList struct {
	first *byteNode
}

func (s *StackLinkedList) Push(c byte) {
	n := byteNode {
		item: c,
		next: s.first,
	}

	s.first = &n
}

func (s *StackLinkedList) Pop() byte {
	n := s.first
	s.first = n.next
	return n.item
}

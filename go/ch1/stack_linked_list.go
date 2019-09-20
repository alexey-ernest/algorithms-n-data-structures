package ch1

type node struct {
	item byte
	next *node
}

type StackLinkedList struct {
	first *node
}

func (s *StackLinkedList) Push(c byte) {
	n := node {
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

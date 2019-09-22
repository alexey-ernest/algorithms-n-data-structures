package ch1

type _stringNode struct {
	item string
	next *_stringNode
}

type StackLinkedListString struct {
	first *_stringNode
	n int
}

func (s *StackLinkedListString) IsEmpty() bool {
	return s.first == nil
}

func (s *StackLinkedListString) Size() int {
	return s.n
}

func (s *StackLinkedListString) Push(c string) {
	n := _stringNode {
		item: c,
		next: s.first,
	}

	s.first = &n
	s.n += 1
}

func (s *StackLinkedListString) Pop() string {
	n := s.first
	s.first = n.next
	s.n -= 1
	return n.item
}

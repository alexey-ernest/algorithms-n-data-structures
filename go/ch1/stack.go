package ch1

type Stack struct {
	data []byte
}

func (s *Stack) Push(c byte) {
	if s.data == nil {
		s.data = make([]byte, 0)
	}

	s.data = append(s.data, c)
}

func (s *Stack) Pop() byte {
	if s.data == nil {
		s.data = make([]byte, 0)
	}

	if len(s.data) == 0 {
		return byte(0)
	}

	el, r := s.data[0], s.data[1:]
	s.data = r
	return el
}

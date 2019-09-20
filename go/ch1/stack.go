package ch1

type Stack struct {
	data []byte
}

func (s *Stack) Push(c byte) {
	s.data = append(s.data, c)
}

func (s *Stack) Pop() byte {
	if len(s.data) == 0 {
		return byte(0)
	}

	el, r := s.data[len(s.data) - 1], s.data[:len(s.data) - 1]
	s.data = r
	return el
}

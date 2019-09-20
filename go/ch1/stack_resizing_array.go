package ch1

//import "fmt"

type StackResizingArray struct {
	data []byte // slice as a pointer to an array
}

func (s *StackResizingArray) Push(c byte) {
	//fmt.Printf("len %d, cap %d\n", len(s.data), cap(s.data))
	if len(s.data) == cap(s.data) {
		// resizing array
		newlen := len(s.data) * 2
		if newlen == 0 {
			newlen = 1 
		}

		newdata := make([]byte, len(s.data), newlen) // creating double sized array
		copy(newdata, s.data) // copy elements
		s.data = newdata
	}
	s.data = append(s.data, c)
}

func (s *StackResizingArray) Pop() byte {	
	if len(s.data) == 0 {
		return byte(0)
	}

	if len(s.data) <= cap(s.data) / 4 {
		// shrinking
		newlen := cap(s.data) / 2
		newdata := make([]byte, len(s.data), newlen)
		copy(newdata, s.data)
		s.data = newdata
	}
	//fmt.Printf("len %d, cap %d\n", len(s.data), cap(s.data))

	r, val := s.data[:len(s.data) - 1], s.data[len(s.data) - 1]
	s.data = r
	return val
}
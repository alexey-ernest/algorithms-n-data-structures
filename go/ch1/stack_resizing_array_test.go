package ch1

import "testing"

func TestPushStackResizingArray(t *testing.T) {
	s := new(StackResizingArray)
	s.Push('a')

	got := s.Pop()
	if got != 'a' {
		t.Errorf("'a' != %q", got)
	}
}

func TestPush2NPopStackResizingArray(t *testing.T) {
	s := new(StackResizingArray)
	s.Push('a')
	s.Push('b')

	got := [2]byte{s.Pop(), s.Pop()}
	if got != [2]byte{'b', 'a'} {
		t.Errorf("['a' 'b'] != %+q", got)
	}
}

func TestPush10PopStackResizingArray(t *testing.T) {
	s := new(StackResizingArray)
	var b1 [10]byte
	for i := 0; i < 10; i += 1 {
		c := byte('a' + i)
		s.Push(c)
		b1[9 - i] = c
	}

	var b2 [10]byte
	for i := 0; i < 10; i += 1 {
		c := s.Pop()
		b2[i] = c
	}

	if b1 != b2 {
		t.Errorf("%q != %q", b1, b2)
	}
}
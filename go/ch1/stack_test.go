package ch1

import "testing"

func TestPushNPopStack(t *testing.T) {
	s := new(Stack)
	s.Push('a')

	got := s.Pop()
	if got != 'a' {
		t.Errorf("'a' != %q", got)
	}
}

func TestPush2NPopStack(t *testing.T) {
	s := new(Stack)
	s.Push('a')
	s.Push('b')

	got := [2]byte{s.Pop(), s.Pop()}
	if got != [2]byte{'b', 'a'} {
		t.Errorf("['a' 'b'] != %+q", got)
	}
}

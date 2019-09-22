package ch1

import "testing"

func TestFixParentheses(t *testing.T) {
	input := "1+2)*3-4)*5-6)))"
	got := FixParentheses(input)
	expected := "((1+2)*((3-4)*(5-6)))"
	if got != expected {
		t.Errorf("%s != %s", expected, got)
	}
}
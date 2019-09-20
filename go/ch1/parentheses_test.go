package ch1

import "testing"

func TestBalancedParenthesesGood(t *testing.T) {
	input := "[()]{}{[()()]()}";
	if !IsBalancedParentheses(input) {
		t.Errorf("%q should be balanced", input);
	}
}

func TestBalancedParenthesesBad(t *testing.T) {
	input := "[(])";
	if IsBalancedParentheses(input) {
		t.Errorf("%q should not be balanced", input);
	}
}

func TestBalancedParenthesesOnlyOpen(t *testing.T) {
	input := "[({";
	if IsBalancedParentheses(input) {
		t.Errorf("%q should not be balanced", input);
	}
}

func TestBalancedParenthesesOnlyClose(t *testing.T) {
	input := "})]";
	if IsBalancedParentheses(input) {
		t.Errorf("%q should not be balanced", input);
	}
}

func TestBalancedParenthesesEmpty(t *testing.T) {
	input := "";
	if !IsBalancedParentheses(input) {
		t.Errorf("empty is balanced", input);
	}
}
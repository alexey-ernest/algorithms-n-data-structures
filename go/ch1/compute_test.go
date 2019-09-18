package ch1

import "testing"

func TestCompute1p1(t *testing.T) {
	got := compute("(1 + 1)")
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestCompute5minus3(t *testing.T) {
	got := compute("(5 - 3)")
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestCompute2m3(t *testing.T) {
	got := compute("(2 * 3)")
	if got != 6 {
		t.Errorf("expected 6, got %d", got)
	}
}

func TestComputeSumAndMult(t *testing.T) {
	got := compute("(1 + (2 * 3))")
	if got != 7 {
		t.Errorf("expected 7, got %d", got)
	}
}

func TestComputeDiv(t *testing.T) {
	got := compute("(8 / 2)")
	if got != 4 {
		t.Errorf("expected 4, got %d", got)
	}
}

func TestCompute2Pow5(t *testing.T) {
	got := compute("(2 ^ 5)")
	if got != 32 {
		t.Errorf("expected 32, got %d", got)
	}
}

func TestCompute(t *testing.T) {
	got := compute("((((1 + (2*3)) * 5) + 5) / 4)")
	if got != 10 {
		t.Errorf("expected 10, got %d", got)
	}
}

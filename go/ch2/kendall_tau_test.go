package ch2

import "testing"

func TestKendallTauBasic(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 1, 2, 5}

	kd := KendallTauDistance(a, b)
	if kd != 4 {
		t.Errorf("Kendall tau distance should be 4")
	}
}

func TestKendallTauBasic2(t *testing.T) {
	a := []int{0, 3, 1, 6, 2, 5, 4}
	b := []int{1, 0, 3, 6, 4, 2, 5}

	kd := KendallTauDistance(a, b)
	if kd != 4 {
		t.Errorf("Kendall tau distance should be 4, got %d", kd)
	}
}
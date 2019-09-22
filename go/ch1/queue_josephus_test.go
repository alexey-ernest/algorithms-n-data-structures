package ch1

import "testing"

func TestJosephus(t *testing.T) {
	answer := []string{"1", "3", "5", "0", "4", "2", "6"}
	got := JosephusEliminateOrder(7, 2)

	for i, v := range got {
		if answer[i] != v {
			t.Errorf("%q != %q", answer, got)
			return
		}
	}
}
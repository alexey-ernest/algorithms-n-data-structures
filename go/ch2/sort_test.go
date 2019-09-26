package ch2

import (
	"testing"
	"math/rand"
)

func validateSort(data []int) bool {
	for i := 1; i < len(data); i += 1 {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func TestSelectionSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = rand.Int()
	}

	s := SelectionSort{}
	s.Sort(SortableInt(input))

	if !validateSort(input) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkSelectionSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}
	s := SelectionSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkSelectionSort1k(b *testing.B) {
	benchmarkSelectionSort(1000, b)
}

func BenchmarkSelectionSort10k(b *testing.B) {
	benchmarkSelectionSort(10000, b)
}

func BenchmarkSelectionSort100k(b *testing.B) {
	benchmarkSelectionSort(100000, b)
}


func TestInsertionSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = int(rand.Int())
	}

	s := InsertionSort{}
	s.Sort(SortableInt(input))

	if !validateSort(input) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkInsertionSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}
	s := InsertionSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkInsertionSort1k(b *testing.B) {
	benchmarkInsertionSort(1000, b)
}

func BenchmarkInsertionSort10k(b *testing.B) {
	benchmarkInsertionSort(10000, b)
}

func BenchmarkInsertionSort100k(b *testing.B) {
	benchmarkInsertionSort(100000, b)
}

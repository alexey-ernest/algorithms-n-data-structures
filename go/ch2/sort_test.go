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

func TestInsertionSortSentinel(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = int(rand.Int())
	}

	s := InsertionSortSentinel{}
	s.Sort(SortableInt(input))

	if !validateSort(input) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkInsertionSortSentinel(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}
	s := InsertionSortSentinel{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkInsertionSortSentinel1k(b *testing.B) {
	benchmarkInsertionSortSentinel(1000, b)
}

func BenchmarkInsertionSortSentinel10k(b *testing.B) {
	benchmarkInsertionSortSentinel(10000, b)
}

func BenchmarkInsertionSortSentinel100k(b *testing.B) {
	benchmarkInsertionSortSentinel(100000, b)
}

func TestShellSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = int(rand.Int())
	}

	s := ShellSort{}
	s.Sort(SortableInt(input))

	if !validateSort(input) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkShellSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}
	s := ShellSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkShellSort1k(b *testing.B) {
	benchmarkShellSort(1000, b)
}

func BenchmarkShellSort10k(b *testing.B) {
	benchmarkShellSort(10000, b)
}

func BenchmarkShellSort100k(b *testing.B) {
	benchmarkShellSort(100000, b)
}

func BenchmarkShellSort1m(b *testing.B) {
	benchmarkShellSort(1000000, b)
}

func BenchmarkShellSort10m(b *testing.B) {
	benchmarkShellSort(10000000, b)
}

func TestMergeSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = int(rand.Int())
	}

	s := MergeSort{}
	s.Sort(input)

	if !validateSort(input) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkMergeSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}
	s := MergeSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(input)
	}
}

func BenchmarkMergeSort1k(b *testing.B) {
	benchmarkMergeSort(1000, b)
}

func BenchmarkMergeSort10k(b *testing.B) {
	benchmarkMergeSort(10000, b)
}

func BenchmarkMergeSort100k(b *testing.B) {
	benchmarkMergeSort(100000, b)
}

func BenchmarkMergeSort1m(b *testing.B) {
	benchmarkMergeSort(1000000, b)
}

func BenchmarkMergeSort10m(b *testing.B) {
	benchmarkMergeSort(10000000, b)
}

package ch2

import (
	"testing"
	"math/rand"
	"sort"
)

func validateSort(data Sortable) bool {
	for i := 1; i < data.Len(); i += 1 {
		if data.Less(i, i-1) {
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

	if !validateSort(SortableInt(input)) {
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

	if !validateSort(SortableInt(input)) {
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

	if !validateSort(SortableInt(input)) {
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

	if !validateSort(SortableInt(input)) {
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

	if !validateSort(SortableInt(input)) {
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

func BenchmarkMergeSort50m(b *testing.B) {
	benchmarkMergeSort(50000000, b)
}

func TestQuickSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = rand.Int()
	}

	s := QuickSort{}
	s.Sort(SortableInt(input))

	if !validateSort(SortableInt(input)) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func TestQuickSortAsc(t *testing.T) {
	input := make([]int, 1000)
	last := 0
	for i := range input {
		last += rand.Intn(10)
		input[i] = last
	}

	s := QuickSort{}
	s.Sort(SortableInt(input))

	if !validateSort(SortableInt(input)) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func TestQuickSortBytes(t *testing.T) {
	input := []byte{'K', 'R', 'A', 'T', 'E', 'L', 'E', 'P', 'U', 'I', 'M', 'Q', 'C', 'X', 'O', 'S'}

	s := QuickSort{}
	s.Sort(SortableByte(input))

	if !validateSort(SortableByte(input)) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkQuickSort(n int, b *testing.B) {
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

func BenchmarkQuickSort1k(b *testing.B) {
	benchmarkQuickSort(1000, b)
}

func BenchmarkQuickSort10k(b *testing.B) {
	benchmarkQuickSort(10000, b)
}

func BenchmarkQuickSort100k(b *testing.B) {
	benchmarkQuickSort(100000, b)
}

func BenchmarkQuickSort1m(b *testing.B) {
	benchmarkQuickSort(1000000, b)
}

func BenchmarkQuickSort10m(b *testing.B) {
	benchmarkQuickSort(10000000, b)
}

func TestQuickSort3Asc(t *testing.T) {
	input := make([]int, 1000)
	last := 0
	for i := range input {
		last += rand.Intn(10)
		input[i] = last
	}

	s := QuickSort3{}
	s.Sort(SortableInt(input))

	if !validateSort(SortableInt(input)) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func TestQuickSort3Bytes(t *testing.T) {
	input := []byte{'R', 'B', 'W', 'W', 'R', 'W', 'B', 'R', 'R', 'W', 'B', 'R'}

	s := QuickSort3{}
	s.Sort(SortableByte(input))

	if !validateSort(SortableByte(input)) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkQuickSort3(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Intn(10)
	}

	s := ShellSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkQuickSort3_1k(b *testing.B) {
	benchmarkQuickSort3(1000, b)
}

func BenchmarkQuickSort3_10k(b *testing.B) {
	benchmarkQuickSort3(10000, b)
}

func BenchmarkQuickSort3_100k(b *testing.B) {
	benchmarkQuickSort3(100000, b)
}

func BenchmarkQuickSort3_1m(b *testing.B) {
	benchmarkQuickSort3(1000000, b)
}

func BenchmarkQuickSort3_10m(b *testing.B) {
	benchmarkQuickSort3(10000000, b)
}

func BenchmarkQuickSort3_50m(b *testing.B) {
	benchmarkQuickSort3(50000000, b)
}

func benchmarkNativeSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		sort.Sort(SortableInt(input))
	}
}

func BenchmarkNativeSort1k(b *testing.B) {
	benchmarkNativeSort(1000, b)
}

func BenchmarkNativeSort10k(b *testing.B) {
	benchmarkNativeSort(10000, b)
}

func BenchmarkNativeSort100k(b *testing.B) {
	benchmarkNativeSort(100000, b)
}

func BenchmarkNativeSort1m(b *testing.B) {
	benchmarkNativeSort(1000000, b)
}

func BenchmarkNativeSort10m(b *testing.B) {
	benchmarkNativeSort(10000000, b)
}

func BenchmarkNativeSort50m(b *testing.B) {
	benchmarkNativeSort(50000000, b)
}

func TestHeapSort(t *testing.T) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = rand.Int()
	}
	input = append([]int{0}, input...)

	s := HeapSort{}
	s.Sort(SortableInt(input))

	if !validateSort(SortableInt(input[1:])) {
		t.Errorf("%+v is not sorted properly", input)
	}
}

func benchmarkHeapSort(n int, b *testing.B) {
	input := make([]int, n)
	for i := range input {
		input[i] = rand.Int()
	}

	s := HeapSort{}

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		s.Sort(SortableInt(input))
	}
}

func BenchmarkHeapSort1m(b *testing.B) {
	benchmarkHeapSort(1000000, b)
}

func BenchmarkHeapSort10m(b *testing.B) {
	benchmarkHeapSort(10000000, b)
}

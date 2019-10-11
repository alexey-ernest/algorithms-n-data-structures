package ch2

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMaxPQAsc(t *testing.T) {
	maxsize := 5
	pq := NewMaxPQ(maxsize+1)
	for i := 0; i < 10; i += 1 {
		pq.Insert(i)
		if pq.Size() > maxsize {
			pq.DelMax()
		}
	}

	res := [5]int{}
	for i := 0; i < 5; i += 1 {
		res[i] = pq.DelMax()
	}

	exp := [5]int{4, 3, 2, 1, 0}
	if  res != exp {
		t.Errorf("%+v != %+v", res, exp)
	}
}

func TestMaxPQDesc(t *testing.T) {
	maxsize := 5
	pq := NewMaxPQ(maxsize+1)
	data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(data); i += 1 {
		pq.Insert(data[i])
		if pq.Size() > maxsize {
			pq.DelMax()
		}
	}

	res := [5]int{}
	for i := 0; i < 5; i += 1 {
		res[i] = pq.DelMax()
	}

	exp := [5]int{5, 4, 3, 2, 1}
	if  res != exp {
		t.Errorf("%+v != %+v", res, exp)
	}
}

func TestMaxPQMixed(t *testing.T) {
	maxsize := 5
	pq := NewMaxPQ(maxsize+1)
	data := []int{1, 3, 8, 4, 2, 9, 10, 5, 6, 7}
	for i := 0; i < len(data); i += 1 {
		pq.Insert(data[i])
		if pq.Size() > maxsize {
			pq.DelMax()
		}
	}

	res := [5]int{}
	for i := 0; i < 5; i += 1 {
		res[i] = pq.DelMax()
	}

	exp := [5]int{5, 4, 3, 2, 1}
	if  res != exp {
		t.Errorf("%+v != %+v", res, exp)
	}
}

func TestMaxPQRandom(t *testing.T) {
	maxsize := 5
	pq := NewMaxPQ(maxsize+1)
	minarr := []int{}
	for i := 0; i < 1000; i += 1 {
		key := rand.Int()
		pq.Insert(key)
		if pq.Size() > maxsize {
			pq.DelMax()
		}

		if len(minarr) < maxsize {
			minarr = append(minarr, key)
		} else {
			minmax := 0
			for j := 1; j < len(minarr); j += 1 {
				if minarr[j] > minarr[minmax] {
					minmax = j
				}
			}
			if key < minarr[minmax] {
				minarr[minmax] = key
			}
		}
	}

	res := [5]int{}
	for i := 0; i < 5; i += 1 {
		res[i] = pq.DelMax()
	}

	// sorting minarr desc using merge sort
	for i := 1; i < len(minarr); i += 1 {
		for j := i; j > 0; j -= 1 {
			if minarr[j] > minarr[j-1] {
				minarr[j], minarr[j-1] = minarr[j-1], minarr[j]
			} else {
				break
			}
		}
	}

	exp := [5]int{}
	for i := range minarr {
		exp[i] = minarr[i]
	}

	if  res != exp {
		t.Errorf("%+v != %+v", res, exp)
	}
}

func benchmarkMaxPQConcurrency(maxsize int, numthreads int, b *testing.B) {
	pq := NewMaxPQ(maxsize+1)
	mu := sync.Mutex{}
	res := make(chan float64, numthreads)

	// init pq to have some data
	for i := 0; i < maxsize; i += 1 {
		pq.Insert(rand.Int())
	}

	// modifying queue from `numthreads` simultaneusely
	b.ResetTimer()
	for i := 0; i < numthreads; i += 1 {
		go (func(idx int) {
			start := time.Now()
			for j := 0; j < b.N; j += 1 {
				key := rand.Int()
				
				mu.Lock()
				pq.Insert(key)
				if pq.Size() > maxsize {
					pq.DelMax()
				}
				mu.Unlock()
			}
			elapsed := time.Since(start)
			avg := float64(b.N)/elapsed.Seconds()
			//fmt.Printf("thread %d: \telapsed %s, \ttotal %d, \tavg %.2f op/s\n", idx, elapsed, b.N, avg)

			res <- avg
		})(i)
	}

	total := 0.0
	for i := 0; i < numthreads; i+= 1 {
		avg := <- res
		total += avg
	}

	fmt.Printf("total %d ops, %d ops/thread, performance %.2f op/s\n", b.N * numthreads, b.N, total)
}

func BenchmarkMaxPQConcurrency1m10threads(b *testing.B) {
	benchmarkMaxPQConcurrency(1000000, 10, b)
}

func BenchmarkMaxPQConcurrency1m100threads(b *testing.B) {
	benchmarkMaxPQConcurrency(1000000, 100, b)
}

func BenchmarkMaxPQConcurrency1m1kthreads(b *testing.B) {
	benchmarkMaxPQConcurrency(1000000, 1000, b)
}

func BenchmarkMaxPQConcurrency1m10kthreads(b *testing.B) {
	benchmarkMaxPQConcurrency(1000000, 10000, b)
}

func benchmarkMaxPQ(n int, b *testing.B) {
	pq := NewMaxPQ(n)

	// init with some values
	for i := 0; i < n; i += 1 {
		pq.Insert(rand.Int())
	}

	b.ResetTimer()
	// measure insertion time in an initialized PQ of size n
	for i := 0; i < b.N; i += 1 {
		pq.DelMax()
		pq.Insert(rand.Int())
	}
}

func BenchmarkMaxPQ10k(b *testing.B) {
	benchmarkMaxPQ(10000, b)
}

func BenchmarkMaxPQ100k(b *testing.B) {
	benchmarkMaxPQ(100000, b)
}

func BenchmarkMaxPQ1m(b *testing.B) {
	benchmarkMaxPQ(1000000, b)
}
package ch2

import (
	"math/rand"
	"testing"
)

func TestIndexMinQueueOne(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 5)
	res := minpq.Min()

	if res != 5 {
		t.Errorf("actual %+v != expected %+v", res, 5)
	}
}

func TestIndexMinQueueTwo(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 6)
	minpq.Insert(1, 5)
	
	res := [2]int{}
	res[0] = minpq.Min()
	minpq.DelMin()
	res[1] = minpq.Min()

	exp := [2]int{5, 6}
	if res != exp {
		t.Errorf("actual %+v != expected %+v", res, exp)
	}
}

func TestIndexMinQueueThree(t *testing.T) {
	minpq := NewIndexMinPQ(10)
	minpq.Insert(0, 6)
	minpq.Insert(1, 5)
	minpq.Insert(2, 4)
	
	res := [3]int{}
	res[0] = minpq.Min()
	minpq.DelMin()
	res[1] = minpq.Min()
	minpq.DelMin()
	res[2] = minpq.Min()
	minpq.DelMin()

	exp := [3]int{4, 5, 6}
	if res != exp {
		t.Errorf("actual %+v != expected %+v", res, exp)
	}

	if !minpq.IsEmpty() {
		t.Errorf("pq should be empty")
	}
}

func TestIndexMinQueueRandom(t *testing.T) {
	minpq := NewIndexMinPQ(100)
	emptyindex := 0
	for i := 0; i < 1000; i += 1 {
		emptyindex = i
		if minpq.Size() == 100 {
			emptyindex = minpq.DelMin()
		}
		minpq.Insert(emptyindex, rand.Intn(100))
	}

	res := [100]int{}
	for i := range res {
		res[i] = minpq.Min()
		minpq.DelMin()
	}

	if !minpq.IsEmpty() {
		t.Errorf("pq should be empty after all")
	}

	for i := 1; i < 100; i += 1 {
		if res[i] < res[i-1] {
			t.Errorf("invalid order")
		}
	}
}
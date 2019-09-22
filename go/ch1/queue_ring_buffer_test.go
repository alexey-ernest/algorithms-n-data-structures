package ch1

import (
	"fmt"
	"testing"
)

func TestQueueRingBuffer(t *testing.T) {
	q := NewQueueRingBuffer(10)
	if q.IsEmpty() != true {
		t.Errorf("newly created queue should be empty")
	}
	if q.Size() != 0 {
		t.Errorf("newly created queue should has Size() == 0")
	}
}

func TestQueueNDequeueRingBuffer(t *testing.T) {
	q := NewQueueRingBuffer(10)
	el := "abc"
	q.Enqueue(el)

	got := q.Dequeue()
	if got != "abc" {
		t.Errorf("%q != %q", el, got)
	}
}

func TestQueueNDequeueWrapAroundRingBuffer(t *testing.T) {
	var a [20]string
	q := NewQueueRingBuffer(10)
	for i := 0; i < 10; i+=1 {
		a[i] = fmt.Sprintf("%s", string('a' + i))
		q.Enqueue(a[i])
	}

	var b [20]string
	for i := 0; i < 5; i +=1 {
		b[i] = q.Dequeue()
	}

	for i := 10; i < 15; i+=1 {
		a[i] = fmt.Sprintf("%s", string('a' + i))
		q.Enqueue(a[i])
	}

	for i := 5; i < 15; i +=1 {
		b[i] = q.Dequeue()
	}

	if a != b {
		t.Errorf("%q != %q", a, b)
	}
}
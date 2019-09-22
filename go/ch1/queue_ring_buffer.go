package ch1

type queueRingBuffer struct {
	data []string
	start int
	end int
	n int
}

func NewQueueRingBuffer(size int) queueRingBuffer {
	buf := make([]string, size, size)
	q := queueRingBuffer {
		data: buf,
	}

	return q
}

func (q *queueRingBuffer) IsEmpty() bool {
	return q.n == 0
}

func (q *queueRingBuffer) Size() int {
	return q.n
}

func (q *queueRingBuffer) Enqueue(s string) {
	// blocking until we have enough space
	for q.n == cap(q.data) {}

	q.data[q.end] = s
	q.end += 1
	if q.end == cap(q.data) {
		q.end = 0
	}

	q.n += 1
}

func (q *queueRingBuffer) Dequeue() string {
	// block until we have some work
	for q.n == 0 {}

	e := q.data[q.start]
	q.start += 1
	if q.start == cap(q.data) {
		q.start = 0
	}

	q.n -= 1
	return e
}
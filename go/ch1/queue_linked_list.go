package ch1

type stringNode struct {
	item string
	next *stringNode
}

type QueueLinkedList struct {
	first *stringNode
	last *stringNode
	n int
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.first == nil;
}

func (q *QueueLinkedList) Size() int {
	return q.n;
}

func (q *QueueLinkedList) Enqueue(s string) {
	n := stringNode {
		item: s,
	}

	oldLast := q.last;
	q.last = &n;
	if q.first == nil {
		q.first = &n;
	} else {
		oldLast.next = &n;
	}

	q.n += 1;
}

func (q *QueueLinkedList) Dequeue() string {
	n := q.first;
	q.first = n.next;
	if q.first == nil {
		q.last = nil;
	}
	q.n -= 1;
	return n.item;
}

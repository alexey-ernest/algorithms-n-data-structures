package ch4

type node struct {
	next *node
	val int
}

type Queue struct {
	first *node
	last *node
	size int
}

func (this *Queue) Enqueue(val int) {
	x := &node {
		val: val,
	}
	if this.last != nil {
		this.last.next = x
		this.last = x
	}
	if this.first == nil {
		this.first = x
		this.last = x
	}
	this.size++
}

func (this *Queue) Size() int {
	return this.size
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) Dequeue() int {
	if this.size == 0 {
		panic("the queue is empty")
	}

	x := this.first
	this.first = x.next
	this.size--

	if this.size == 0 {
		this.last = nil
	}

	return x.val
}
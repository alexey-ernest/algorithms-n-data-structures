package ch1

import(
	"testing"
	"fmt"
)

func TestQueueDequeueLinkedList(t *testing.T) {
	var a [10]string;
	var queue QueueLinkedList;
	for i := 0; i < 10; i+=1 {
		a[i] = fmt.Sprintf("%s", string('a' + i));
		queue.Enqueue(a[i]);
	}

	var b [10]string;
	for i := 0; i < 10; i +=1 {
		b[i] = queue.Dequeue();
	}

	if a != b {
		t.Errorf("%q != %q", a, b);
	}
}
package ch1

import "fmt"

func JosephusEliminateOrder(n, m int) []string {
	// init queue
	var queue QueueLinkedList;
	for i := 0; i < n; i+=1 {
		queue.Enqueue(fmt.Sprintf("%d", i));
	}

	// building the order
	var res []string;
	counter := 1;
	for !queue.IsEmpty() {
		el := queue.Dequeue();
		
		if counter != m {
			queue.Enqueue(el);
		} else {
			res = append(res, el);
			counter = 0;
		}

		counter += 1;
	}

	return res;
}
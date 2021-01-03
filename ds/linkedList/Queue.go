package linkedList

import "testing"

// first in first out
type Queue struct {
	data []interface{}
	size int
}

type QueueOperations interface {
	add(interface{})
	//Inserts the specified element into this queue if it is possible to do so immediately without violating capacity restrictions, returning true upon success and throwing an IllegalStateException if no space is currently available.
	element()
	//Retrieves, but does not remove, the head of this queue.
	peek() interface{}
	//Retrieves, but does not remove, the head of this queue, or returns null if this queue is empty.
	remove()
	//Retrieves and removes the head of this queue.
}

func (queue *Queue) add(val interface{}) {
	queue.data = append(queue.data, val)
	queue.size++
}

func (queue *Queue) element() {

}

func (queue *Queue) peek() interface{} {
	return queue.data[0]
}

func (queue *Queue) remove() {
	if queue.size > 0 {
		queue.data = queue.data[:queue.size-1]
		queue.size--
	}
}

func TestQueue(t *testing.T) {

}

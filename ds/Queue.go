package ds

// first in first out
type Queue struct {
}

type QueueOperations interface {
	add()
	//Inserts the specified element into this queue if it is possible to do so immediately without violating capacity restrictions, returning true upon success and throwing an IllegalStateException if no space is currently available.
	element()
	//Retrieves, but does not remove, the head of this queue.
	peek()
	//Retrieves, but does not remove, the head of this queue, or returns null if this queue is empty.
	remove()
	//Retrieves and removes the head of this queue.
}

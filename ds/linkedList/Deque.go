package linkedList

type Deque struct {
	data []interface{}
	size int
}

type DequeOperations interface {
	add()
	//Inserts the specified element into the queue represented by this deque (in other words, at the tail of this deque) if it is possible to do so immediately without violating capacity restrictions, returning true upon success and throwing an IllegalStateException if no space is currently available.
	addAll()
	//Adds all of the elements in the specified collection at the end of this deque, as if by calling addLast(E) on each one, in the order that they are returned by the collection's iterator.
	addFirst()
	//Inserts the specified element at the front of this deque if it is possible to do so immediately without violating capacity restrictions, throwing an IllegalStateException if no space is currently available.
	addLast()
	//Inserts the specified element at the end of this deque if it is possible to do so immediately without violating capacity restrictions, throwing an IllegalStateException if no space is currently available.
	contains()
	//Returns true if this deque contains the specified element.
	descendingIterator()
	//Returns an iterator over the elements in this deque in reverse sequential order.
	element()
	//Retrieves, but does not remove, the head of the queue represented by this deque (in other words, the first element of this deque).
	getFirst()
	//Retrieves, but does not remove, the first element of this deque.
	getLast()
	//Retrieves, but does not remove, the last element of this deque.
	iterator()
	//Returns an iterator over the elements in this deque in proper sequence.
	offer()
	//Inserts the specified element into the queue represented by this deque (in other words, at the tail of this deque) if it is possible to do so immediately without violating capacity restrictions, returning true upon success and false if no space is currently available.
	offerFirst()
	//Inserts the specified element at the front of this deque unless it would violate capacity restrictions.
	offerLast()
	//Inserts the specified element at the end of this deque unless it would violate capacity restrictions.
	peek()
	//Retrieves, but does not remove, the head of the queue represented by this deque (in other words, the first element of this deque), or returns null if this deque is empty.
	peekFirst()
	//Retrieves, but does not remove, the first element of this deque, or returns null if this deque is empty.
	peekLast()
	//Retrieves, but does not remove, the last element of this deque, or returns null if this deque is empty.
	poll()
	//Retrieves and removes the head of the queue represented by this deque (in other words, the first element of this deque), or returns null if this deque is empty.
	pollFirst()
	//Retrieves and removes the first element of this deque, or returns null if this deque is empty.
	pollLast()
	//Retrieves and removes the last element of this deque, or returns null if this deque is empty.
	pop()
	//Pops an element from the stack represented by this deque.
	push()
	//Pushes an element onto the stack represented by this deque (in other words, at the head of this deque) if it is possible to do so immediately without violating capacity restrictions, throwing an IllegalStateException if no space is currently available.
	remove()
	//Retrieves and removes the head of the queue represented by this deque (in other words, the first element of this deque).
	removeOne()
	//Removes the first occurrence of the specified element from this deque.
	removeFirst()
	//Retrieves and removes the first element of this deque.
	removeFirstOccurrence()
	//Removes the first occurrence of the specified element from this deque.
	removeLast()
	//Retrieves and removes the last element of this deque.
	removeLastOccurrence()
	//Removes the last occurrence of the specified element from this deque.
	size()
	//Returns the number of elements in this deque.
}

type DequeMyOperation interface {
	Prepend(interface{})
	Append(interface{})
	Pop() interface{}
	PeekFirst() interface{}
	PeekLast() interface{}
	IsEmpty() bool
	GetSize() int
	Search(interface{}) int
}

func (deque *Deque) Append(val interface{}) {
	deque.data = append(deque.data, val)
	deque.size++
}

func (deque *Deque) Prepend(val interface{}) {
	//deque.data = append([]interface{}{val}, deque.data)
	deque.data = append([]interface{}{val}, deque.data...)
	deque.size++
}

func (deque *Deque) Pop() interface{} {
	if deque.size > 0 {
		res := deque.data[deque.size-1]
		deque.data = deque.data[:deque.size-1]
		deque.size--
		return res
	}
	return -1
}

func (deque *Deque) PeekFirst() interface{} {
	return deque.data[0]
}
func (deque *Deque) PeekLast() interface{} {
	return deque.data[deque.size-1]
}
func (deque *Deque) IsEmpty() bool {
	return deque.size == 0
}
func (deque *Deque) GetSize() int {
	return deque.size
}
func (deque *Deque) Search(val interface{}) int {
	for k, v := range deque.data {
		if v == val {
			return k
		}
	}
	return -1
}

package ds

type Queue struct {
}

type QueueOperations interface {
	front()
	back()
	pop()
	push()
	empty()
	size()
}

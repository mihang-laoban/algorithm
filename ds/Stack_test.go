package ds

// first in last out
type Stack struct {
}

type StackOperations interface {
	push()
	pop()
	peek()
	empty()
	search()
}

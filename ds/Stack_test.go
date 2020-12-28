package ds

type Stack struct {
}

type StackOperations interface {
	pop()
	top()
	size()
	empty()
	push()
}

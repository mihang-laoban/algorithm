package linkedList

import (
	"fmt"
)

// first in last out

type Stack struct {
	data []interface{}
	size int
}

type StackOperations interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Empty() bool
	Search(interface{}) int
}

func (stack *Stack) Push(val interface{}) {
	stack.data = append(stack.data, val)
	stack.size++
}

func (stack *Stack) Pop() (last interface{}) {
	if stack.size > 0 {
		last = stack.data[stack.size-1]
		stack.data = stack.data[:stack.size-1]
		stack.size--
		return
	}
	return nil
}

func (stack *Stack) Peek() interface{} {
	return stack.data[stack.size-1]
}

func (stack *Stack) Empty() bool {
	if stack.size == 0 && len(stack.data) == 0 {
		return true
	}
	return false
}

func (stack *Stack) Search(val interface{}) int {
	for k, v := range stack.data {
		if v == val {
			return k
		}
	}
	return -1
}

func TestStack() {
	stack := Stack{}
	fmt.Println(stack.Empty())
	stack.Push(1)
	fmt.Println(stack.Pop())
}

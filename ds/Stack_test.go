package ds

import (
	"fmt"
	"testing"
)

// first in last out

type Stack struct {
	data []interface{}
	size int
}

type StackOperations interface {
	push(interface{})
	pop() interface{}
	peek() interface{}
	empty() bool
	search(interface{}) int
}

func (stack *Stack) push(val interface{}) {
	stack.data = append(stack.data, val)
	stack.size++
}

func (stack *Stack) pop() (last interface{}) {
	if stack.size > 0 {
		last = stack.data[stack.size-1]
		stack.data = stack.data[:stack.size-1]
		stack.size--
		return
	}
	return nil
}

func (stack *Stack) peek() interface{} {
	return stack.data[stack.size-1]
}

func (stack *Stack) empty() bool {
	if stack.size == 0 && len(stack.data) == 0 {
		return true
	}
	return false
}

func (stack *Stack) search(val interface{}) int {
	for k, v := range stack.data {
		if v == val {
			return k
		}
	}
	return -1
}

func TestStack(t *testing.T) {
	stack := Stack{}
	fmt.Println(stack.empty())
	stack.push(1)
	fmt.Println(stack.pop())
}

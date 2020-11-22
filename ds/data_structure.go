package ds

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) Insert(x int) {
	if linkedList.head == nil {
		linkedList.head = &Node{Val: x}
		return
	}
	tmp := linkedList.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: x}
}

func (linkedList *LinkedList) Display() {
	tmp := linkedList.head
	for tmp.Next != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

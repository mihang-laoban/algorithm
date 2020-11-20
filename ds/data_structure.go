package ds

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	head *Node
	Val  int
}

func (list *LinkedList) Insert(val int) {
	if list.head == nil {
		list.head = &Node{Val: val}
		return
	}
	tmp := list.head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: val}
}

func (list *LinkedList) Show() {
	for list.head != nil {
		fmt.Println(list.head.Val)
		list.head = list.head.Next
	}
}

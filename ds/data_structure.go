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
	list.Val = val
}

func (list *LinkedList) Show() {
	fmt.Println(list.Val)
}

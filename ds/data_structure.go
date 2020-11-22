package ds

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	Head *Node
}

func (linkedList *LinkedList) Insert(x int) {
	if linkedList.Head == nil {
		linkedList.Head = &Node{Val: x}
		return
	}
	tmp := linkedList.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: x}
}

func (linkedList *LinkedList) InnerDisplay() {
	tmp := linkedList.Head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (linkedList *LinkedList) Display() {
	linkedList._display(linkedList.Head)
}

func (linkedList *LinkedList) _display(node *Node) {
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (linkedList *LinkedList) Reverse1() {
	linkedList.Head = linkedList._reverse1(linkedList.Head)
}

func (linkedList *LinkedList) Reverse2() {
	linkedList.Head = linkedList._reverse2(linkedList.Head)
}

func (linkedList *LinkedList) _reverse1(head *Node) *Node {
	cur := head
	var pre *Node
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func (LinkedList *LinkedList) _reverse2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := LinkedList._reverse2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmp
}

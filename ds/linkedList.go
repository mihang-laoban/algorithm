package ds

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	head   *Node
	length int
}

func (LinkedList *LinkedList) Append(x int) {
	if LinkedList.head == nil {
		LinkedList.length++
		LinkedList.head = &Node{Val: x}
		return
	}
	tmp := LinkedList.head
	LinkedList.length++
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: x}
}

func (LinkedList *LinkedList) Size() int {
	return LinkedList.length
}

func (LinkedList *LinkedList) InnerDisplay() {
	tmp := LinkedList.head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (LinkedList *LinkedList) Display() {
	LinkedList._display(LinkedList.head)
}

func (LinkedList *LinkedList) DisplayR() {
	LinkedList._displayR(LinkedList.head)
	fmt.Println()
}

func (LinkedList *LinkedList) _display(node *Node) {
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (LinkedList *LinkedList) _displayR(node *Node) {
	fmt.Print(node.Val)
	if node.Next != nil {
		LinkedList._displayR(node.Next)
	}
}

func (LinkedList *LinkedList) Reverse1() {
	LinkedList.head = LinkedList._reverse1(LinkedList.head)
}

func (LinkedList *LinkedList) Reverse2() {
	LinkedList.head = LinkedList._reverse2(LinkedList.head)
}

func (LinkedList *LinkedList) _reverse1(head *Node) *Node {
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

func (LinkedList *LinkedList) Prepend(x int) {
	LinkedList.length++
	LinkedList.head = LinkedList._prepend(x)
}

func (LinkedList *LinkedList) _prepend(x int) *Node {
	tmp := &Node{Val: x}
	tmp.Next = LinkedList.head
	return tmp
}

func (LinkedList *LinkedList) IsEmpty() bool {
	if LinkedList.head == nil {
		return true
	}
	return false
}

func (LinkedList *LinkedList) _reverse3(head *Node) *Node {
	var pre *Node
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

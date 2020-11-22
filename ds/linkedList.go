package ds

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type LinkedList struct {
	Head *Node
}

func (LinkedList *LinkedList) Append(x int) {
	if LinkedList.Head == nil {
		LinkedList.Head = &Node{Val: x}
		return
	}
	tmp := LinkedList.Head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: x}
}

func (LinkedList *LinkedList) InnerDisplay() {
	tmp := LinkedList.Head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (LinkedList *LinkedList) Display() {
	LinkedList._display(LinkedList.Head)
}

func (LinkedList *LinkedList) _display(node *Node) {
	tmp := node
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (LinkedList *LinkedList) Reverse1() {
	LinkedList.Head = LinkedList._reverse1(LinkedList.Head)
}

func (LinkedList *LinkedList) Reverse2() {
	LinkedList.Head = LinkedList._reverse2(LinkedList.Head)
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

fun
package ds

import "fmt"

func (LinkedList *LinkedList) Append(x int) {
	if LinkedList.Head == nil {
		LinkedList.Length++
		LinkedList.Head = &Node{Val: x}
		return
	}
	tmp := LinkedList.Head
	LinkedList.Length++
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Node{Val: x}
}

func (LinkedList *LinkedList) Size() int {
	return LinkedList.Length
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

func (LinkedList *LinkedList) DisplayR() {
	LinkedList._displayR(LinkedList.Head)
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
	LinkedList.Head = LinkedList._reverse3(LinkedList.Head)
}

func (LinkedList *LinkedList) Reverse2() {
	LinkedList.Head = LinkedList._reverse2(LinkedList.Head)
}

//func (LinkedList *LinkedList) _reverse1(head *Node) *Node {
//}

//func (LinkedList *LinkedList) _reverse1(head *Node) *Node {
//	cur := head
//	var pre *Node
//	for cur != nil {
//		tmp := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = tmp
//	}
//	return pre
//}

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
	LinkedList.Length++
	LinkedList.Head = LinkedList._prepend(x)
}

func (LinkedList *LinkedList) _prepend(x int) *Node {
	tmp := &Node{Val: x}
	tmp.Next = LinkedList.Head
	return tmp
}

func (LinkedList *LinkedList) IsEmpty() bool {
	if LinkedList.Head == nil {
		return true
	}
	return false
}

func (LinkedList *LinkedList) _reverse3(node *Node) *Node {
	cur := node
	var pre *Node
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

package linkedList

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  interface{}
}

type LinkedList struct {
	Head   *Node
	Length int
}

type Operation interface {
	Add(x interface{})
	Size() int
	InnerDisplay()
	Display()
	DisplayR()
	Reverse()
	ReverseR()
	Prepend(interface{})
	IsEmpty() bool
}

func (LinkedList *LinkedList) Add(x interface{})     { LinkedList._append(x) }
func (LinkedList *LinkedList) Size() int             { return LinkedList.Length }
func (LinkedList *LinkedList) InnerDisplay()         { _innerDisplay(LinkedList) }
func (LinkedList *LinkedList) Display()              { LinkedList._display(LinkedList.Head) }
func (LinkedList *LinkedList) DisplayR()             { LinkedList._displayR(LinkedList.Head) }
func (LinkedList *LinkedList) Reverse()              { LinkedList.Head = LinkedList._reverse(LinkedList.Head) }
func (LinkedList *LinkedList) ReverseR()             { LinkedList.Head = LinkedList._reverseR(LinkedList.Head) }
func (LinkedList *LinkedList) Prepend(x interface{}) { LinkedList.Head = LinkedList._prepend(x) }
func (LinkedList *LinkedList) IsEmpty() bool         { return LinkedList.Head == nil }

func (LinkedList *LinkedList) _append(x interface{}) {
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

func _innerDisplay(LinkedList *LinkedList) {
	tmp := LinkedList.Head
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}

func (LinkedList *LinkedList) _display(node *Node) {
	tmp := node
	for tmp != nil {
		fmt.Printf("%d => ", tmp.Val)
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

func (LinkedList *LinkedList) _reverse(head *Node) *Node {
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

func (LinkedList *LinkedList) _reverseR(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := LinkedList._reverseR(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmp
}

func (LinkedList *LinkedList) _prepend(x interface{}) *Node {
	LinkedList.Length++
	tmp := &Node{Val: x}
	tmp.Next = LinkedList.Head
	return tmp
}

func test() {
	node := LinkedList{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(node.IsEmpty())
	for i := 0; i < len(arr); i++ {
		node.Add(arr[i])
	}
	node.Prepend(0)
	node.Display()
	node.Reverse()
	node.DisplayR()
	fmt.Println()
	fmt.Println(node.IsEmpty())
	fmt.Println(node.Size())
}

func IsCircular(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func circle(root *Node) {
	second := root.Next
	for root.Next != nil {
		root = root.Next
	}
	root.Next = second
}

func GetCircleLink() *Node {
	node := LinkedList{}
	arr := []int{3, 2, 0, -4}
	for _, v := range arr {
		node.Add(v)
	}
	circle(node.Head)
	return node.Head
}

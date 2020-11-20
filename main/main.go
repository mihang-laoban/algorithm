package main

import (
	"dp/Questions/DP"
	"dp/ds"
	"fmt"
)

func dpCollection() {
	DP.FindLongestSharedSubSeq()
}

func main() {

	//arr := []int{1, 2, 3, 4, 5}

	test := ds.LinkedList{}
	test.Insert(15)
	test.Show()

	//head := ds.Node{}
	//
	//createLinkedList(&head, arr)
	//
	//printLinkedList(&head)
	//printLinkedListRecur(&head)
	//
	//reversionLinkedList(&head)
}

func reversionLinkedList(node *ds.Node) {

}

func createLinkedList(head *ds.Node, arr []int) {

}

func printLinkedList(head *ds.Node) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Print(tmp.Val, "-")
	}
	fmt.Println()
}

func printLinkedListRecur(head *ds.Node) {
	fmt.Print(head.Val, "-")
	if head.Next != nil {
		printLinkedListRecur(head.Next)
	}
	fmt.Println()
}

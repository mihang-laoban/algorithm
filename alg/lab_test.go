package alg

import (
	linkedList2 "dp/alg/linkedList"
	. "dp/ds/linkedList"
	"fmt"
	"testing"
)

func TestComponentLab(t *testing.T) {
	linkedList()
}

func linkedList() {
	head := ArrayToLinkedList([]int{1, 2, 3, 4})
	linkedList2.ReverseListL()
	//dummy := &ListNode{}
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	fmt.Println(LinkedListToArray(cur))
}

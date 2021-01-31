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
	cur := linkedList2.ReverseListR(head)

	fmt.Println(LinkedListToArray(cur))
}

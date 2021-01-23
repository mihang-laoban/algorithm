package linkedList

import (
	. "dp/ds/linkedList"
	"fmt"
	"testing"
)

func TestIsCircle(t *testing.T) {
	node := GetCircleLink()
	fmt.Println(IsCircular(node))
	fmt.Println(isCircular(node))
	TestLinkedList()
}

func isCircular(node *Node) bool {
	return true
}

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
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMergeList(t *testing.T) {
	l1 := ArrayToLinkedList([]int{1,2,4})
	l2 := ArrayToLinkedList([]int{1,3,4})
	head := MergeLinkedList(l1, l2)
	fmt.Println(LinkedListToArray(head))
}

func MergeLinkedList(l1, l2 *NodeInt) *NodeInt {
	head := &NodeInt{Val: -1}
	tmp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = &NodeInt{Val: l1.Val}
			l1 = l1.Next
		} else {
			tmp.Next = &NodeInt{Val: l2.Val}
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	}else {
		tmp.Next = l1
	}
	return head.Next
}
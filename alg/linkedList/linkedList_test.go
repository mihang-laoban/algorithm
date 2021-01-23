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
	l1 := ArrayToLinkedList([]int{1, 2, 4})
	l2 := ArrayToLinkedList([]int{1, 3, 4})
	//head := MergeLinkedListL(l1, l2)
	head := MergeLinkedListR(l1, l2)
	fmt.Println(LinkedListToArray(head))
}

/*给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSwapPairs(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4})
	//fmt.Println(LinkedListToArray(SwapPairsR(head)))
	fmt.Println(LinkedListToArray(SwapPairsL(head)))
}

func SwapPairsR(head *NodeInt) *NodeInt {
	// 如果头节点或者下一个节点尾空，则无法交换
	if head == nil || head.Next == nil {
		return head
	}
	// 存储第二个节点
	newHead := head.Next
	// 第二个节点指向第三个节点交换以后的结果
	head.Next = SwapPairsR(newHead.Next)
	// 第二个节点成为新的头节点
	newHead.Next = head
	return newHead

	/*	if head == nil || head.Next == nil {
			return head
		}
		third := SwapPairsR(head.Next.Next)
		second := head.Next
		head.Next = third
		second.Next = head
		return second*/
}
func SwapPairsL(head *NodeInt) *NodeInt {
	dummyHead := &NodeInt{head, 0}
	cur := dummyHead
	/*
		c, 1, 2, 3, 4
		f = 1
		s = 2
		c.next = s
		f.next = s.next
		s.next = f
		c = f
	*/
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next
		cur.Next = second
		first.Next = second.Next
		second.Next = first
		cur = first
	}
	return dummyHead.Next
}

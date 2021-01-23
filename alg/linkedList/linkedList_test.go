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

/*给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5

说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func Test(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(LinkedListToArray(ReverseKGroup(head, 3)))
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 链表头节点
	dum := &ListNode{Val: -1}
	// 表头指向入参头节点
	dum.Next = head
	// 初始化起点和终点
	start, end := dum, dum
	for end.Next != nil {
		// 遍历到第K个元素
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果到达最后一组没有遍历完则不反转
		if end == nil {
			break
		}
		// 记录起点和下一个起点
		cur, next := start.Next, end.Next
		// 断开链接，设置反转终点
		end.Next = nil
		// 起点指向反转以后的第一个节点
		start.Next = ReverseListL(cur)
		// 重新链接，此时的start已经上一组的终点
		cur.Next = next
		// 重新设置起点
		start, end = cur, cur
	}
	return dum.Next
}

func ReverseListL(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

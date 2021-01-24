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

/*反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestReverseLinkedList(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	//fmt.Println(LinkedListToArray(ReverseListR(head)))
	//fmt.Println(LinkedListToArray(ReverseKListR(head, 3)))
	//fmt.Println(LinkedListToArray(ReverseListBetweenR(head, 2, 4)))
	fmt.Println(LinkedListToArray(ReverseListBetweenL(head, 2, 4)))
	//fmt.Println(LinkedListToArray(ReverseListBetweenL2(head, 2, 4)))
}

func ReverseListBetweenL(head *ListNode, m, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	guard, point := dummy, dummy.Next
	for i := 0; i < m-1; i++ {
		guard, point = guard.Next, point.Next
	}
	for i := 0; i < n-m; i++ {
		removed := point.Next
		point.Next = point.Next.Next
		removed.Next = guard.Next
		guard.Next = removed
	}
	return dummy.Next
}

func ReverseListBetweenL2(head *ListNode, m, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	node := dummy
	//找到需要反转的那一段的上一个节点。
	for i := 1; i < m; i++ {
		node = node.Next
	}
	//node.next就是需要反转的这段的起点。
	cur := node.Next
	var tmp, pre *ListNode
	//反转m到n这一段
	for i := m; i <= n; i++ {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	//将反转的起点的next指向next。
	node.Next.Next = tmp
	//需要反转的那一段的上一个节点的next节点指向反转后链表的头结点
	node.Next = pre
	return dummy.Next
}

func ReverseListBetweenR(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return ReverseKListR(head, n)
	}
	head.Next = ReverseListBetweenR(head.Next, m-1, n-1)
	return head
}

func ReverseKListR(head *ListNode, n int) *ListNode {
	var (
		successor      *ListNode
		_reverseFirstK func(*ListNode, int) *ListNode
	)

	_reverseFirstK = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			successor = head.Next
			return head
		}
		last := _reverseFirstK(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return last
	}

	return _reverseFirstK(head, n)
}

/*
用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。

你可以假设这个整数除了 0 本身，没有任何前导的 0。

这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。

示例:

输入: [1,2,3]
输出: [1,2,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPlusOne(t *testing.T) {
	head := ArrayToLinkedList([]int{9})
	fmt.Println(LinkedListToArray(PlusOne(head)))
}

/*给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
1->4->5,
1->3->4,
2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMergeKLists(t *testing.T) {
	head1 := ArrayToLinkedList([]int{1, 4, 5})
	head2 := ArrayToLinkedList([]int{1, 3, 4})
	head3 := ArrayToLinkedList([]int{2, 6})
	//fmt.Println(LinkedListToArray(MergeKLists([]*ListNode{head1, head2, head3})))
	fmt.Println(LinkedListToArray(MergeKListsPriorityQueue([]*ListNode{head1, head2, head3})))
}

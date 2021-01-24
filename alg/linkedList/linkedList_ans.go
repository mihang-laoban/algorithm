package linkedList

import (
	. "dp/ds/linkedList"
)

func MergeLinkedListL(l1, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	tmp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			tmp.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return head.Next
}

func MergeLinkedListR(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeLinkedListR(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeLinkedListR(l1, l2.Next)
		return l2
	}
}

func SwapPairsR(head *ListNode) *ListNode {
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
func SwapPairsL(head *ListNode) *ListNode {
	dummyHead := &ListNode{head, -1}
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
		start.Next = ReverseListR(cur)
		// 重新链接，此时的start已经上一组的终点
		cur.Next = next
		// 重新设置起点
		start, end = cur, cur
	}
	return dum.Next
}

func ReverseListR(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := ReverseListR(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
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

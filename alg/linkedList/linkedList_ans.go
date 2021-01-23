package linkedList

import . "dp/ds/linkedList"

func MergeLinkedListL(l1, l2 *NodeInt) *NodeInt {
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
	} else {
		tmp.Next = l1
	}
	return head.Next
}

func MergeLinkedListR(l1, l2 *NodeInt) *NodeInt {
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

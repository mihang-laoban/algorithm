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

package linkedList

import (
	"container/heap"
	. "dp/ds/linkedList"
)

func MergeLinkedListL(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else {
		tmp.Next = l1
	}
	return dummy.Next
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
	dummyHead := &ListNode{Next: head, Val: -1}
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

//"1299","1229","1922",普通情况，末尾元素+1，链表长度不变，返回头节点
//“9”，“99”，“999”这种特殊情况，返回 哑节点，也就是没有移动一步的 slow 节点
func PlusOne(head *ListNode) *ListNode {
	fast, slow := head, &ListNode{Val: 0}
	slow.Next = head
	//2.遍历链表
	for fast != nil {
		if fast.Val != 9 {
			slow = fast
		}
		fast = fast.Next
	}
	//3.末位加1
	slow.Val++
	cur := slow.Next
	for cur != nil {
		cur.Val = 0
		cur = cur.Next
	}
	if slow.Next == head {
		return slow
	}
	return head
}

func MergeKLists(lists []*ListNode) *ListNode {
	var mergeKLists func(int, int) *ListNode
	mergeKLists = func(low, hight int) *ListNode {
		if low == hight {
			return lists[low]
		}
		if low > hight {
			return nil
		}
		mid := (low + hight) >> 1
		return MergeLinkedListL(mergeKLists(low, mid), mergeKLists(mid+1, hight))
	}
	return mergeKLists(0, len(lists)-1)
}

/*本题考查最小堆的用法 最小堆里面的每个元素可以是一个结构体，只要正确实现了Less方法即可*/
func MergeKListsPriorityQueue(lists []*ListNode) *ListNode {
	pq := &minHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(pq, lists[i])
		}
	}
	heap.Init(pq)
	head := &ListNode{}
	cur := head
	for pq.Len() > 0 {
		curNode := heap.Pop(pq).(*ListNode)
		if curNode.Next != nil {
			heap.Push(pq, curNode.Next)
		}
		cur.Next = curNode
		cur = cur.Next
	}
	return head.Next
}

type minHeap []*ListNode //由链表组成的最小堆
func (h *minHeap) Len() int {
	return len(*h)
}
func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *minHeap) Pop() interface{} {
	headNode := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return headNode
}
func (h *minHeap) Push(node interface{}) {
	newNode := node.(*ListNode)
	*h = append(*h, newNode)
}

func SortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for i := head; i != nil; i = i.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		pre, cur := dummy, dummy.Next
		for cur != nil {
			l1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			l2 := cur.Next
			cur.Next = nil
			cur = l2
			// ----------------------|处理奇数个元素的情况|-------------------------
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			cur = next

			pre.Next = MergeLinkedListL(l1, l2)

			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return dummy.Next
}

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := MiddleNode(head)
	l1, l2 := head, mid.Next
	mid.Next = nil // 断开链表1
	l2 = ReverseListL(l2)
	reorderList(l1, l2)
}

func reorderList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		// 存出两个链表头节点的下一个节点
		tmp1 := l1.Next
		tmp2 := l2.Next

		l1.Next = l2 // 交错
		l1 = tmp1    // 向后移动一位

		l2.Next = l1 // 交错
		l2 = tmp2    // 向后移动一位
	}
}

func Partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallTmp, largeTmp := small, large
	for head != nil {
		if head.Val < x {
			smallTmp.Next = head
			smallTmp = smallTmp.Next
		} else {
			largeTmp.Next = head
			largeTmp = largeTmp.Next
		}
		head = head.Next
	}
	largeTmp.Next, smallTmp.Next = nil, large.Next
	return small.Next
}

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy.Next
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := []int{}, []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var ans *ListNode
	var a, b, carry int
	for len(s1) > 0 || len(s2) > 0 || carry != 0 {
		if len(s1) == 0 {
			a = 0
		} else {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) == 0 {
			b = 0
		} else {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		ans = &ListNode{Next: ans, Val: cur}
	}
	return ans
}

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}

func KthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func DeleteDuplicates(head *ListNode) *ListNode {
	cur := head
	// cur != nil 处理头节点为空的情况
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

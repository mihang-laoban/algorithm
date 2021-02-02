package linkedList

import (
	"container/heap"
	. "dp/ds/linkedList"
	. "dp/ds/tree"
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
	last := ReverseListR(head.Next) // 递推到最后把尾节点一直向上抛出，回归时的当前头节点就是上一个节点
	head.Next.Next = head           // 头节点的下一个节点向前指向头节点形成循环链表(目标达成)
	head.Next = nil                 // 头节点断开
	return last
}

/*
	关键步骤：
	1.需要当前节点指向前面一个节点，由于要继续遍历，需要保存当前节点的后继节点，所以我们需要先记录 tmp := cur.Next
	2.当前节点可以指向前驱节点了, cur.Next = pre
	3.此时，前驱节点可以更新为当前节点了 pre = cur
	4.继续看后面的元素，当前节点指向之前预存的后继节点cur = tmp
*/
func ReverseListL(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next // 存出后续节点 tmp = 2 > 3 > 4
		cur.Next = pre  // 头节点的后继节点指向前驱节点 cur(1) > pre(nil)
		pre = cur       // 前驱节点替换为当前节点 pre = 1 > nil
		cur = tmp       // 当前节点替换为存储的后续节点 cur = 2 > 3 > 4
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

func DeleteDuplicatesIIR(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		if head != nil {
			return DeleteDuplicatesIIR(head.Next)
		} else {
			return nil
		}
	} else {
		head.Next = DeleteDuplicatesIIR(head.Next)
		return head
	}
}

func IsPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var pre, cur *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		cur, slow, fast = slow, slow.Next, fast.Next.Next
		cur.Next = pre
		pre = cur
	}
	if fast != nil { // 处理奇数的情况
		slow = slow.Next
	}
	for cur != nil && slow != nil {
		if cur.Val != slow.Val {
			return false
		}
		cur, slow = cur.Next, slow.Next
	}
	return true
}

func IsPalindrome2(head *ListNode) interface{} {
	tmp := head
	var check func(*ListNode) bool
	check = func(cur *ListNode) bool {
		if cur != nil {
			if !check(cur.Next) {
				return false
			}
			if cur.Val != tmp.Val {
				return false
			}
			tmp = tmp.Next
		}
		return true
	}
	return check(head)
}

func IsPalindrome1(head *ListNode) bool {
	size, tmp, cur := 0, []int{}, head
	for cur != nil {
		tmp, cur = append(tmp, cur.Val), cur.Next
		size++
	}
	for i := 0; i < size; i++ {
		if tmp[i] != tmp[size-1-i] {
			return false
		}
	}
	return true
}

func DeleteDuplicatesIIL(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil {
		// 如果当前节点和下一个节点值相等则向后移动一位
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		// 前置节点如果和当前节点是同一个节点则一起向后移动一位
		if pre.Next == cur {
			pre = pre.Next
		} else {
			// 否则指向已经下一位和当前位值不相等的位置
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// nlogn - logn
func SortedListToBST1(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var preSlow *ListNode
	// 找到中间节点
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 根节点为中间值
	root := &TreeNode{Val: slow.Val}
	if preSlow != nil {
		// 只传入前半段，即左子树
		preSlow.Next = nil
		root.Left = SortedListToBST1(head)
	}
	root.Right = SortedListToBST1(slow.Next)
	return root
}

/*
   0
  / \
-10  5
   \  \
    3  9
*/

// n - n
func SortedListToBST2(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var build func(int, int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		root := &TreeNode{Val: nums[mid]}
		root.Left = build(start, mid-1)
		root.Right = build(mid+1, end)
		return root
	}
	return build(0, len(nums)-1)
}

// n - logn
func SortedListToBST3(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	cur, length := head, 0
	for head != nil {
		length++
		head = head.Next
	}
	var buildBST func(int, int) *TreeNode
	buildBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		// 只看左半
		left := buildBST(start, mid-1)
		root := &TreeNode{Val: cur.Val}
		cur = cur.Next
		root.Left = left
		// 只看右半
		root.Right = buildBST(mid+1, end)
		return root
	}
	return buildBST(0, length-1)
}

func GetRandomList(nums [][]interface{}) *RandomListNode {
	dummy := &RandomListNode{}
	pre := dummy
	nodeMap := map[int]*RandomListNode{}
	arr := []*RandomListNode{}
	for i := 0; i < len(nums); i++ {
		cur := &RandomListNode{
			Val:    nums[i][0].(int),
			Next:   nil,
			Random: nil,
		}
		arr = append(arr, cur)
		nodeMap[i] = cur
		pre.Next = cur
		pre = pre.Next
	}
	for i := 0; i < len(nums); i++ {
		if nums[i][1] != nil {
			arr[i].Random = nodeMap[nums[i][1].(int)]
		}
	}
	return dummy.Next
}

func CopyRandomListNew(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	// 复制链表
	cur := head
	for cur != nil {
		tmp := &RandomListNode{Val: cur.Val}
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
	// 处理random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 拆分链表
	first, second, tmp := head, head.Next, head.Next
	for first != nil {
		first.Next = first.Next.Next
		if second.Next != nil {
			second.Next = second.Next.Next
		}
		first, second = first.Next, second.Next
	}
	return tmp
}

func CopyRandomList(head *RandomListNode) *RandomListNode {
	visited := map[*RandomListNode]*RandomListNode{}
	var copyNodesR func(*RandomListNode) *RandomListNode
	copyNodesR = func(head *RandomListNode) *RandomListNode {
		if head == nil {
			return head
		}
		if v, ok := visited[head]; ok {
			if ok {
				return v
			}
		}
		node := &RandomListNode{Val: head.Val}
		visited[head] = node
		node.Next = copyNodesR(head.Next)
		node.Random = copyNodesR(head.Random)
		return node
	}
	return copyNodesR(head)
}

func DuplicateLinkedList(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		tmp := &ListNode{Val: cur.Val}
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
	return head
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	lastSort, cur := head, head.Next
	for cur != nil {
		if lastSort.Val <= cur.Val {
			lastSort = lastSort.Next
		} else {
			pre := dummy
			// 从头找到第一个大于当前元素的元素
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}
			lastSort.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
		}
		cur = lastSort.Next
	}
	return dummy.Next
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func SplitListToParts(root *ListNode, k int) []*ListNode {
	cur, size := root, 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	width, rem, res := size/k, size%k, []*ListNode{}
	cur = root
	for i := 0; i < k; i++ {
		head, seq := cur, 0
		if i < rem {
			seq = 0
		} else {
			seq = 1
		}
		for j := 0; j < width-seq; j++ {
			if cur != nil {
				cur = cur.Next
			}
		}
		if cur != nil {
			pre := cur
			cur = cur.Next
			pre.Next = nil
		}
		res = append(res, head)
	}
	return res
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length, last := 1, head
	for last.Next != nil {
		length++
		last = last.Next
	}
	last.Next = head
	newHead := head
	for i := 1; i < length-k%length; i++ {
		newHead = newHead.Next
	}
	head, newHead.Next = newHead.Next, nil
	return head
}

func RemoveDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	record := map[int]bool{}
	record[head.Val] = true
	cur := head
	for cur.Next != nil {
		// 如果存在，删除当前联表
		if record[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			// 如果当前值不存在，记录
			record[cur.Next.Val] = true
			cur = cur.Next
		}
	}
	return head
}

func RemoveDuplicateNodes2(head *ListNode) *ListNode {
	a, pre := map[int]bool{}, head
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := a[cur.Val]; ok {
			pre.Next = cur.Next
		} else {
			a[cur.Val] = true
			pre = cur
		}
	}
	return head
}

type PhoneDirectory struct {
	nums         []int
	used         []bool
	total, index int
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func DirConstructor(maxNumbers int) PhoneDirectory {
	dir := PhoneDirectory{total: maxNumbers}
	dir.used = make([]bool, maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		dir.nums = append(dir.nums, i)
	}
	return dir
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	if this.index >= this.total {
		return -1
	}
	val := this.nums[this.index]
	this.index++
	this.used[val] = true
	return val
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	if number < this.total {
		return !this.used[number]
	}
	return false
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
	if number < this.total && this.used[number] {
		this.used[number] = false
		this.index--
		this.nums[this.index] = number
	}
}

func GetDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res<<1 + head.Val
		head = head.Next
	}
	return res
}

func NumComponents(head *ListNode, G []int) int {
	ans, cur, record := 0, head, make(map[int]bool, len(G))
	for _, value := range G {
		record[value] = true
	}
	for cur != nil {
		// 当前值在记录中，下一个值不存在或者或者下一个存在但并不在记录中，则为一个组件
		if record[cur.Val] && (cur.Next == nil || !record[cur.Next.Val]) {
			ans++
		}
		cur = cur.Next
	}
	return ans
}

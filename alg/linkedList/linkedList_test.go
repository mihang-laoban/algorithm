package linkedList

import (
	. "dp/ds/linkedList"
	. "dp/ds/tree"
	"fmt"
	"testing"
)

func TestIsCircle(t *testing.T) {
	node := GetCircleLink()
	fmt.Println(IsCircular(node))
	TestLinkedList()
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

/*给你一个链表，每 index 个节点一组进行翻转，请你返回翻转后的链表。
index 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 index 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 index = 2 时，应当返回: 2->1->4->3->5
当 index = 3 时，应当返回: 3->2->1->4->5

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
	fmt.Println(LinkedListToArray(MergeKLists2([]*ListNode{head1, head2, head3})))
	//fmt.Println(LinkedListToArray(MergeKListsPriorityQueue([]*ListNode{head1, head2, head3})))
}

func MergeKLists2(lists []*ListNode) *ListNode {
	size := len(lists)
	head := &ListNode{}
	if size == 0 {
		return head.Next
	}
	for i := size - 1; i >= 0; i-- {
		if lists[i] == nil {
			size--
			lists[i] = lists[size]
		}
		heapify(lists, i, size)
	}
	tail := head
	for size > 0 {
		tail.Next = lists[0]
		tail = tail.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			size--
			if size == 0 {
				break
			}
			lists[0] = lists[size]
		}
		heapify(lists, 0, size)
	}
	return head.Next
}

func heapify(lists []*ListNode, i, size int) {
	temp := lists[i]
	for k := i<<1 + 1; k < size; k = k<<1 + 1 {
		if k+1 < size && lists[k].Val > lists[k+1].Val {
			k++
		}
		if temp.Val > lists[k].Val {
			lists[i] = lists[k]
			i = k
		}
	}
	lists[i] = temp
}

/*给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSortList(t *testing.T) {
	head := ArrayToLinkedList([]int{4, 2, 1, 3, 5})
	fmt.Println(LinkedListToArray(SortList(head)))
}

/*给定一个头结点为 head 的非空单链表，返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。

示例 1：
输入：[1,2,3,4,5]
输出：此列表中的结点 3 (序列化形式：[3,4,5])
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.

示例 2：
输入：[1,2,3,4,5,6]
输出：此列表中的结点 4 (序列化形式：[4,5,6])
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。*/

func TestMiddleNode(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(MiddleNode(head).Val)
}

/*给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestReorderList(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	ReorderList(head)
	fmt.Println(LinkedListToArray(head))
}

/*给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。

示例：

输入：head = 1->4->3->2->5->2, x = 3
输出：1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPartition(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(LinkedListToArray(Partition(head, 3)))
}

/*请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
现有一个链表 -- head = [4,5,1,9]，它可以表示为:

示例 1：
输入：head = [4,5,1,9], node = 5
输出：[4,1,9]
解释：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.

示例 2：
输入：head = [4,5,1,9], node = 1
输出：[4,5,9]
解释：给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestDeleteNode(t *testing.T) {
	head := ArrayToLinkedList([]int{4, 5, 1, 9})
	target := SearchNode(head, 5)
	DeleteNode(target)
	fmt.Println(LinkedListToArray(head))
}

/*给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
进阶：你能尝试使用一趟扫描实现吗？

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestRemoveNthFromEnd(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(LinkedListToArray(RemoveNthFromEnd(head, 2)))
}

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：
输入：(7 -> 2 -> 4 -> 2) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 6
*/
func TestAddTwoNumbers(t *testing.T) {
	head1 := ArrayToLinkedList([]int{7, 2, 4, 2})
	head2 := ArrayToLinkedList([]int{5, 6, 4})
	fmt.Println(LinkedListToArray(AddTwoNumbers(head1, head2)))
}

/*
删除链表中等于给定值 val 的所有节点。

示例:
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

1,1,1
[]
*/

func TestRemoveElements(t *testing.T) {
	//head := ArrayToLinkedList([]int{6, 1, 2, 6, 3, 4, 5, 6})
	//head := ArrayToLinkedList([]int{1, 1, 1})
	head := ArrayToLinkedList([]int{1})
	fmt.Println(LinkedListToArray(RemoveElements(head, 1)))
}

/*实现一种算法，找出单向链表中倒数第 index 个节点。返回该节点的值。

注意：本题相对原题稍作改动
示例：

输入： 1->2->3->4->5 和 index = 2
输出： 4
说明：

给定的 index 保证是有效的。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestKthToLast(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(KthToLast(head, 2))
}

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
示例 1:
输入: 1->1->2
输出: 1->2

示例 2:
输入: 1->1->2->3->3
输出: 1->2->3
*/
func TestDeleteDuplicates(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 1, 2, 3, 3})
	fmt.Println(LinkedListToArray(DeleteDuplicates(head)))
}

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestIsPalindrome(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 8, 3, 2, 1})
	//head := ArrayToLinkedList([]int{1, 0, 1})
	fmt.Println(IsPalindrome1(head))
	fmt.Println(IsPalindrome2(head))
	fmt.Println(IsPalindrome3(head))
}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5

示例 2:
输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestDeleteDuplicatesII(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	//fmt.Println(LinkedListToArray(DeleteDuplicatesIIL(head)))
	fmt.Println(LinkedListToArray(DeleteDuplicatesIIR(head)))
}

/*

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, nil, 5], 它可以表示下面这个高度平衡二叉搜索树：

     0
    / \
  -3   9
  /   /
-10  5

   0
  / \
-10  5
   \  \
    3  9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func TestTree(t *testing.T) {
	head := ArrayToLinkedList([]int{-10, -3, 0, 5, 9})
	fmt.Println(TreeToArray(SortedListToBST1(head)))
	head = ArrayToLinkedList([]int{-10, -3, 0, 5, 9})
	fmt.Println(TreeToArray(SortedListToBST2(head)))
	head = ArrayToLinkedList([]int{-10, -3, 0, 5, 9})
	fmt.Println(TreeToArray(SortedListToBST3(head)))
}

/*给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。
我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
val：一个表示 RandomListNode.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  nil 。

示例 1：

输入：head = [[7,nil],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,nil],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：

输入：head = [[3,nil],[3,0],[3,nil]]
输出：[[3,nil],[3,0],[3,nil]]
示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 nil。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/copy-list-with-random-pointer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestCopyRandomList(t *testing.T) {
	head := GetRandomList([][]interface{}{
		[]interface{}{7, nil},
		[]interface{}{13, 0},
		[]interface{}{11, 4},
		[]interface{}{10, 2},
		[]interface{}{1, 0},
	})
	newHead1 := CopyRandomList(head)
	newHead2 := CopyRandomListNew(head)
	fmt.Println(newHead1, newHead2)
}

/*对链表进行插入排序。

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：
插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

示例 1：
输入: 4->2->1->3
输出: 1->2->3->4

示例 2：
输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insertion-sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestInsertionSortList(t *testing.T) {
	head := ArrayToLinkedList([]int{-1, 5, 3, 4, 0})
	fmt.Println(LinkedListToArray(InsertionSortList(head)))
}

/*给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。*/

func TestOddEvenList(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(LinkedListToArray(OddEvenList(head)))
}

/*给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 nil。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。

进阶：
你是否可以使用 O(1) 空间解决此题？

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 nil
解释：链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestCir(t *testing.T) {
	head := ArrayToLinkedList([]int{3, 2, 0, -1})
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head.Next
	fmt.Println(DetectCycle(head))
}

/*编写一个程序，找到两个单链表相交的起始节点。

示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

示例 2：
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。


示例 3：
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：nil
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 nil。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestGetIntersectionNode(t *testing.T) {
	head := ArrayToLinkedList([]int{8, 4, 5})
	a := ArrayToLinkedList([]int{4, 1})
	b := ArrayToLinkedList([]int{5, 0, 1})
	a.Next.Next = head
	b.Next.Next.Next = head
	fmt.Println(GetIntersectionNode(a, b).Val)
}

/*给定一个链表，旋转链表，将链表每个节点向右移动 index 个位置，其中 index 是非负数。

示例 1:
输入: 1->2->3->4->5->NULL, index = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

示例 2:
输入: 0->1->2->NULL, index = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestRotateRight(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(LinkedListToArray(RotateRight(head, 2)))
}

/*给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
返回一个符合上述规则的链表的列表。

举例： 1->2->3->4, index = 5 // 5 结果 [ [1], [2], [3], [4], nil ]

示例 1：
输入:
root = [1, 2, 3], index = 5
输出: [[1],[2],[3],[],[]]
解释:
输入输出各部分都应该是链表，而不是数组。
例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且 root.next.next.next = nil。
第一个输出 output[0] 是 output[0].val = 1, output[0].next = nil。
最后一个元素 output[4] 为 nil, 它代表了最后一个部分为空链表。

示例 2：
输入:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], index = 3
输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
解释:
输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-linked-list-in-parts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSplitListToParts(t *testing.T) {
	head1 := ArrayToLinkedList([]int{1, 2, 3})
	heads1 := SplitListToParts(head1, 5)
	for i := 0; i < len(heads1); i++ {
		fmt.Print(LinkedListToArray(heads1[i]), " ")
	}
	head2 := ArrayToLinkedList([]int{1, 2, 3, 4, 5, 6, 7})
	heads2 := SplitListToParts(head2, 3)
	fmt.Println()
	for i := 0; i < len(heads2); i++ {
		fmt.Print(LinkedListToArray(heads2[i]), " ")
	}
}

/*编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

输入：[1, 2, 3, 3, 2, 1]
输出：[1, 2, 3]
示例2:

输入：[1, 1, 1, 1, 2]
输出：[1, 2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestRemoveDuplicateNodes(t *testing.T) {
	head := ArrayToLinkedList([]int{1, 2, 3, 3, 2, 1})
	fmt.Println(LinkedListToArray(RemoveDuplicateNodes2(head)))
}

/*给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
请你返回该链表所表示数字的 十进制值 。

示例 1：
输入：head = [1,0,1]
输出：5
解释：二进制数 (101) 转化为十进制数 (5)

示例 2：
输入：head = [0]
输出：0

示例 3：
输入：head = [1]
输出：1

示例 4：
输入：head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
输出：18880

示例 5：
输入：head = [0,0]
输出：0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestGetDecimalValue(t *testing.T) {
	//head := ArrayToLinkedList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	head := ArrayToLinkedList([]int{1, 0, 1})
	fmt.Println(GetDecimalValue(head))
}

/*设计一个电话目录管理系统，让它支持以下功能：
	get: 分配给用户一个未被使用的电话号码，获取失败请返回 -1
	check: 检查指定的电话号码是否被使用
	release: 释放掉一个电话号码，使其能够重新被分配

示例：
// 初始化电话目录，包括 3 个电话号码：0，1 和 2。
PhoneDirectory2 directory = new PhoneDirectory2(3);

// 可以返回任意未分配的号码，这里我们假设它返回 0。
directory.get();

// 假设，函数返回 1。
directory.get();

// 号码 2 未分配，所以返回为 true。
directory.check(2);

// 返回 2，分配后，只剩一个号码未被分配。
directory.get();

// 此时，号码 2 已经被分配，所以返回 false。
directory.check(2);

// 释放号码 2，将该号码变回未分配状态。
directory.release(2);

// 号码 2 现在是未分配状态，所以返回 true。
directory.check(2);

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-phone-directory
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestPhoneDirectory(t *testing.T) {
	// 初始化电话目录，包括 3 个电话号码：0，1 和 2。
	directory := DirConstructor(3)
	// 可以返回任意未分配的号码，这里我们假设它返回 0。
	fmt.Print(directory.Get(), " ")
	// 假设，函数返回 1。
	fmt.Print(directory.Get(), " ")
	// 号码 2 未分配，所以返回为 true。
	fmt.Print(directory.Check(2), " ")
	// 返回 2，分配后，只剩一个号码未被分配。
	fmt.Print(directory.Get(), " ")
	// 此时，号码 2 已经被分配，所以返回 false。
	fmt.Print(directory.Check(2), " ")
	// 释放号码 2，将该号码变回未分配状态。
	directory.Release(2)
	// 号码 2 现在是未分配状态，所以返回 true。
	fmt.Print(directory.Check(2), " ")
	fmt.Println()
}

/*给定链表头结点 head，该链表上的每个结点都有一个 唯一的整型值 。
同时给定列表 G，该列表是上述链表中整型值的一个子集。
返回列表 G 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。

示例 1：
输入:
head: 0->1->2->3
G = [0, 1, 3]
输出: 2
解释:
链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。

示例 2：
输入:
head: 0->1->2->3->4
G = [0, 3, 1, 4]
输出: 2
解释:
链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-components
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestNumComponents(t *testing.T) {
	head := ArrayToLinkedList([]int{0, 1, 2, 3, 4})
	fmt.Println(NumComponents(head, []int{0, 3, 1, 4}))
}

/*给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。
下图中蓝色边和节点展示了操作后的结果：
请你返回结果链表的头指针。

示例 1：
输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中第三和第四个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。

示例 2：
输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
解释：上图中蓝色的边和节点为答案链表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-in-between-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestMergeInBetween(t *testing.T) {
	head1 := ArrayToLinkedList([]int{0, 1, 2, 3, 4, 5})
	head2 := ArrayToLinkedList([]int{1000000, 1000001, 1000002})
	fmt.Println(LinkedListToArray(MergeInBetween(head1, 3, 4, head2)))
	//head1 := ArrayToLinkedList([]int{0, 1, 2, 3, 4, 5, 6})
	//head2 := ArrayToLinkedList([]int{1000000, 1000001, 1000002, 1000003, 1000004})
	//fmt.Println(LinkedListToArray(MergeInBetween(head1, 2, 5, head2)))
}

/*将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。
对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
特别地，我们希望可以 就地 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。

示例 1：
输入：root = [4,2,5,1,3]
输出：[1,2,3,4,5]
解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。

示例 2：
输入：root = [2,1,3]
输出：[1,2,3]

示例 3：
输入：root = []
输出：[]
解释：输入是空树，所以输出也是空链表。

示例 4：
输入：root = [1]
输出：[1]

    4
   / \
  2   5
 / \
1   3
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestTreeToDoublyList(t *testing.T) {
	nums := []interface{}{4, 2, 5, 1, 3}
	size := len(nums)
	root := ArrayToTree(nums)
	head := TreeToDoublyList(root)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = head.Val
		head = head.Right
	}
	fmt.Println(res)
}

func TreeToDoublyList(root *TreeNode) *TreeNode {
	// 前驱节点，当前头节点
	var pre, head *TreeNode
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 中序遍历从最左孩子开始
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈处理当前节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 首次遍历，前驱节点为空，则记录头节点
		if pre == nil {
			head = root
		} else {
			// 非首次遍历，前驱节点与刚出栈的根节点相互指向建立连接
			pre.Right = root
			root.Left = pre
		}
		// 前驱节点记录当前根节点，切换到右孩子进行处理
		pre = root
		root = root.Right
	}
	// 单向循环结束以后，需要将最左孩子与最右孩子相互指向建立连接
	if pre != nil && head != nil {
		pre.Right = head
		head.Left = pre
	}
	return head
}

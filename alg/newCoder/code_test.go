package newCoder

import (
	"container/heap"
	linkedList2 "dp/alg/linkedList"
	. "dp/ds"
	"dp/ds/linkedList"
	"dp/ds/tree"
	"dp/tools"
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

/*题目描述
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值
[要求]
set和get方法的时间复杂度为O(1)
某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案

第一次操作后：最常使用的记录为("1", 1)
第二次操作后：最常使用的记录为("2", 2)，("1", 1)变为最不常用的
第三次操作后：最常使用的记录为("3", 2)，("1", 1)还是最不常用的
第四次操作后：最常用的记录为("1", 1)，("2", 2)变为最不常用的
第五次操作后：大小超过了3，所以移除此时最不常使用的记录("2", 2)，加入记录("4", 4)，并且为最常使用的记录，然后("3", 2)变为最不常使用的记录
*/
func TestLRU_coder(t *testing.T) {
	/* [1, -1]
	set(1, 1) [1:1]
	set(2, 2) [2:2, 1:1]
	set(3, 2) [3:2, 2:2, 1:1]
	get(1)    [1:1, 3:2, 2:2] -> 1
	set(4, 4) [4:4, 1:1, 3:2]
	get(2)    [4:4, 1:1, 3:2] -> -1
	*/
	fmt.Println(LRU([][]int{[]int{1, 1, 1}, []int{1, 2, 2}, []int{1, 3, 2}, []int{2, 1}, []int{1, 4, 4}, []int{2, 2}}, 3))
	fmt.Println(LRU1([][]int{[]int{1, 1, 1}, []int{1, 2, 2}, []int{1, 3, 2}, []int{2, 1}, []int{1, 4, 4}, []int{2, 2}}, 3))
}

type node struct {
	k, v      int
	pre, next *node
}

type lru struct {
	head, tail     *node
	size, capacity int
	cache          map[int]*node
}

func construct(k int) lru {
	this := lru{
		head:     &node{},
		tail:     &node{},
		cache:    map[int]*node{},
		capacity: k,
	}
	this.head.next = this.tail
	this.tail.pre = this.head
	return this
}

func (this *lru) set(k, v int) {
	if cur, ok := this.cache[k]; ok {
		cur.v = v
		this.moveToHead(cur)
	} else {
		node := &node{k: k, v: v}
		this.addToHead(node)
		this.cache[k] = node
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.k)
			this.size--
		}
	}
}

func (this *lru) get(k int) int {
	if cur, ok := this.cache[k]; ok {
		this.moveToHead(cur)
		return cur.v
	}
	return -1
}

func (this *lru) moveToHead(cur *node) {
	this.remove(cur)
	this.addToHead(cur)
}

func (this *lru) addToHead(cur *node) {
	cur.pre = this.head
	cur.next = this.head.next
	this.head.next.pre = cur
	this.head.next = cur
}

func (this *lru) remove(cur *node) {
	cur.pre.next = cur.next
	cur.next.pre = cur.pre
}

func (this *lru) removeTail() *node {
	cur := this.tail.pre
	this.remove(cur)
	return cur
}

func LRU1(operators [][]int, k int) []int {
	res := []int{}
	lru := construct(k)
	for _, v := range operators {
		if v[0] == 1 {
			lru.set(v[1], v[2])
		} else {
			res = append(res, lru.get(v[1]))
		}
	}
	return res
}

// k是缓存大小
func LRU(operators [][]int, k int) []int {
	res := []int{}
	lru := LRUConstructor(k)
	for _, op := range operators {
		if op[0] == 1 {
			lru.Set(op[1], op[2])
		} else if op[0] == 2 {
			res = append(res, lru.Get(op[1]))
		}
	}
	return res
}

/*请实现有重复数字的升序数组的二分查找
给定一个 元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1

[1,2,4,4,5],4
2

[1,2,4,4,5],3
-1

[1,1,1,1,1],1
0
*/

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{1, 2, 4, 4, 5}, 4))
	fmt.Println(search([]int{1, 2, 4, 4, 5}, 3))
	fmt.Println(search([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, 3))
}

func TestClimb(t *testing.T) {
	fmt.Println(ClimbStairs(4))
	fmt.Println(ClimbStairsDP(4))
}

/*请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false


说明：

你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


进阶：

你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。


示例：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestStack4Queen(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.In2Out()
	}
	cur := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return cur
}

func (this *MyQueue) In2Out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.In2Out()
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/*给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：
num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestAddStrings(t *testing.T) {
	fmt.Println(AddStrings("515", "357"))
}

/*给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestLowestCommonAncestor(t *testing.T) {
	root := tree.ArrayToTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	fmt.Println(LowestCommonAncestor(root, root.Left, root.Right).Val)
}

/*给定一棵二叉树以及这棵树上的两个节点 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
示例1
输入
[3,5,1,6,2,0,8,#,#,7,4],6,4
返回值
5
    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4
*/

func TestLowestCommonAncestor2(t *testing.T) {
	root := tree.ArrayToTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	fmt.Println(LowestCommonAncestor2(root, 6, 4))
}

/*给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。
输入
"1AB2345CD","12345EF"
返回值
"2345"
*/

func TestLCS(t *testing.T) {
	fmt.Println(LCS("1AB2345CD", "12345EF"))
	fmt.Println(LCS2("1AB2345CD", "12345EF"))
}

func LCS2(str1 string, str2 string) string {
	return ""
}

//给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个容器，请返回容器能装多少水。
//[3,1,2,5,2,4]
//5

func TestMaxWater(t *testing.T) {
	fmt.Println(MaxWater([]int{3, 1, 2, 5, 2, 4}))
}

/*给定一个数组arr，返回arr的最长无的重复子串的长度(无重复指的是所有数字都不相同)。
[2,3,4,5]
4
[2,2,3,4,3]
3
*/

func TestMaxLength(t *testing.T) {
	fmt.Println(MaxLength([]int{2, 3, 4, 5}))
	fmt.Println(MaxLength([]int{2, 2, 3, 4, 3}))
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
	head1 := linkedList.ArrayToLinkedList([]int{7, 2, 4, 2})
	head2 := linkedList.ArrayToLinkedList([]int{5, 6, 4})
	fmt.Println(linkedList.LinkedListToArray(linkedList2.AddTwoNumbers(head1, head2)))
	fmt.Println(linkedList.LinkedListToArray(AddTwoNumbers(head1, head2)))
}

/*给定数组arr，设长度为n，输出arr的最长递增子序列。（如果有多个答案，请输出其中字典序最小的）

输入
[2,1,5,3,6,4,8,9,7]
返回值
[1,3,4,8,9]

示例2
输入
[1,2,8,6,4]
返回值
[1,2,4]
说明
其最长递增子序列有3个，（1，2，8）、（1，2，6）、（1，2，4）其中第三个字典序最小，故答案为（1，2，4）*/

func TestLIS(t *testing.T) {
	fmt.Println(LIS1([]int{2, 1, 5, 3, 6, 4, 8, 9, 7}))
	fmt.Println(LIS([]int{1, 2, 8, 6, 4}))
}

func LIS1(arr []int) []int {
	return []int{}
}

/*对于一个字符串，请设计一个高效算法，计算其中最长回文子串的长度。
给定字符串A以及它的长度n，请返回最长回文子串的长度。

示例1
输入
"abc1234321ab",12

返回值
7
*/

func TestGetLongestPalindrome(t *testing.T) {
	fmt.Println(GetLongestPalindrome("abc1234321ab", 12))
}

/*题目描述
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。

示例1
输入
{1,2,3,4,5,6}
返回值
{1,3,5,2,4,6}

示例2
输入
{1,4,6,3,7}
返回值
{1,6,7,4,3}
说明
奇数节点有1,6,7，偶数节点有4,3。重排后为1,6,7,4,3
*/

func TestOddEvenList(t *testing.T) {
	head := linkedList.ArrayToLinkedList([]int{1, 2, 3, 4, 5, 6})
	cur := OddEvenList(head)
	fmt.Println(linkedList.LinkedListToArray(cur))
}

/*
题目描述
给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。
示例1
输入
[
	[1,3,5,9],
	[8,1,3,4],
	[5,0,6,1],
	[8,8,4,0],
]
返回值
12
*/

func TestMinPathSum(t *testing.T) {
	fmt.Println(MinPathSum([][]int{[]int{1, 3, 5, 9}, []int{8, 1, 3, 4}, []int{5, 0, 6, 1}, []int{8, 8, 4, 0}}))
}

/*
题目描述
给你一个非空模板串S，一个文本串T，问S在T中出现了多少次

示例1
输入
"ababab","abababab"
返回值
2

示例2
输入
"abab","abacabab"
返回值
1

*/

func TestKmp(t *testing.T) {
	fmt.Println(Kmp("ababab", "abababab"))
	fmt.Println(Kmp("abab", "abacabab"))
}

/*

一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
示例1
输入
1
返回值
1

示例2
输入
4
返回值
5

*/

func TestJumpFloor(t *testing.T) {
	fmt.Println(JumpFloor(4))
}

/*给定一棵二叉搜索树，请找出其中的第k小的TreeNode结点。
示例1
输入
{5,3,7,2,4,6,8},3

    5
   / \
  3   7
 / \ / \
2  4 6  8

返回值
{4}
说明
按结点数值大小顺序第三小结点的值为4*/

func TestKthNode(t *testing.T) {
	root := tree.ArrayToTree([]interface{}{8, 6, 10, 5, 7, 9, 11})
	res := KthNode(root, 0)
	if res == nil {
		fmt.Println(res)
		return
	}
	fmt.Println(res.Val)
}

/*给定一个m x n大小的矩阵（m行，n列），按螺旋的顺序返回矩阵中的所有元素。
示例1
输入
[
	[1,2,3],
	[4,5,6],
	[7,8,9],
]
返回值
[1,2,3,6,9,8,7,4,5]*/

// todo
func TestSpiralOrder(t *testing.T) {
	fmt.Println(SpiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(SpiralOrder1([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(SpiralOrder1([][]int{[]int{2, 3}}))
}

func SpiralOrder1(matrix [][]int) []int {
	return []int{}
}

/*

题目描述
输入一个整数，输出该数32位二进制表示中1的个数。其中负数用补码表示。
示例1
输入
10
返回值
2

*/

func TestNumberOf1(t *testing.T) {
	// 1100
	// 1011
	fmt.Println(NumberOf1(10))
}

/*
题目描述
已知int一个有序矩阵mat，同时给定矩阵的大小n和m以及需要查找的元素x，且矩阵的行和列都是从小到大有序的。设计查找算法返回所查找元素的二元数组，代表该元素的行号和列号(均从零开始)。保证元素互异。

示例1
输入
[[1,2,3],[4,5,6]],2,3,6

[
	[1,2,3],
	[4,5,6],
]

返回值
[1,2]

*/

func TestFindElement(t *testing.T) {
	fmt.Println(FindElement([][]int{[]int{1, 2, 3}, []int{4, 5, 6}}, 2, 3, 6))
}

/*有一个NxN整数矩阵，请编写一个算法，将矩阵顺时针旋转90度。
给定一个NxN的矩阵，和矩阵的阶数N,请返回旋转后的NxN矩阵,保证N小于等于300。

示例1
输入
[
	[1,2,3],
	[4,5,6],
	[7,8,9],
],

3

返回值
[
	[7,4,1],
	[8,5,2],
	[9,6,3],
]
*/

func TestRotateMatrix(t *testing.T) {
	fmt.Println(RotateMatrix([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, 3))
}

/*
题目描述
给定一个由0和1组成的2维矩阵，返回该矩阵中最大的由1组成的正方形的面积
示例1
输入
[
	[1,0,1,0,0],
	[1,0,1,1,1],
	[1,1,1,1,1],
	[1,0,0,1,0],
]
返回值
4
*/

func TestMaxSquare(t *testing.T) {
	fmt.Println(MaxSquare([][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}))
	fmt.Println(MaxSquare1([][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}))
}

func MaxSquare1(matrix [][]byte) int {
	max := 0
	return max * max
}

func TestIntersect(t *testing.T) {
	fmt.Println(Intersect([]int{2, 3, 4, 1}, []int{3, 4}))
	fmt.Println(Intersect([]int{2, 3, 4, 1, 5, 7, 0, 2}, []int{7, 3, 4}))
}

/*
题目描述
实现函数 int sqrt(int x).
计算并返回x的平方根（向下取整）
示例1
输入
2
返回值
1
*/

func TestSqrt(t *testing.T) {
	fmt.Println(sqrt(8))
}

func sqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	s := x >> 1
	for s*s > x {
		s >>= 1
	}
	for s*s <= x {
		s++
	}
	return s - 1
}

/*
题目描述
给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
窗口大于数组长度的时候，返回空

示例1
输入
[2,3,4,2,6,2,5,1],3
返回值
[4,4,6,6,6,5]
*/

func TestMaxInWindows(t *testing.T) {
	fmt.Println(MaxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3))
}

/*
实现函数 atoi 。函数的功能为将字符串转化为整数
提示：仔细思考所有可能的输入情况。这个问题没有给出输入的限制，你需要自己考虑所有可能的情况。
*/
func TestAtoi(t *testing.T) {
	fmt.Println(Atoi("-123"))
}

/*给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

示例 1：
输入：k = 1, n = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
如果它没碎，那么肯定能得出 f = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。

示例 2：
输入：k = 2, n = 6
输出：3

示例 3：
输入：k = 3, n = 14
输出：4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/super-egg-drop
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSuperEggDrop(t *testing.T) {
	fmt.Println(SuperEggDrop(3, 14))
	fmt.Println(superEggDrop(3, 14))
}

func SuperEggDrop(k int, n int) int {
	dp, step := make([]int, k+1), 0
	for ; dp[k] < n; step++ {
		for i := k; i > 0; i-- {
			dp[i] = 1 + dp[i] + dp[i-1]
		}
	}
	return step
}

func superEggDrop(k int, n int) int {
	// 第k枚鸡蛋，在第n层至少需要多少层数
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}
	m := 0
	for dp[k][m] < n {
		m++
		for j := 1; j <= k; j++ {
			dp[j][m] = dp[j][m-1] + dp[j-1][m-1] + 1
		}
	}
	return m
}

/*

题目描述
给定一个十进制数M，以及需要转换的进制数N。将十进制数M转化为N进制数
示例1
输入
7,2
返回值
"111"

*/

func TestTen2N(t *testing.T) {
	fmt.Println(Ten2N(17, 12))
	fmt.Println(Hello(17, 12))
}

func Hello(M int, N int) string {
	return ""
}

func Ten2N(M int, N int) string {
	res, sign := "", ""
	if M < 0 {
		sign = "-"
		M = -M
	}
	for M > 0 {
		ac := M % N
		M = M / N
		if ac > 9 {
			res = string(ac-10+'A') + res
		} else {
			res = string(ac+'0') + res
		}
	}
	return sign + res
}

/*
题目描述
给定一个无序单链表，实现单链表的排序(按升序排序)。
示例1
输入
复制
[1,3,2,4,5]
返回值
复制
{1,2,3,4,5}
*/

func TestSortInList(t *testing.T) {
	head := linkedList.ArrayToLinkedList([]int{1, 3, 2, 4, 5, 11, 0, -2, -1, 3})
	res := linkedList.LinkedListToArray(SortInList(head))
	fmt.Println(res)
}

func SortInList(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	return merge(SortInList(head), SortInList(cur))
}

func merge(l1, l2 *linkedList.ListNode) *linkedList.ListNode {
	head := &linkedList.ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

/*
最大数
题目描述
给定一个数组由一些非负整数组成，现需要将他们进行排列并拼接，使得最后的结果最大，返回值需要是string类型 否则可能会溢出
示例1
输入
[30,1]
返回值
"301"
*/

func TestJoinMaximum(t *testing.T) {
	fmt.Println(JoinMaximum([]int{30, 8, 23}))
}

func JoinMaximum(nums []int) string {
	memo := []string{}
	for _, v := range nums {
		memo = append(memo, strconv.Itoa(v))
	}
	sort.Slice(memo, func(i, j int) bool {
		num1, _ := strconv.Atoi(memo[i] + memo[j])
		num2, _ := strconv.Atoi(memo[j] + memo[i])
		return num1 > num2
	})
	res := ""
	for _, v := range memo {
		res += v
	}

	if res[0] == '0' {
		return "0"
	}
	return res
}

/*题目描述
请写一个整数计算器，支持加减乘三种运算和括号。

示例1
输入
"1+2"
返回值
3

示例2
输入
"(2*(3-4))*5"
返回值
-10

示例3
输入
"3+2*3*4-1"
返回值
26

3+3*4/2-1 = 8
*/
func TestDiffWaysToCompute(t *testing.T) {
	//fmt.Println(Calculator("(2*(3-4))*5"))
	fmt.Println(Calculator("3+2*3*4-1"))
	fmt.Println(Calculator1("(2*(3-4))*5"))
	fmt.Println(Calculator1("(2*(3-4))*5"))
}

func Calculator1(s string) int {
	return 0
}

func Calculator(s string) int {
	stack := []int{}
	var sign byte = '+'
	num := 0
	// 如果字符串中还有元素
	for len(s) > 0 {
		// 取出第一个元素
		ch := s[0]
		s = s[1:]
		// 如果是左括号 将括号内的数取出来
		if ch == '(' {
			i, count := 0, 1
			// 计算括号终止位置
			for count > 0 {
				if s[i] == '(' {
					count++
				}
				if s[i] == ')' {
					count--
				}
				i++
			}
			// 计算括号内运算
			num = Calculator(s[0 : i-1])
			s = s[i:]
			if len(s) != 0 {
				continue
			}
		}
		// 如果是数字, 计算结果
		if '0' <= ch && ch <= '9' {
			num = 10*num + int(ch-'0')
			// 如果字符串为空，则当前数字为最后一个数字
			if len(s) != 0 {
				continue
			}
		}
		switch sign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, pre*num)
		case '/':
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, pre/num)
		}
		// 记录符号下一轮计算使用
		sign, num = ch, 0
	}

	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}

/*题目描述
假设你有一个数组，其中第 i 个元素是股票在第 i 天的价格。
你有一次买入和卖出的机会。（只有买入了股票以后才能卖出）。请你设计一个算法来计算可以获得的最大收益。

示例1
输入
[1,4,2]
返回值
3

示例2
输入
[2,4,1]
返回值
2*/

func TestMaxProfit(t *testing.T) {
	fmt.Println(MaxProfit([]int{1, 4, 2}))
	fmt.Println(MaxProfit([]int{2, 4, 1}))
}

func MaxProfit(prices []int) int {
	res := 0
	minPrice := math.MaxInt64
	for _, value := range prices {
		if value < minPrice {
			minPrice = value
		}
		if res < value-minPrice {
			res = value - minPrice
		}
	}
	return res
}

type keyValue struct {
	Ch    string
	Count int
}

type KeyHeap []keyValue

func (h KeyHeap) Len() int {
	return len(h)
}

func (h KeyHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Ch < h[j].Ch
	}
	return h[i].Count > h[j].Count
}

func (h KeyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KeyHeap) Push(x interface{}) {
	*h = append(*h, x.(keyValue))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

/**
 * return topK string
 * @param strings string字符串一维数组 strings
 * @param k int整型 the k
 * @return string字符串二维数组
 */
func TopKStrings(strings []string, k int) (res [][]string) {
	// write code here
	h := &KeyHeap{}
	heap.Init(h)

	var hashMap = make(map[string]int)
	for _, ch := range strings {
		hashMap[ch]++
	}
	for key, value := range hashMap {
		heap.Push(h, keyValue{
			Ch:    key,
			Count: value,
		})
	}
	for i := 0; i < k; i++ {
		item := heap.Pop(h).(keyValue)
		res = append(res, []string{
			item.Ch, strconv.Itoa(item.Count),
		})
	}
	return
}

func TestTopKStrings(t *testing.T) {
	fmt.Println(TopKStrings([]string{"1", "2", "3", "4"}, 2))
}

/*实现一个特殊功能的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。
示例1
输入
[[1,3],[1,2],[1,1],[3],[2],[3]]
返回值
[1,2]*/

func TestGetMinStack(t *testing.T) {
	fmt.Println(GetMinStack([][]int{[]int{1, 3}, []int{1, 2}, []int{1, 1}, []int{3}, []int{2}, []int{3}}))
}

type Stack struct {
	s1 []int
	s2 []int
}

func ConstructStack() *Stack {
	return &Stack{s1: []int{}, s2: []int{}}
}

func GetMinStack(op [][]int) []int {
	s := ConstructStack()
	res := []int{}
	for i := 0; i < len(op); i++ {
		if op[i][0] == 1 {
			if len(s.s1) == 0 {
				s.s1 = append(s.s1, op[i][1])
				s.s2 = append(s.s2, op[i][1])
			} else {
				s.s1 = append(s.s1, op[i][1])
				s.s2 = append(s.s2, tools.Min(op[i][1], s.s2[len(s.s2)-1]))
			}
		} else if op[i][0] == 2 {
			s.s1 = s.s1[:len(s.s1)-1]
			s.s2 = s.s2[:len(s.s2)-1]
		} else {
			res = append(res, s.s2[len(s.s2)-1])
		}
	}
	return res
}

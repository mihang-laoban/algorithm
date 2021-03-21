package newCoder

import (
	linkedList2 "dp/alg/linkedList"
	. "dp/ds"
	"dp/ds/linkedList"
	"dp/ds/tree"
	"fmt"
	"math"
	"strings"
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
	head := linkedList.ArrayToLinkedList([]int{1, 3, 2, 4, 5})
	res := linkedList.LinkedListToArray(SortInList(head))
	fmt.Println(res)
}

func SortInList(head *linkedList.ListNode) *linkedList.ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	first, second := head, slow.Next
	slow.Next = nil
	res := Merge(SortInList(first), second)
	return res
}

func Merge(l1, l2 *linkedList.ListNode) *linkedList.ListNode {
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
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
}

func TestAtoi(t *testing.T) {
	fmt.Println(atoi("    11333713950"))
}

func atoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	num, isPositive := 0, true
	if str[0] == '+' {
		str = str[1:]
		isPositive = true
	} else if str[0] == '-' {
		str = str[1:]
		isPositive = false
	}

	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		num = num*10 + (int(str[i] - '0'))
	}

	if isPositive {
		if num > math.MaxInt32 {
			return math.MaxInt32
		}
		return num
	}
	if -num < math.MinInt32 {
		return math.MinInt32
	}
	return -num
}

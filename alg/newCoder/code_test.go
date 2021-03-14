package newCoder

import (
	linkedList2 "dp/alg/linkedList"
	. "dp/ds"
	"dp/ds/linkedList"
	"dp/ds/tree"
	"dp/tools"
	"fmt"
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

func AddStrings(s string, t string) string {
	ans, count := "", 0
	for i, j := len(s)-1, len(t)-1; i >= 0 || j >= 0 || count != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(s[i] - '0')
		}
		if j >= 0 {
			y = int(t[j] - '0')
		}
		result := x + y + count
		count = result / 10
		ans = strconv.Itoa(result%10) + ans
	}
	return ans
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

func LowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	ans := root
	for {
		if p.Val < ans.Val && q.Val < ans.Val {
			ans = ans.Right
		} else if p.Val > ans.Val && q.Val > ans.Val {
			ans = ans.Left
		} else {
			return ans
		}
	}
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

func LowestCommonAncestor2(root *tree.TreeNode, c1, c2 int) int {
	return findLowestCommonAncestor2(root, c1, c2).Val
}

func findLowestCommonAncestor2(root *tree.TreeNode, c1, c2 int) *tree.TreeNode {
	if root == nil || root.Val == c1 || root.Val == c2 {
		return root
	}
	l := findLowestCommonAncestor2(root.Left, c1, c2)
	r := findLowestCommonAncestor2(root.Right, c1, c2)
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
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
	s1, s2 := len(str1), len(str2)
	dp := make([][]int, s1+1)
	for i := 0; i < s1+1; i++ {
		dp[i] = make([]int, s2+1)
	}
	index, size := 0, 0
	for i := 1; i < s1+1; i++ {
		for j := 1; j < s2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if size < dp[i][j] {
					size = dp[i][j]
					index = j
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	if size == 0 {
		return "-1"
	}
	return str2[index-size : index]
}

func LCS(str1 string, str2 string) string {
	size1, size2 := len(str1), len(str2)
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}
	maxLen, str2Index := 0, 0
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if maxLen < dp[i][j] {
					maxLen = dp[i][j]
					str2Index = j
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	if maxLen == 0 {
		return "-1"
	} else {
		return str2[str2Index-maxLen : str2Index]
	}
}

//给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个容器，请返回容器能装多少水。
//[3,1,2,5,2,4]
//5

func TestMaxWater(t *testing.T) {
	fmt.Println(MaxWater([]int{3, 1, 2, 5, 2, 4}))
}

func MaxWater(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax, res := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if lMax < height[l] {
				lMax = height[l]
			} else {
				res += lMax - height[l]
			}
			l++
		} else {
			if rMax < height[r] {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			r--
		}
	}
	return res
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

func MaxLength(arr []int) int {
	recorder := map[int]int{}
	cur, maxLen := 0, 0
	for i, v := range arr {
		if _, ok := recorder[v]; ok {
			cur = tools.Max(cur, recorder[v]+1)
		}
		recorder[v] = i
		maxLen = tools.Max(maxLen, i-cur+1)
	}
	return maxLen
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

func AddTwoNumbers(head1 *linkedList.ListNode, head2 *linkedList.ListNode) *linkedList.ListNode {
	nums1, nums2 := []int{}, []int{}
	for head1 != nil {
		nums1 = append(nums1, head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		nums2 = append(nums2, head2.Val)
		head2 = head2.Next
	}
	var add int
	var target *linkedList.ListNode
	for len(nums1) > 0 || len(nums2) > 0 || add > 0 {
		var a, b int
		if len(nums1) > 0 {
			a = nums1[len(nums1)-1]
			nums1 = nums1[:len(nums1)-1]
		}
		if len(nums2) > 0 {
			b = nums2[len(nums2)-1]
			nums2 = nums2[:len(nums2)-1]
		}
		res := a + b + add
		add = res / 10
		rest := res % 10
		target = &linkedList.ListNode{Val: rest, Next: target}
	}
	return target
}

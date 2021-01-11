package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"math"
	"sort"
	"testing"
)

func init() {
	Max(1, 2)
}

func traverse(order int) {
	preFunc := []func(*TreeNode) []int{PreOrderLabel, PreOrderRecur, PreOrderLoop, MyPreOrder}
	inFunc := []func(*TreeNode) []int{InOrderLabel, InOrderRecur, InOrderLoop, MyInOrder}
	postFunc := []func(*TreeNode) []int{PostOrderLabel, PostOrderRecur, PostOrderLoop, MyPostOrder}

	for _, root := range GetRootForTraverse() {
		switch order {
		case PRE:
			fmt.Println("Pre:")
			Ordering(root, preFunc)
		case IN:
			fmt.Println("In:")
			Ordering(root, inFunc)
		case POST:
			fmt.Println("Post:")
			Ordering(root, postFunc)
		}
	}
}

func MyPreOrder(root *TreeNode) (res []int)  { return }
func MyInOrder(root *TreeNode) (res []int)   { return }
func MyPostOrder(root *TreeNode) (res []int) { return }

func AllOrderTraverse() {
	for _, v := range []int{ /*PRE, IN,*/ POST} {
		traverse(v)
	}
}

func SerAndDes() {
	// "[1,2,3,nil,nil,4,5,nil,nil,nil,nil]"
	root := GetRootForSerialize()
	str := Serialize(root)
	res := Deserialize(str)
	fmt.Println(Serialize(res))
}

/*给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
  2
 / \
1   3
输出: true
示例 2:

输入:
  5
 / \
1   4
   / \
  3   6
输出: false
解释: 输入为: [5,1,4,nil,nil,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func BuildTreeToValidate() (root *TreeNode) {
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	return
}

func checker(root *TreeNode) bool {
	type Node struct {
		isCur bool
		node  *TreeNode
	}
	last, stack := math.MinInt64, []*Node{{true, root}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.node == nil {
			continue
		}
		if cur.isCur {
			stack = append(stack, &Node{true, cur.node.Right})
			stack = append(stack, &Node{false, cur.node})
			stack = append(stack, &Node{true, cur.node.Left})
		} else {
			if last >= cur.node.Val {
				return false
			}
			last = cur.node.Val
		}
	}
	return true
}

/*输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:
给定二叉树 [3,9,20,nil,nil,15,7]

  3
 / \
9  20
  /  \
15    7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,nil,nil,4,4]

      1
     / \
    2   2
   / \
  3   3
 / \
4   4

	  1
	 / \
	2   2
   /     \
  3       3
 /
4
[1,2,2,3,nil,nil,3,4,nil,nil,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestIsBalancedBottom(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 2, 3, nil, nil, 3, 4, nil, nil, 4})
	fmt.Println(balance(root) >= 0)
}

func balance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left, right := balance(node.Left), balance(node.Right)
	if left == -1 || right == 1 || Abs(left-right) > 1 {
		return -1
	}
	return Max(left, right) + 1
}

func TestGo(t *testing.T) {
	AllOrderTraverse()
	//root := BuildTreeToValidate()
	//fmt.Println(checker(root))
}

func TestTreeAndArray(t *testing.T) {
	arr := []interface{}{5, 1, 4, nil, nil, 3, 6}
	fmt.Println(arr)
	root := ArrayToTree(arr)
	res := TreeToArray(root)
	fmt.Println(res)
}

/*给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
  2
 / \
1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,nil,nil,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestValidateBinaryTree(t *testing.T) {
	nums := []interface{}{5, 1, 4, nil, nil, 3, 6}
	root := ArrayToTree(nums)
	fmt.Println(isValidBST(root))
	fmt.Println(isValidBSTInOrder(root))
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBSTInOrder(root *TreeNode) bool {
	stack := []*TreeNode{}
	inOrder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inOrder {
			return false
		}
		inOrder = root.Val
		root = root.Right
	}
	return true
}

/*
    5
   / \
  2   6
 / \
1   3

*/
//https://leetcode-cn.com/problems/number-of-islands/solution/number-of-islands-shen-du-you-xian-bian-li-dfs-or-/
func Test(t *testing.T) {
	root := InitTree([]int{5, 2, 6, 1, 3}).GetRoot()
	//52136
	fmt.Println(Pre(root))
	//12356
	fmt.Println(In(root))
	//fmt.Println(InOrderLoop(root))
	//13265
	fmt.Println(Post(root))
	//fmt.Println(BFS(root))
}

func Pre(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append(res, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return
}

func In(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			root = cur.Right
		}
	}
	return
}

/*
    5
   / \
  2   6
 / \
1   3

*/
func Post(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append([]int{root.Val}, res...)
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return
}

/*给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,nil,nil,15,7],

  3
 / \
9  20
  /  \
15    7
返回其层序遍历结果：

[
[3],
[9,20],
[15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestBFS(t *testing.T) {
	root := ArrayToTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(BFS2(root))
}

func BFSArr(root *TreeNode) (res [][]int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmp, size := []int{}, len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmp)
	}
	return
}

func BFS2(root *TreeNode) (res []int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return
}

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
]
输出：1

示例 2：

输入：grid = [
["1","1","0","0","0"],
["1","1","0","0","0"],
["0","0","1","0","0"],
["0","0","0","1","1"]
]
输出：3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIslandNum(t *testing.T) {
	grid := [][]string{
		[]string{"1", "1", "0", "0", "0"},
		[]string{"1", "1", "0", "0", "0"},
		[]string{"0", "0", "1", "0", "0"},
		[]string{"0", "0", "0", "1", "1"},
	}

	fmt.Println(IslandBFS(grid))
	fmt.Println(IslandBfs(grid))
}

func IslandBfs(grid [][]string) int {
	count := 0
	return count
}

/*假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。


示例 1:

输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。
示例 2:

输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/assign-cookies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestAssignCookies(t *testing.T) {
	children := []int{1, 2}
	cookies := []int{1, 2, 3}
	fmt.Println(FindContentChildren(children, cookies))
}

func FindContentChildren(g, s []int) (count int) {
	sort.Ints(g)
	sort.Ints(s)
	// 查找饼干
	for j := 0; count < len(g) && j < len(s); j++ {
		// 如果饼干能满足当前孩子，就看下一个孩子
		// 如果不能就看下一个饼干
		if g[count] <= s[j] {
			count++
		}
	}
	return count
}

//https://leetcode-cn.com/problems/sqrtx/
//https://leetcode-cn.com/problems/valid-perfect-square/
//https://leetcode-cn.com/problems/search-rotate-array-lcci/

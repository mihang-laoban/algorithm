package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

//https://leetcode-cn.com/problems/number-of-islands/solution/number-of-islands-shen-du-you-xian-bian-li-dfs-or-/
func TestTraverse(t *testing.T) {
	root := ArrayToTree([]interface{}{5, 2, 6, 1, 3, 4})
	//52136
	fmt.Println("pre")
	fmt.Println(Pre(root))
	fmt.Println(Pre2(root))
	fmt.Println(MorrisPre(root))
	fmt.Println(MorrisPre2(root))
	fmt.Println("in")
	fmt.Println(In(root))
	fmt.Println(In2(root))
	fmt.Println(MorrisIn(root))
	fmt.Println("post")
	//13265
	fmt.Println(Post(root))
	fmt.Println(MorrisPost(root))
	fmt.Println("bfs")
	//52613
	fmt.Println(Bfs(root))
}

func MorrisPre2(root *TreeNode) (res []int) {

	return
}

/*
    5
   / \
  2   6
 / \
1   3
*/
func MorrisPre(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	cur := root //当前开始遍历的节点
	for cur != nil {
		post := cur.Left
		if post != nil {
			for post.Right != nil && post.Right != cur { //找到当前左子树的最右侧节点，且这个节点应该在指向根结点之前，否则整个节点又回到了根结点。
				post = post.Right
			}
			if post.Right == nil { //这个时候如果最右侧这个节点的右指针没有指向根结点，创建连接然后往下一个左子树的根结点进行连接操作。
				post.Right = cur
				res = append(res, cur.Val)
				cur = cur.Left
				continue
			} else { //当左子树的最右侧节点有指向根结点，此时说明我们已经回到了根结点并重复了之前的操作，同时在回到根结点的时候我们应该已经处理完 左子树的最右侧节点 了，把路断开。
				post.Right = nil
			}
		} else {
			res = append(res, cur.Val)
		}
		cur = cur.Right //一直往右边走，参考图
	}
	return
}

func MorrisIn(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	cur := root //当前开始遍历的节点
	for cur != nil {
		post := cur.Left
		if post != nil {
			for post.Right != nil && post.Right != cur { //找到当前左子树的最右侧节点，且这个节点应该在指向根结点之前，否则整个节点又回到了根结点。
				post = post.Right
			}
			if post.Right == nil { //这个时候如果最右侧这个节点的右指针没有指向根结点，创建连接然后往下一个左子树的根结点进行连接操作。
				post.Right = cur
				cur = cur.Left
				continue
			} else { //当左子树的最右侧节点有指向根结点，此时说明我们已经回到了根结点并重复了之前的操作，同时在回到根结点的时候我们应该已经处理完 左子树的最右侧节点 了，把路断开。
				post.Right = nil
			}
		}
		res = append(res, cur.Val)
		cur = cur.Right //一直往右边走，参考图
	}
	return
}

/*
    5
   / \
  2   6
 / \ /
1  3 4
*/

func MorrisPost(root *TreeNode) (res []int) {
	reverse := func(nums []int) {
		for i, n := 0, len(nums); i < n>>1; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
	}

	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	cur := root
	for cur != nil {
		post := cur.Left
		if post != nil {
			for post.Right != nil && post.Right != cur {
				post = post.Right
			}
			if post.Right == nil {
				post.Right = cur
				cur = cur.Left
				continue
			}
			post.Right = nil
			addPath(cur.Left)
		}
		cur = cur.Right
	}
	addPath(root)
	return
}

func Pre2(root *TreeNode) (res []int) {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return
}

func Pre(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
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
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return
}

func In2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
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
	for len(stack) > 0 || root != nil {
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

func Bfs(root *TreeNode) (res []int) {
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

func TestBuildTreeToValidate(t *testing.T) {
	AllOrderTraverse()
	root := BuildTreeToValidate()
	fmt.Println(checker(root))
}

func TestValidateBinaryTree(t *testing.T) {
	nums := []interface{}{5, 1, 4, nil, nil, 3, 6}
	root := ArrayToTree(nums)
	fmt.Println(isValidBST(root))
	fmt.Println(isValidBSTInOrder(root))
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
	//fmt.Println(BFS2(root))
	fmt.Println(BFSArr(root))
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

/*
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

说明：不要使用任何内置的库函数，如  sqrt。

示例 1：

输入：16
输出：True
示例 2：

输入：14
输出：False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TestPerfectSquare(t *testing.T) {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquareOld(16))
}

/*
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。


示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSearchRotatedArray(t *testing.T) {
	target, nums := 5, []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(target, nums))
}

/*给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。

进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？

示例 1：
输入：root = [1,3,nil,nil,2]
输出：[3,1,nil,nil,2]
解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

示例 2：
输入：root = [3,1,4,nil,nil,2]
输出：[2,1,4,nil,nil,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
https://leetcode-cn.com/problems/recover-binary-search-tree/solution/hui-fu-er-cha-sou-suo-shu-by-leetcode-solution/
*/
func TestRecoverTree(t *testing.T) {
	tree1 := ArrayToTree([]interface{}{1, 3, nil, nil, 2})
	//tree2 := ArrayToTree([]interface{}{3, 1, 4, nil, nil, 2})
	recoverT(tree1)
	RecoverTree(tree1)
	//RecoverTreeMorris(tree2)
	fmt.Println(BFStoArray(tree1))
	//fmt.Println(BFStoArray(tree2))
}

/*翻转一棵二叉树。
示例：
输入：
    4
   / \
  2   7
 / \ / \
1  3 6  9
输出：

	4
   / \
  7   2
 / \ / \
9  6 3  1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestReverseTree(t *testing.T) {
	root := ArrayToTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
	root = ReverseTreeR(root)
	root = ReverseTreeL(root)
	fmt.Println(BFS(root))
}

/*输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

    5
   / \
  2   6
 / \
1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestVerifyPostOrder(t *testing.T) {
	//fmt.Println(VerifyPostOrder([]int{1, 6, 3, 2, 5}))
	fmt.Println(VerifyPostOrder([]int{1, 3, 2, 6, 5}))
}

/*将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
     0
    / \
  -3   9
  /   /
-10  5

[0 -10 5 <nil> -3 <nil> 9]
     0
    / \
  -10  5
    \   \
    -3   9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestSortedArrayToBST(t *testing.T) {
	head := SortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(TreeToArray(head))
}

/*给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
 / \
15  7
返回它的最大深度 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMaxDepth(t *testing.T) {
	fmt.Println(MaxDepth(ArrayToTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
}

/*翻转一棵二叉树。
示例：
输入：

    4
   / \
  2   7
 / \ / \
1  3 6  9
输出：

    4
   / \
  7   2
 / \ / \
9  6 3  1

备注:
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestInvertTree(t *testing.T) {
	root := ArrayToTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
	newRoot := InvertTree(root)
	fmt.Println(TreeToArray(newRoot))
}

/*
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

  3
 / \
9   20
   /  \
  15   7
返回其自底向上的层序遍历为：
[
	[15,7],
	[9,20],
	[3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestLevelOrderBottom(t *testing.T) {
	root := ArrayToTree([]interface{}{})
	res := LevelOrderBottom(root)
	fmt.Println(res)
}

/*给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。
注意：
给定的目标值 target 是一个浮点数
题目保证在该二叉搜索树中只会存在一个最接近目标值的数
示例：

输入: root = [4,2,5,1,3]，目标值 target = 3.714286
    4
   / \
  2   5
 / \
1   3
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/closest-binary-search-tree-value
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestClosestValue(t *testing.T) {
	//root := ArrayToTree([]interface{}{4, 2, 5, 1, 3})
	//fmt.Println(ClosestValue(root, 3.714286))
	root := ArrayToTree([]interface{}{1500000000, 1400000000})
	fmt.Println(ClosestValue(root, -1500000000.0))
}

/*给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树
    1
   / \
  2   3
 / \
4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestDiameterOfBinaryTree(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(DiameterOfBinaryTree(root))
}

/*给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

示例 1:

输入:
Tree 1                     Tree 2				Tree 3
    1                         2                    2
   / \                       / \                  / \
  3   2                     1   3                1   3
 /                           \   \              / \
5                             4   7            4   7
输出:
合并后的树:
    3                 3
   / \               / \
  4   5             4   5
 / \   \           / \
5   4   7         9   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMergeTrees(t *testing.T) {
	root1 := ArrayToTree([]interface{}{1, 3, 2, 5})
	root2 := ArrayToTree([]interface{}{2, 1, 3, 4, 7})
	MergeTreesL(root1, root2)
	fmt.Println(TreeToArray(root1))
	//fmt.Println(TreeToArray(MergeTreesL(root1, root2)))
}

/*给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

  1
 / \
2   2
 \   \
  3   3
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIsSymmetric(t *testing.T) {
	root1 := ArrayToTree([]interface{}{1, 2, 2, 3, 4, 4, 3})
	root2 := ArrayToTree([]interface{}{1, 2, 2, 3, 3})
	fmt.Println(IsSymmetricL(root1))
	fmt.Println(IsSymmetricL(root2))
}

/*给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。
示例:

输入:
  1
 / \
2   3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestBinaryTreePaths(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 5})
	fmt.Println(BinaryTreePathsR(root))
}

/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：false

示例 3：
输入：root = [1,2], targetSum = 0
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestHasPathSum(t *testing.T) {
	root := ArrayToTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1})
	fmt.Println(HasPathSum(root, 22))
}

/*给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true

示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false

示例 3：
输入：p = [1,2,1], q = [1,1,2]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/same-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestIsSameTree(t *testing.T) {
	root1 := ArrayToTree([]interface{}{1, 2, 1})
	root2 := ArrayToTree([]interface{}{1, 1, 2})
	fmt.Println(IsSameTree(root1, root2))
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

/*给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
更正式地说，root.val = min(root.left.val, root.right.val) 总成立。
给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。

示例 1：
  2
 / \
2   5
   / \
  5   7
输入：root = [2,2,5,null,null,5,7]
输出：5
解释：最小的值是 2 ，第二小的值是 5 。

示例 2：
输入：root = [2,2,2]
输出：-1
解释：最小的值是 2, 但是不存在第二小的值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFindSecondMinimumValue(t *testing.T) {
	//root := ArrayToTree([]interface{}{2, 2, 5, nil, nil, 5, 7})
	root := ArrayToTree([]interface{}{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1}) // 2
	fmt.Println(FindSecondMinimumValue(root))
}

/*给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。

示例：
输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树节点对象(TreeNode object)，而不是数组。

给定的树 [4,2,6,1,3,null,null] 可表示为下图:

    4
   / \
  2   6
 / \
1   3

最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMinDiffInBST(t *testing.T) {
	//root := ArrayToTree([]interface{}{4, 2, 6, 1, 3, nil, nil})
	root := ArrayToTree([]interface{}{90, 69, nil, 49, 89, nil, 52})
	//root := ArrayToTree([]interface{}{1, 0, 48, nil, nil, 12, 49})
	fmt.Println(MinDiffInBSTR(root))
	fmt.Println(MinDiffInBSTL(root))
}

/*给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

    3
   / \
  4   5
 / \
1   2
给定的树 t：
  4
 / \
1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：
    3
   / \
  4   5
 / \
1   2
   /
  0
给定的树 t：

  4
 / \
1   2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subtree-of-another-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIsSubtree(t *testing.T) {
	root1 := ArrayToTree([]interface{}{3, 4, 5, 1, 2})
	root2 := ArrayToTree([]interface{}{4, 1, 2})
	fmt.Println(IsSubtree(root1, root2))
}

/*
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
    1
   / \
  2   3
 /
4
输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:

输入: 二叉树: [1,2,3,null,4]
  1
 / \
2   3
 \
  4
输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-string-from-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestTree2str(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4})
	fmt.Println(Tree2str(root))
	fmt.Println(Tree2str2(root))
}

/*给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

	1
	 \
	  2
	 /
	2
返回[2].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFindMode(t *testing.T) {
	root := ArrayToTree([]interface{}{1, nil, 2, 2})
	fmt.Println(FindMode(root))
	fmt.Println(FindModeMorris(root))
}

func FindMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

func FindModeMorris(root *TreeNode) (answer []int) {
	var base, count, maxCount int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return
}

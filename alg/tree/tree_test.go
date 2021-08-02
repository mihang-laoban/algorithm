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
	root := ArrayToTree([]interface{}{5, 2, 6, 1, 3})
	fmt.Println("My order:")
	fmt.Println(MyOrder(root))
	//52136
	fmt.Println("pre")
	fmt.Println(Pre(root))
	fmt.Println(Pre2(root))
	fmt.Println(MorrisPre(root))
	// 12356
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

func MyOrder(root *TreeNode) (res []int) {
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
		pre := cur.Left
		if pre != nil {
			for pre.Right != nil && pre.Right != cur { //找到当前左子树的最右侧节点，且这个节点应该在指向根结点之前，否则整个节点又回到了根结点。
				pre = pre.Right
			}
			if pre.Right == nil { //这个时候如果最右侧这个节点的右指针没有指向根结点，创建连接然后往下一个左子树的根结点进行连接操作。
				pre.Right = cur
				res = append(res, cur.Val)
				cur = cur.Left
				continue
			}
			pre.Right = nil
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
		pre := cur.Left
		if pre != nil {
			for pre.Right != nil && pre.Right != cur { //找到当前左子树的最右侧节点，且这个节点应该在指向根结点之前，否则整个节点又回到了根结点。
				pre = pre.Right
			}
			if pre.Right == nil { //这个时候如果最右侧这个节点的右指针没有指向根结点，创建连接然后往下一个左子树的根结点进行连接操作。
				pre.Right = cur
				cur = cur.Left
				continue
			}
			pre.Right = nil //当左子树的最右侧节点有指向根结点，此时说明我们已经回到了根结点并重复了之前的操作，同时在回到根结点的时候我们应该已经处理完 左子树的最右侧节点 了，把路断开。
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
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for node != nil {
			res = append(res, node.Val)
			node = node.Right
		}
		ReverseArr1(res[resSize:])
	}

	cur := root
	for cur != nil {
		pre := cur.Left
		if pre != nil {
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
				continue
			}
			pre.Right = nil
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
	fmt.Println(IsBalancedBottom(root))
	fmt.Println(IsBalancedTop(root))
}

/*给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例1:

输入:
  2
 / \
1   3
输出: true
示例2:

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
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

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

说明：不要使用任何内置的库函数，如 sqrt。

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
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为[4,5,6,7,0,1,2] ）。
请你在数组中搜索target ，如果数组中存在这个目标值，则返回它的索引，否则返回-1。


示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例2：
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
	//recoverT(tree1)
	//RecoverTree(tree1)
	RecoverTreeMorris(tree1)
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

/*输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

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
本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。
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
说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
 / \
15  7
返回它的最大深度3 。

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

/*给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

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
*/

func TestLevelOrder(t *testing.T) {
	root := ArrayToTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(LevelOrder(root))
}

/*给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树
    1
   / \
  2   3
 / \
4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestDiameterOfBinaryTree(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(DiameterOfBinaryTree(root))
}

/*给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。

示例1:

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
	MergeTreesR(root1, root2)
	fmt.Println(TreeToArray(root1))
	//fmt.Println(TreeToArray(MergeTreesL(root1, root2)))
}

/*给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

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
说明:叶子节点是指没有子节点的节点。
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
给你二叉树的根节点root 和一个表示目标和的整数targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。
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

/*给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
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

/*给定一个二叉搜索树的根节点root，返回树中任意两节点的差的最小值。

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

/*给定两个非空二叉树 s 和 t，检验s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

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

/*
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
示例：
输入：
1
 \
  3
 /
2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。*/
func TestGetMinimumDifference(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	fmt.Println(GetMinimumDifference1(root))
}

/*
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
示例 1：

输入：
  3
 / \
9   20
   /  \
  15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestAverageOfLevels(t *testing.T) {
	root := ArrayToTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	fmt.Println(AverageOfLevels(root))
}

/*给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

案例 1:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9
输出: True

Target = 28
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFindTarget(t *testing.T) {
	root := ArrayToTree([]interface{}{1})
	fmt.Println(FindTarget(root, 2))
	fmt.Println(FindTarget(root, 28))
}

func FindTarget(root *TreeNode, k int) bool {
	queue := []*TreeNode{root}
	target := map[int]bool{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if target[cur.Val] {
			return true
		}
		target[k-cur.Val] = true
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return false
}

/*给定一个二叉树，计算 整个树 的坡度 。
一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
整个树 的坡度就是其所有节点的坡度之和。

示例 1：
  1
 / \
2   3
输入：root = [1,2,3]
输出：1
解释：
节点 2 的坡度：|0-0| = 0（没有子节点）
节点 3 的坡度：|0-0| = 0（没有子节点）
节点 1 的坡度：|2-3| = 1（左子树就是左子节点，所以和是 2 ；右子树就是右子节点，所以和是 3 ）
坡度总和：0 + 0 + 1 = 1

示例 2：
    4
   / \
  2   9
 / \   \
3   5   7
输入：root = [4,2,9,3,5,null,7]
输出：15
解释：
节点 3 的坡度：|0-0| = 0（没有子节点）
节点 5 的坡度：|0-0| = 0（没有子节点）
节点 7 的坡度：|0-0| = 0（没有子节点）
节点 2 的坡度：|3-5| = 2（左子树就是左子节点，所以和是 3 ；右子树就是右子节点，所以和是 5 ）
节点 9 的坡度：|0-7| = 7（没有左子树，所以和是 0 ；右子树正好是右子节点，所以和是 7 ）
节点 4 的坡度：|(3+5+2)-(9+7)| = |10-16| = 6（左子树值为 3、5 和 2 ，和是 10 ；右子树值为 9 和 7 ，和是 16 ）
坡度总和：0 + 0 + 0 + 2 + 7 + 6 = 15

示例 3：
      21
     /  \
    7    14
   / \    \
  1   1    22
 / \
3   3
输入：root = [21,7,14,1,1,2,2,3,3]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-tilt
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFindTilt(t *testing.T) {
	root := ArrayToTree([]interface{}{4, 2, 9, 3, 5, nil, 7})
	fmt.Println(FindTilt(root))
}

/*给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

     3
    / \
   9   20
  / \
15   7

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMinDepth(t *testing.T) {
	root1 := ArrayToTree([]interface{}{3, 9, 20, 15, 7, nil, nil})
	fmt.Println(MinDepth1(root1))
	fmt.Println(MinDepth2(root1))
}

/*计算给定二叉树的所有左叶子之和。

示例：

  3
 / \
9  20
  /  \
 15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

    1
   / \
  2   3
 / \
4   5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSumOfLeftLeaves(t *testing.T) {
	root1 := ArrayToTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	root2 := ArrayToTree([]interface{}{1, 2, 3, 4, 5})
	fmt.Println(SumOfLeftLeavesBFS(root1))
	fmt.Println(SumOfLeftLeavesDFS(root2))
}

/*路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

示例 1：
  1
 / \
2   3
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

示例 2：
   -10
   / \
  9   20
     /  \
    15   7
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestMaxPathSum(t *testing.T) {
	root := ArrayToTree([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	fmt.Println(MaxPathSum(root))
}

/*272. 最接近的二叉搜索树值 II
给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的 k 个值。

注意：

给定的目标值 target 是一个浮点数
你可以默认 k 值永远是有效的，即 k ≤ 总结点数
题目保证该二叉搜索树中只会存在一种 k 个值集合最接近目标值
示例：

输入: root = [4,2,5,1,3]，目标值 = 3.714286，且 k = 2

    4
   / \
  2   5
 / \
1   3

输出: [4,3]
拓展：
假设该二叉搜索树是平衡的，请问您是否能在小于 O(n)（n 为总结点数）的时间复杂度内解决该问题呢？*/

func TestClosestKValues(t *testing.T) {
	root := ArrayToTree([]interface{}{4, 2, 5, 1, 3})
	fmt.Println(ClosestKValues(root, 3.714286, 2))
}

/*序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

示例 1：
输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[1,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSerAndDes(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, nil, nil, 4, 5})
	codec := ConstructorCodec()
	//str := codec.SerializeDFS(root)
	str := codec.SerializeBFS(root)
	fmt.Println(str)
	newRoot := codec.DeserializeDFS(str)
	fmt.Println(TreeToArray(newRoot))
}

/*在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。每一个节点只有一个父节点，除了根节点没有父节点。
输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n 中的两个不同顶点间，这条附加的边不属于树中已存在的边。
结果图是一个以边组成的二维数组edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是 vi 的一个父节点。
返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。

示例 1：
输入：edges = [[1,2],[1,3],[2,3]]
输出：[2,3]

示例 2：
输入：edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
输出：[4,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/redundant-connection-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestFindRedundantDirectedConnection(t *testing.T) {
	edges := [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}
	fmt.Println(FindRedundantDirectedConnection(edges))
}

func FindRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	numNodes := len(edges)
	uf := newUnionFind(numNodes + 1)
	parent := make([]int, numNodes+1) // parent[i] 表示 i 的父节点
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) { // from 和 to 已连接
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树[3,9,20,null,null,15,7],

  3
 / \
9  20
  /  \
15    7
返回锯齿形层序遍历如下：

[
[3],
[20,9],
[15,7]
]

    1
   / \
  2   3
 /     \
4       5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestZigzagLevelOrder(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4, nil, nil, 5})
	arr := ZigzagLevelOrder(root)
	fmt.Println(arr)
}

/*给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入:[1,2,3,3,5,nil,4,nil,8]
输出:[1, 3, 4]
解释:

    1       <---
   / \
  2   3     <---
 / \   \
3   5   4   <---
 \
  8			<---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestRightSideView(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 3, 5, nil, 4, nil, 8})
	fmt.Println(RightSideViewBFS(root))
	fmt.Println(RightSideViewDFSRecur(root))
}

/*根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

  3
 / \
9   20
   /  \
  15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestBuildTree(t *testing.T) {
	root := BuildTreePreIn([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	root1 := BuildTree1([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(TreeToArray(root))
	fmt.Println(TreeToArray(root1))
}

func BuildTree1(pre, in []int) *TreeNode {
	return nil
}

/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder =[9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

  3
 / \
9  20
  /  \
 15   7


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestBuildTreeInPost(t *testing.T) {
	inOrder := []int{9, 3, 15, 20, 7}
	postOrder := []int{9, 15, 7, 20, 3}
	root := BuildTreeInPost(inOrder, postOrder)
	fmt.Println(TreeToArray(root))
}

func BuildTreeInPost(inOrder []int, postOrder []int) *TreeNode {
	idxMap := map[int]int{}
	for k, v := range inOrder {
		idxMap[v] = k
	}
	var build func(int, int) *TreeNode
	build = func(inLeft, inRight int) *TreeNode {
		// 无剩余节点
		if inLeft > inRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		cur := postOrder[len(postOrder)-1]
		postOrder = postOrder[:len(postOrder)-1]
		root := &TreeNode{Val: cur}

		// 根据 cur 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		root.Right = build(idxMap[cur]+1, inRight)
		root.Left = build(inLeft, idxMap[cur]-1)
		return root
	}
	return build(0, len(inOrder)-1)
}

/*给定一个二叉树，确定它是否是一个完全二叉树。
百度百科中对完全二叉树的定义如下：
若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~2h个节点。）

示例 1：

输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
示例 2：

输入：[1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIsCompleteTree(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4, 5, nil, 7})
	fmt.Println(IsCompleteTree(root))
}

func IsCompleteTree(root *TreeNode) bool {
	pre, queue := root, []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur != nil && pre == nil {
			return false
		}
		if cur != nil {
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
		pre = cur
	}
	return true
}

/*给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：
    1
   / \
  2   5
 / \   \
3   4   6

1 > 2 > 3 > 4 > 5 > 6

输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFlatten(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 5, 3, 4, nil, 6})
	Flatten(root)
	fmt.Println(TreeToArray(root))
}

func Flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			pre := cur.Left
			next := pre
			for next.Right != nil {
				next = next.Right
			}
			next.Right = cur.Right
			cur.Left = nil
			cur.Right = pre
		}
		cur = cur.Right
	}
}

/*给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:
输入:

    5
   / \
  4   5
 / \   \
1   1   5

输出:
2

示例 2:
输入:

    1
   / \
  4   5
 / \   \
4   4   5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-univalue-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestLongestUniValuePath(t *testing.T) {
	root1 := ArrayToTree([]interface{}{5, 4, 5, 1, 1, nil, 5})
	root2 := ArrayToTree([]interface{}{1, 4, 5, 4, 4, nil, 5})
	fmt.Println(LongestUniValuePath(root1))
	fmt.Println(LongestUniValuePath(root2))
}

func LongestUniValuePath(root *TreeNode) int {
	res := 0
	var arrowLength func(*TreeNode) int
	arrowLength = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := arrowLength(root.Left), arrowLength(root.Right)
		var arrowL, arrowR int
		if root.Left != nil && root.Left.Val == root.Val {
			arrowL += l + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			arrowR += r + 1
		}
		res = Max(res, arrowL+arrowR)
		return Max(arrowL, arrowR)
	}
	arrowLength(root)
	return res
}

/*给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

示例 1：
  1
 / \
0   2
输入：root = [1,0,2], low = 1, high = 2
输出：[1,null,2]

示例 2：
  3
 / \
0   4
 \
  2
 /
1

输入：root = [3,0,4,null,2,null,null,1], low = 1, high = 3
输出：[3,2,null,1]

示例 3：
输入：root = [1], low = 1, high = 2
输出：[1]

示例 4：
输入：root = [1,null,2], low = 1, high = 3
输出：[1,null,2]

示例 5：
输入：root = [1,null,2], low = 2, high = 4
输出：[2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trim-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestTrimBST(t *testing.T) {
	root := ArrayToTree([]interface{}{3, 0, 4, nil, 2, nil, nil, 1})
	fmt.Println(TreeToArray(TrimBST(root, 1, 3)))
}

func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > high {
		return TrimBST(root.Left, low, high)
	}
	if root.Val < low {
		return TrimBST(root.Right, low, high)
	}
	root.Left = TrimBST(root.Left, low, high)
	root.Right = TrimBST(root.Right, low, high)
	return root
}

/*给你一棵二叉搜索树（BST）、它的根结点 root以及目标值 V。
请将该树按要求拆分为两个子树：其中一个子树结点的值都必须小于等于给定的目标值 V；另一个子树结点的值都必须大于目标值 V；树中并非一定要存在值为 V的结点。
除此之外，树中大部分结构都需要保留，也就是说原始树中父节点 P 的任意子节点 C，假如拆分后它们仍在同一个子树中，那么结点 P 应仍为 C 的父结点。
你需要返回拆分后两个子树的根结点 TreeNode，顺序随意。

示例：
输入：root = [4,2,6,1,3,5,7], V = 2
输出：[[2,1],[4,3,6,null,null,5,7]]
解释：
注意根结点 output[0] 和 output[1] 都是 TreeNode对象，不是数组。

给定的树 [4,2,6,1,3,5,7] 可化为如下示意图：

    4
   / \
  2   6
 / \ / \
1  3 5  7

输出的示意图如下：

  4
 / \
3   6       和    2
   / \           /
  5   7         1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSplitBST(t *testing.T) {
	root := ArrayToTree([]interface{}{4, 2, 6, 1, 3, 5, 7})
	roots := SplitBST(root, 2)
	for _, value := range roots {
		fmt.Println(TreeToArray(value))
	}
}

func SplitBST(root *TreeNode, v int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	} else if root.Val <= v {
		bns := SplitBST(root.Right, v)
		root.Right = bns[0]
		bns[0] = root
		return bns
	} else {
		bns := SplitBST(root.Left, v)
		root.Left = bns[1]
		bns[1] = root
		return bns
	}
}

/*给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

  5
 / \
2   6
 \   \
  4   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestDeleteNode(t *testing.T) {
	root := ArrayToTree([]interface{}{5, 3, 6, 2, 4, nil, 7})
	fmt.Println(TreeToArray(DeleteNode(root, 3)))
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

/*给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

示例 1：

   6
  / \
 /   \
3     5
 \   /
  2 0
   \
    1

输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
- [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
- 空数组，无子节点。
- [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
- 空数组，无子节点。
- 只有一个元素，所以子节点是一个值为 1 的节点。
- [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
- 只有一个元素，所以子节点是一个值为 0 的节点。
- 空数组，无子节点。

示例 2：
输入：nums = [3,2,1]
输出：[3,null,2,null,1]

3
 \
  2
   \
    1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestConstructMaximumBinaryTree(t *testing.T) {
	fmt.Println(TreeToArray(ConstructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})))
}

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums))
}

func construct(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}
	ma := max(nums, l, r)
	root := &TreeNode{Val: nums[ma]}
	root.Left = construct(nums, l, ma)
	root.Right = construct(nums, ma+1, r)
	return root
}

func max(nums []int, l, r int) int {
	maxi := l
	for i := l; i < r; i++ {
		if nums[maxi] < nums[i] {
			maxi = i
		}
	}
	return maxi
}

/*给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第k个最小元素（从 1 开始计数）。

示例 1：

  3
 / \
1   4
 \
  2

输入：root = [3,1,4,null,2], k = 1
输出：1
示例 2：

      5
     / \
    3   6
   / \
  2   4
 /
1

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestKthSmallest(t *testing.T) {
	root := ArrayToTree([]interface{}{5, 3, 6, 2, 4, nil, nil, 1})
	fmt.Println(KthSmallest(root, 3))
}

func KthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return root.Val
}

/*给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPathSum(t *testing.T) {
	root := ArrayToTree([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	fmt.Println(PathSum(root, 8))
}

func PathSum(root *TreeNode, sum int) int {
	res := 0
	prefix := map[int]int{0: 1}
	var dfs func(*TreeNode, int, int)
	dfs = func(root *TreeNode, sum, curSum int) {
		if root == nil {
			return
		}
		curSum += root.Val
		if _, ok := prefix[curSum-sum]; ok {
			res += prefix[curSum-sum]
		}
		prefix[curSum]++
		dfs(root.Left, sum, curSum)
		dfs(root.Right, sum, curSum)
		prefix[curSum]--
	}
	dfs(root, sum, 0)
	return res
}

/*给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。
每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

示例 1:
输入:

    1
   / \
  3   2
 / \   \
5   3   9
输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。

示例 2:
输入:

    1
   /
  3
 / \
5   3

输出: 2
解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
示例3:

输入:

    1
   / \
  3   2
 /
5

输出: 2
解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
示例 4:

输入:

      1
     / \
    3   2
   /     \
  5       9
 /         \
6           7
输出: 8
解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestWidthOfBinaryTree(t *testing.T) {
	root1 := ArrayToTree([]interface{}{6, 3, 2, 5, nil, nil, 9, 6, nil, nil, nil, nil, nil, nil, 7})
	root2 := ArrayToTree([]interface{}{1, 3, 2, 5, 3, nil, 9})
	fmt.Println(WidthOfBinaryTree(root1))
	fmt.Println(WidthOfBinaryTree1(root2))
	fmt.Println(WidthOfBinaryTree(root2))
}

func WidthOfBinaryTree1(root *TreeNode) int {
	return 0
}

/*给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:
  2
 / \
1   3
输出:
1


示例 2:

输入:
    1
   / \
  2   3
 /   / \
4   5   6
   /
  7
输出:
7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-bottom-left-tree-value
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestFindBottomLeftValue(t *testing.T) {
	root := ArrayToTree([]interface{}{1, 2, 3, 4, nil, 5, 6, nil, nil, nil, nil, 7})
	fmt.Println(FindBottomLeftValue(root))
}

func FindBottomLeftValue(root *TreeNode) int {
	cur, q := &TreeNode{}, []*TreeNode{root}
	for len(q) > 0 {
		cur = q[0]
		q = q[1:]
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
	}
	return cur.Val
}

/*返回与给定前序遍历preorder 相匹配的二叉搜索树（binary search tree）的根结点。
(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于node.left的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，前序遍历首先显示节点node 的值，然后遍历 node.left，接着遍历 node.right。）
题目保证，对于给定的测试用例，总能找到满足要求的二叉搜索树。

示例：

输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestBstFromPreorder(t *testing.T) {
	fmt.Println(TreeToArray(BstFromPreorder([]int{8, 5, 1, 7, 10, 12})))
}

func BstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	var insert func(int, *TreeNode) *TreeNode
	insert = func(val int, root *TreeNode) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		} else if val > root.Val {
			root.Right = insert(val, root.Right)
		} else if val < root.Val {
			root.Left = insert(val, root.Left)
		}
		return root
	}

	for i := 1; i < len(preorder); i++ {
		insert(preorder[i], root)
	}
	return root
}

/*给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
	int val;
	Node *left;
	Node *right;
	Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL。

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

示例：
    1
   / \
  2   3
 / \ / \
4  5 6  7

输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestConnect(t *testing.T) {
	type Node struct {
		Val   int
		Left  *Node
		Right *Node
		Next  *Node
	}

	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right = &Node{Val: 3}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}

	var connect func(*Node) *Node
	connect = func(root *Node) *Node {
		if root == nil {
			return root
		}
		// 每次循环从该层的最左侧节点开始
		for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
			// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
			for node := leftmost; node != nil; node = node.Next {
				// 左节点指向右节点
				node.Left.Next = node.Right

				// 右节点指向下一个左节点
				if node.Next != nil {
					node.Right.Next = node.Next.Left
				}
			}
		}
		// 返回根节点
		return root
	}

	res := connect(root)
	fmt.Println(res)
}

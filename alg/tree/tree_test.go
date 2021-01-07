package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"math"
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

//https://leetcode-cn.com/problems/number-of-islands/solution/number-of-islands-shen-du-you-xian-bian-li-dfs-or-/

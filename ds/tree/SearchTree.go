package tree

import (
	. "dp/tools"
	"fmt"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type BiSearchTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func (bst *BiSearchTree) Add(data int) {
	bst.create = new(TreeNode)
	bst.create.Val = data

	if !bst.IsEmpty() {
		bst.cur = bst.root
		for {
			if data < bst.cur.Val {
				//如果要插入的值比当前节点的值小，则当前节点指向当前节点的左孩子，如果
				//左孩子为空，就在这个左孩子上插入新值
				if bst.cur.Left == nil {
					bst.cur.Left = bst.create
					bst.create.Parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.Left
				}

			} else if data > bst.cur.Val {
				//如果要插入的值比当前节点的值大，则当前节点指向当前节点的右孩子，如果
				//右孩子为空，就在这个右孩子上插入新值
				if bst.cur.Right == nil {
					bst.cur.Right = bst.create
					bst.create.Parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.Right
				}

			} else {
				//如果要插入的值在树中已经存在，则退出
				return
			}
		}

	} else {
		bst.root = bst.create
		bst.root.Parent = nil
	}
}

func (bst *BiSearchTree) Delete(data int) {
	var (
		deleteNode func(node *TreeNode)
		node       = bst.Search(data)
	)

	deleteNode = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			//如果要删除的节点没有孩子，直接删掉它就可以(毫无挂念~.~!)
			if node == bst.root {
				bst.root = nil
			} else {
				if node.Parent.Left == node {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
			}

		} else if node.Left != nil && node.Right == nil {
			//如果要删除的节点只有左孩子或右孩子，让这个节点的父节点指向它的指针指向它的
			//孩子即可
			if node == bst.root {
				node.Left.Parent = nil
				bst.root = node.Left
			} else {
				node.Left.Parent = node.Parent
				if node.Parent.Left == node {
					node.Parent.Left = node.Left
				} else {
					node.Parent.Right = node.Left
				}
			}

		} else if node.Left == nil && node.Right != nil {
			if node == bst.root {
				node.Right.Parent = nil
				bst.root = node.Right
			} else {
				node.Right.Parent = node.Parent
				if node.Parent.Left == node {
					node.Parent.Left = node.Right
				} else {
					node.Parent.Right = node.Right
				}
			}

		} else {
			//如果要删除的节点既有左孩子又有右孩子，就把这个节点的直接后继的值赋给这个节
			//点，然后删除直接后继节点即可
			successor := bst.GetSuccessor(node.Val)
			node.Val = successor.Val
			deleteNode(successor)
		}
	}
	deleteNode(node)
}
func (bst BiSearchTree) GetRoot() *TreeNode {
	if bst.root != nil {
		return bst.root
	}
	return nil
}

func (bst BiSearchTree) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

func (bst BiSearchTree) InOrderTravel() {
	var inOrderTravel func(node *TreeNode)

	inOrderTravel = func(node *TreeNode) {
		if node != nil {
			inOrderTravel(node.Left)
			fmt.Printf("%d ", node.Val)
			inOrderTravel(node.Right)
		}
	}

	inOrderTravel(bst.root)
}

func (bst BiSearchTree) PreOrderTravel() {
	var PreOrderTravel func(node *TreeNode)

	PreOrderTravel = func(node *TreeNode) {
		if node != nil {
			fmt.Printf("%d ", node.Val)
			PreOrderTravel(node.Left)
			PreOrderTravel(node.Right)
		}
	}

	PreOrderTravel(bst.root)
}

func (bst BiSearchTree) PostOrderTravel() {
	var PostOrderTravel func(node *TreeNode)

	PostOrderTravel = func(node *TreeNode) {
		if node != nil {
			PostOrderTravel(node.Left)
			PostOrderTravel(node.Right)
			fmt.Printf("%d ", node.Val)
		}
	}

	PostOrderTravel(bst.root)
}

func (bst BiSearchTree) Search(data int) *TreeNode {
	//和Add操作类似，只要按照比当前节点小就往左孩子上拐，比当前节点大就往右孩子上拐的思路
	//一路找下去，知道找到要查找的值返回即可
	bst.cur = bst.root
	for {
		if bst.cur == nil {
			return nil
		}

		if data < bst.cur.Val {
			bst.cur = bst.cur.Left
		} else if data > bst.cur.Val {
			bst.cur = bst.cur.Right
		} else {
			return bst.cur
		}
	}
}

func (bst BiSearchTree) GetDeepth() int {
	var getDepth func(node *TreeNode) int

	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		lDepth, rDepth := getDepth(node.Left), getDepth(node.Right)
		if lDepth > rDepth {
			return lDepth + 1
		} else {
			return rDepth + 1
		}
	}

	return getDepth(bst.root)
}

func (bst BiSearchTree) GetMin() int {
	//根据二叉查找树的性质，树中最左边的节点就是值最小的节点
	if bst.root == nil {
		return -1
	}
	bst.cur = bst.root
	for {
		if bst.cur.Left != nil {
			bst.cur = bst.cur.Left
		} else {
			return bst.cur.Val
		}
	}
}

func (bst BiSearchTree) GetMax() int {
	//根据二叉查找树的性质，树中最右边的节点就是值最大的节点
	if bst.root == nil {
		return -1
	}
	bst.cur = bst.root
	for {
		if bst.cur.Right != nil {
			bst.cur = bst.cur.Right
		} else {
			return bst.cur.Val
		}
	}
}

func (bst BiSearchTree) GetPredecessor(data int) *TreeNode {
	getMax := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.Right != nil {
				node = node.Right
			} else {
				return node
			}
		}
	}

	node := bst.Search(data)
	if node != nil {
		if node.Left != nil {
			//如果这个节点有左孩子，那么它的直接前驱就是它左子树的最右边的节点，因为比这
			//个节点值小的节点都在左子树，而这些节点中值最大的就是这个最右边的节点
			return getMax(node.Left)
		} else {
			//如果这个节点没有左孩子，那么就沿着它的父节点找，知道某个父节点的父节点的右
			//孩子就是这个父节点，那么这个父节点的父节点就是直接前驱
			for {
				if node == nil || node.Parent == nil {
					break
				}
				if node == node.Parent.Right {
					return node.Parent
				}
				node = node.Parent
			}
		}
	}

	return nil
}

func (bst BiSearchTree) GetSuccessor(data int) *TreeNode {
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.Left != nil {
				node = node.Left
			} else {
				return node
			}
		}
	}

	//参照寻找直接前驱的函数对比着看
	node := bst.Search(data)
	if node != nil {
		if node.Right != nil {
			return getMin(node.Right)
		} else {
			for {
				if node == nil || node.Parent == nil {
					break
				}
				if node == node.Parent.Left {
					return node.Parent
				}
				node = node.Parent
			}
		}
	}

	return nil
}

func (bst *BiSearchTree) Clear() {
	bst.root = nil
	bst.cur = nil
	bst.create = nil
}

/*
			15
		   /  \
		  6    18
		 / \   / \
		3   7 17  20
	   / \   \
	  2   4   13
			  / \
			 9   14
*/

func InitTree(arr []int) (bst BiSearchTree) {
	for _, v := range arr {
		bst.Add(v)
	}
	return
}

func Run(bst BiSearchTree) {
	fmt.Println("[in]The nodes of the BiSearchTree is: ")
	bst.InOrderTravel()
	fmt.Println()
	fmt.Println("[post]The nodes of the BiSearchTree is: ")
	bst.PostOrderTravel()
	fmt.Println()
	fmt.Println("[pre]The nodes of the BiSearchTree is: ")
	bst.PreOrderTravel()
	fmt.Println()

	fmt.Printf("The deepth of the tree is: %d\n", bst.GetDeepth())
	fmt.Printf("The min is: %d\n", bst.GetMin())
	fmt.Printf("The max is: %d\n", bst.GetMax())

	if bst.Search(17) != nil {
		fmt.Printf("The 17 exists.\n")
	}

	fmt.Printf("root node's predecessor is: %d\n", bst.GetPredecessor(bst.GetRoot().Val).Val)
	fmt.Printf("root node's successor is: %d\n", bst.GetSuccessor(bst.GetRoot().Val).Val)

	bst.Delete(13)
	fmt.Printf("Nodes after delete the 13: ")
	bst.InOrderTravel()
	fmt.Printf("\n")
}

func IsBalancedBottom(root *TreeNode) bool {
	return HeightBottom(root) >= 0
}

func HeightBottom(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := HeightBottom(root.Left)
	right := HeightBottom(root.Right)
	if left == -1 || right == -1 || Abs(left-right) > 1 {
		return -1
	}
	return Max(left, right) + 1
}

func IsBalancedTop(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsBalancedTop(root.Left) && IsBalancedTop(root.Right) && Abs(HeightTop(root.Left)-HeightTop(root.Right)) < 2
}

func HeightTop(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return Max(HeightTop(node.Right), HeightTop(node.Left)) + 1
}

func PreOrderLabel(root *TreeNode) (res []int) {
	type Node struct {
		isCur bool
		node  *TreeNode
	}
	// 根节点入栈，标记为不记录结果，继续遍历
	stack := []*Node{{true, root}}
	for len(stack) > 0 {
		// 处理栈中弹出的最后一个节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果节点为空则跳过
		if cur.node == nil {
			continue
		}
		if cur.isCur {
			stack = append(stack, &Node{true, cur.node.Right})
			stack = append(stack, &Node{true, cur.node.Left})
			stack = append(stack, &Node{false, cur.node}) // 标记下一个添加结果集的节点
		} else {
			res = append(res, cur.node.Val)
		}
	}
	return
}

func InOrderLabel(root *TreeNode) (res []int) {
	type Node struct {
		isCur bool
		node  *TreeNode
	}
	// 根节点入栈，标记为不记录结果，继续遍历
	stack := []*Node{{true, root}}
	for len(stack) > 0 {
		// 处理栈中弹出的最后一个节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果节点为空则跳过
		if cur.node == nil {
			continue
		}
		if cur.isCur {
			stack = append(stack, &Node{true, cur.node.Right})
			stack = append(stack, &Node{false, cur.node}) // 标记下一个添加结果集的节点
			stack = append(stack, &Node{true, cur.node.Left})
		} else {
			res = append(res, cur.node.Val)
		}
	}
	return
}

func PostOrderLabel(root *TreeNode) (res []int) {
	type Node struct {
		isCur bool
		node  *TreeNode
	}
	// 根节点入栈，标记为不记录结果，继续遍历
	stack := []*Node{{true, root}}
	for len(stack) > 0 {
		// 处理栈中弹出的最后一个节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果节点为空则跳过
		if cur.node == nil {
			continue
		}
		if cur.isCur {
			stack = append(stack, &Node{false, cur.node}) // 标记下一个添加结果集的节点
			stack = append(stack, &Node{true, cur.node.Right})
			stack = append(stack, &Node{true, cur.node.Left})
		} else {
			res = append(res, cur.node.Val)
		}
	}
	return
}

func PreOrderRecur(root *TreeNode) (res []int) {
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root != nil {
			res = append(res, root.Val)
			preOrder(root.Left)
			preOrder(root.Right)
		}
	}
	preOrder(root)
	return
}

func InOrderRecur(root *TreeNode) (res []int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			res = append(res, node.Val)
			inOrder(node.Right)
		}
	}
	inOrder(root)
	return
}

func PostOrderRecur(root *TreeNode) (res []int) {
	var postOrder func(*TreeNode)
	postOrder = func(root *TreeNode) {
		if root != nil {
			postOrder(root.Left)
			postOrder(root.Right)
			res = append(res, root.Val)
		}
	}
	postOrder(root)
	return
}

func PreOrderLoop(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 添加根节点到结果集
			res = append(res, root.Val)
			// 根节点入栈，直到最左边的叶子节点
			stack = append(stack, root)
			root = root.Left
		}
		// 切换到最后一个节点的右子节点
		root = stack[len(stack)-1].Right
		// 弹出栈中最后一个元素
		stack = stack[:len(stack)-1]
	}
	return
}

func InOrderLoop(root *TreeNode) []int {
	res, stack := []int{}, []*TreeNode{}
	// 如果根节点不为空，并且栈中有元素
	for root != nil || len(stack) > 0 {
		// 遍历到最左边的叶子节点，并一直添加左子节点到栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出栈中最后一个元素，等价于root = stack.pop（）
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 将最后一个元素添加到结果集
		res = append(res, root.Val)
		// 切换到右子树
		root = root.Right
	}
	return res
}

func PostOrderLoop(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

func GetRootExample() (roots []*TreeNode) {
	// [1 3 2]
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}

	// [2 1]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}

	// [1 2]
	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	return append(roots, []*TreeNode{root1, root2, root3}...)
}

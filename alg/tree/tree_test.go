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

func inOrder() {
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}
	// [1 3 2]

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	// [2 1]

	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	// [1 2]

	fmt.Println(InOrderLabel(root1))

	fmt.Println(InOrderTraversal1(root1))
	//fmt.Println(InOrderTraversal1(root2))
	//fmt.Println(InOrderTraversal1(root3))

	fmt.Println(InOrderTraversal2(root1))
	//fmt.Println(InOrderTraversal2(root2))
	//fmt.Println(InOrderTraversal2(root3))
	//
	fmt.Println(InOrderTraversalLoop(root1))
	//fmt.Println(InOrderTraversalLoop(root2))
	//fmt.Println(InOrderTraversalLoop(root3))
}

func preOrder() {
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}
	// [1 3 2]

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	// [2 1]

	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	// [1 2]

	fmt.Println(PreOrderRecursion(root1))
	fmt.Println(PreOrderRecursion(root2))
	fmt.Println(PreOrderRecursion(root3))

	fmt.Println(PreOrderLoop(root1))
	fmt.Println(PreOrderLoop(root2))
	fmt.Println(PreOrderLoop(root3))
}

func postOrder() {
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}
	// [1 3 2]

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	// [2 1]

	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	// [1 2]

	fmt.Println(InOrderLabel(root1))
	fmt.Println(PostOrderRecursion(root1))
	//fmt.Println(PostOrderRecursion(root2))
	//fmt.Println(PostOrderRecursion(root3))

	fmt.Println(PostOrderLoop(root1))
	//fmt.Println(PostOrderLoop(root2))
	//fmt.Println(PostOrderLoop(root3))
}

func PostOrderRecursion(root *TreeNode) (res []int) {
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

func InOrderLoop(root *TreeNode) []int {
	return nil
}

func PreOrderRecursion(root *TreeNode) (res []int) {
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

func TestGo(t *testing.T) {
	//preOrder()
	postOrder()
	//inOrder()
}

package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"strconv"
	"strings"
)

const PRE = 0
const IN = 1
const POST = 2

func Serialize(root *TreeNode) string {
	// 创建节点队列
	queue := []*TreeNode{root}
	values := []string{}
	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			values = append(values, strconv.Itoa(cur.Val))
			// 入队
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		} else { // 如果空节点则设置为nil
			values = append(values, "nil")
		}
	}
	return "[" + strings.Join(values, ",") + "]"
}

func Deserialize(data string) (root *TreeNode) {
	if data == "[]" {
		return
	}
	data = data[1 : len(data)-1]
	i, values := 1, strings.Split(data, ",")
	// 用第一个元素创建根节点
	root = &TreeNode{Val: ToInt(values[0])}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		//  取出头节点
		cur := queue[0]
		// 出队
		queue = queue[1:]
		// 左子节点入队
		if values[i] != "nil" {
			cur.Left = &TreeNode{Val: ToInt(values[i])}
			queue = append(queue, cur.Left)
		}
		i++
		// 右子节点入队
		if values[i] != "nil" {
			cur.Right = &TreeNode{Val: ToInt(values[i])}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return
}

func IsBalancedBottom(root *TreeNode) bool {
	// =0的情况针对数组为空
	return HeightBottom(root) >= 0
}

func HeightBottom(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := HeightBottom(root.Left), HeightBottom(root.Right)
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

func Ordering(root *TreeNode, functions []func(*TreeNode) []int) {
	for _, fun := range functions {
		fmt.Println(fun(root))
	}
}

// 根左右
func PreOrderLoop(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 走到最左子节点，递推时记录结果，右子树入栈
		if root != nil {
			// 添加根节点到结果集
			res = append(res, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		} else {
			// 回归
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return
}

// 左根右
func InOrderLoop(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	// 如果根节点不为空，并且栈中有元素
	for root != nil || len(stack) > 0 {
		// 递推，左子树入栈
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			// 回归时，先记录左子树，后记录根
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 将最后一个元素添加到结果集
			res = append(res, root.Val)
			// 切换到右子树
			root = root.Right
		}
	}
	return
}

// 左右根
func PostOrderLoop(root *TreeNode) (res []int) {
	deque := []*TreeNode{}
	for len(deque) > 0 || root != nil {
		if root != nil {
			// 新元素添加到队列头，根节点位于队尾
			res = append([]int{root.Val}, res...)
			if root.Left != nil {
				deque = append(deque, root.Left)
			}
			root = root.Right // 根节点靠哪个孩子近就往哪边移动
		} else {
			root = deque[len(deque)-1]
			deque = deque[:len(deque)-1]
		}
	}
	return
}

func BFS(root *TreeNode) (res []int) {
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

func GetRootForTraverse() (roots []*TreeNode) {
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

func GetRootForSerialize() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	return root
}

func BFSArray(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// tmp记录当前层的值，size记录当前层的大小
		tmp, size := []int{}, len(queue)
		for i := 0; i < size; i++ {
			// 取出第一个
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

func IslandBFS(grid [][]string) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 遇到第一片陆地，则查看周边是否也是陆地
			if grid[i][j] == "1" {
				bfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]string, i int, j int) {
	queue := [][]int{[]int{i, j}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur[0], cur[1]
		// 如果不超出边界，并且也是陆地，则沉掉当前陆地
		if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && grid[i][j] == "1" {
			grid[i][j] = "0"
			// 继续检查周边是否也是陆地
			queue = append(queue, []int{i + 1, j})
			queue = append(queue, []int{i - 1, j})
			queue = append(queue, []int{i, j + 1})
			queue = append(queue, []int{i, j - 1})
		}
	}
}

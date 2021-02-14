package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"math"
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
func PostOrderLoop(root *TreeNode) (queue []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			// 新元素添加到队列头，根节点位于队尾
			queue = append([]int{root.Val}, queue...)
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			root = root.Right // 根节点靠哪个孩子近就往哪边移动
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
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

func RecoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pre *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && root.Val < pre.Val {
				y = root
				if x == nil {
					x = pre
				} else {
					break
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func RecoverTreeMorris1(root *TreeNode) {
	if root == nil {
		return
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
		}
		cur = cur.Right
	}
}

func RecoverTreeMorris(root *TreeNode) {
	var x, y, last *TreeNode
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
		}
		if last != nil && cur.Val < last.Val {
			if x == nil {
				x = last
			}
			y = cur
		}
		last = cur
		cur = cur.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

func BFStoArray(root *TreeNode) (res []interface{}) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			res = append(res, nil)
		} else {
			res = append(res, cur.Val)
		}
		if cur == nil {
			continue
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		} else {
			queue = append(queue, nil)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		} else {
			queue = append(queue, nil)
		}
	}
	res = Prune(res)
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

func MyPreOrder(root *TreeNode) (res []int) {
	return
}

func MyInOrder(root *TreeNode) (res []int) {
	return
}
func MyPostOrder(root *TreeNode) (res []int) {
	return
}

func AllOrderTraverse() {
	for _, v := range []int{PRE, IN, POST} {
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

func isPerfectSquareOld(num int) interface{} {
	if num < 2 {
		return true
	}
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func isPerfectSquare(n int) interface{} {
	start := 1
	for n > 0 {
		n -= start
		start += 2
	}
	return n == 0
}

func search(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	// 如果只有一个数字
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 当右半边递增
			//如果目标值在右半边
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 当左半边递增
			//如果目标值在右半边
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(MaxDepth(root.Right), MaxDepth(root.Left)) + 1
}

func LevelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, level)
	}
	size := len(res)
	for i := 0; i < size>>1; i++ {
		res[i], res[size-1-i] = res[size-1-i], res[i]
	}
	return res
}

func recoverT(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pre *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && root.Val < pre.Val {
				y = root
				if x == nil {
					x = pre
				} else {
					break
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
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

func ReverseTreeL(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cur.Right, cur.Left = cur.Left, cur.Right
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return root
}

func ReverseTreeR(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	ReverseTreeR(root.Right)
	ReverseTreeR(root.Left)
	return root
}

func VerifyPostOrder(postOrder []int) bool {
	stack, root := []int{}, math.MaxInt64
	// 根 右 左
	for i := len(postOrder) - 1; i >= 0; i-- {
		// 如果当前后序节点大于根节点则不符合定义
		if postOrder[i] > root {
			return false
		}
		// 如果栈不为空，且栈中最后一个元素大于当前元素(它必然是当前元素最近的右子节点)，更新根节点为右子节点
		for len(stack) > 0 && stack[len(stack)-1] > postOrder[i] {
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
		// 每次循环都会把当前元素入栈
		stack = append(stack, postOrder[i])
	}
	return true
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

func SortedArrayToBST(nums []int) *TreeNode {
	var find func([]int, int, int) *TreeNode
	find = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		root := &TreeNode{Val: nums[mid]}
		root.Left = find(nums, start, mid-1)
		root.Right = find(nums, mid+1, end)
		return root
	}
	return find(nums, 0, len(nums)-1)
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

func ClosestValue(root *TreeNode, target float64) int {
	l, r := float64(root.Val), float64(root.Val)
	for root != nil {
		if float64(root.Val) > target {
			l = float64(root.Val)
			root = root.Left
		} else if float64(root.Val) < target {
			r = float64(root.Val)
			root = root.Right
		} else {
			return root.Val
		}
	}
	if math.Abs(l-target) > math.Abs(r-target) {
		return int(r)
	}
	return int(l)
}

func DiameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var cal func(*TreeNode) int
	cal = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := cal(root.Left)
		right := cal(root.Right)
		res = Max(res, left+right)
		return Max(left, right) + 1
	}
	cal(root)
	return res
}

func MergeTreesR(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = MergeTreesR(t1.Left, t2.Left)
	t1.Right = MergeTreesR(t1.Right, t2.Right)
	return t1
}

func MergeTreesL(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	stack := [][2]*TreeNode{[2]*TreeNode{t1, t2}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 两颗树的当前节点都不为空
		if cur[0] != nil && cur[1] != nil {
			cur[0].Val += cur[1].Val
		}

		if cur[0].Right != nil && cur[1].Right != nil {
			stack = append(stack, [2]*TreeNode{cur[0].Right, cur[1].Right})
		}
		if cur[0].Left != nil && cur[1].Left != nil {
			stack = append(stack, [2]*TreeNode{cur[0].Left, cur[1].Left})
		}

		if cur[0].Left == nil && cur[1].Left != nil {
			cur[0].Left = cur[1].Left
		}
		if cur[0].Right == nil && cur[1].Right != nil {
			cur[0].Right = cur[1].Right
		}
	}
	return t1
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

func IsSymmetricL(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left)
		queue = append(queue, r.Right)
		queue = append(queue, l.Right)
		queue = append(queue, r.Left)
	}
	return true
}

func IsSymmetricR(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
	}
	return check(root, root)
}

func BinaryTreePathsR(root *TreeNode) []string {
	res := []string{}
	var build func(*TreeNode, string)
	build = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		if root.Right == nil && root.Left == nil {
			path += strconv.Itoa(root.Val)
			res = append(res, path)
			return
		}
		path += strconv.Itoa(root.Val) + "->"
		build(root.Left, path)
		build(root.Right, path)
	}
	build(root, "")
	return res
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

func FindSecondMinimumValue(root *TreeNode) int {
	var find func(*TreeNode, int) int
	find = func(root *TreeNode, mi int) int {
		if root == nil {
			return -1
		}
		if root.Val > mi {
			return root.Val
		}
		left := find(root.Left, mi)
		right := find(root.Right, mi)
		if left == -1 {
			return right
		}
		if right == -1 {
			return left
		}
		return Min(right, left)
	}
	return find(root, root.Val)
}

func MinDiffInBSTL(root *TreeNode) int {
	mi, pre := math.MaxInt64, -1
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != -1 {
			mi = Min(mi, root.Val-pre)
		}
		pre = root.Val
		root = root.Right
	}
	return mi
}

func MinDiffInBSTR(root *TreeNode) int {
	mi, pre := math.MaxInt64, -1
	var find func(root *TreeNode)
	find = func(root *TreeNode) {
		if root == nil {
			return
		}
		find(root.Left)
		if pre != -1 {
			mi = Min(mi, root.Val-pre)
		}
		pre = root.Val
		find(root.Right)
	}
	find(root)
	return mi
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	var check func(*TreeNode, *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return check(s, t) || IsSubtree(s.Right, t) || IsSubtree(s.Left, t)
}

func Tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Right == nil && root.Left == nil {
		return strconv.Itoa(root.Val)
	}
	if root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + Tree2str(root.Left) + ")"
	}
	return strconv.Itoa(root.Val) + "(" + Tree2str(root.Left) + ")(" + Tree2str(root.Right) + ")"
}

func Tree2str2(t *TreeNode) string {
	if t == nil {
		return ""
	}
	builder := strings.Builder{}

	var construct func(*TreeNode)
	construct = func(t *TreeNode) {
		if t == nil {
			return
		}
		builder.WriteString(strconv.Itoa(t.Val))
		if t.Left == nil && t.Right == nil {
			return
		}
		builder.WriteString("(")
		construct(t.Left)
		builder.WriteString(")")
		if t.Right != nil {
			builder.WriteString("(")
			construct(t.Right)
			builder.WriteString(")")
		}
	}
	construct(t)
	return builder.String()
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
		}
		update(cur.Val)
		cur = cur.Right
	}
	return
}

func GetMinimumDifference1(root *TreeNode) int {
	mi, preVal := math.MaxInt64, -1
	if root == nil {
		return 0
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
		}
		mi = Min(mi, cur.Val-preVal)
		preVal = cur.Val
		cur = cur.Right
	}

	return mi
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GetMinimumDifference2(root *TreeNode) int {
	maxOrMin := func(node *TreeNode, max bool) *TreeNode {
		cur := node
		last := node
		for cur != nil {
			last = cur
			if max {
				cur = cur.Right
			} else {
				cur = cur.Left
			}
		}
		return last
	}

	r := math.MaxInt32
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		min := maxOrMin(node.Left, true)
		if min != nil {
			//左侧最大值
			tmp := node.Val - min.Val
			if tmp < r {
				r = tmp
			}
		}
		max := maxOrMin(node.Right, false)
		if max != nil {
			//右侧最小值
			tmp := max.Val - node.Val
			if tmp < r {
				r = tmp
			}
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return r
}

func AverageOfLevels(root *TreeNode) []float64 {
	stack := []*TreeNode{root}
	res := []float64{}
	for len(stack) > 0 {
		size := len(stack)
		levelSum := 0
		for i := 0; i < size; i++ {
			cur := stack[0]
			stack = stack[1:]
			levelSum += cur.Val
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
		res = append(res, float64(levelSum)/float64(size))
	}
	return res
}

func FindTilt(root *TreeNode) int {
	res := 0
	var bfs func(*TreeNode) int
	bfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := bfs(root.Left)
		r := bfs(root.Right)
		res += Abs(l - r)
		return root.Val + l + r
	}
	bfs(root)
	return res
}

func MinDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}
	return depth
}

func MinDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	mi := math.MaxInt64
	if root.Left != nil {
		mi = Min(MinDepth1(root.Left), mi)
	}
	if root.Right != nil {
		mi = Min(MinDepth1(root.Right), mi)
	}
	return mi + 1
}

func SumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			if cur.Left.Left == nil && cur.Left.Right == nil {
				res += cur.Left.Val
			}
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}

func SumOfLeftLeavesDFS(root *TreeNode) int {
	res := 0
	var counter func(*TreeNode)
	counter = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				res += root.Left.Val
			}
			counter(root.Left)
		}
		if root.Right != nil {
			counter(root.Right)
		}
	}
	counter(root)
	return res
}

func MaxPathSum(root *TreeNode) int {
	ma := math.MinInt64
	var find func(*TreeNode) int
	find = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := Max(find(root.Left), 0)
		r := Max(find(root.Right), 0)
		cur := root.Val + l + r
		ma = Max(cur, ma)
		return root.Val + Max(l, r)
	}
	find(root)
	return ma
}

func ClosestKValues(root *TreeNode, target float64, k int) []int {
	res := []int{}

	getRes := func(val int) {
		if len(res) < k {
			res = append(res, val)
		} else if AbsFloat(float64(res[0])-float64(target)) > AbsFloat(float64(val)-float64(target)) {
			res = res[1:]
			res = append(res, val)
		}
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
		}
		getRes(cur.Val)
		cur = cur.Right
	}

	return res
}

type Codec struct {
	l []string
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) SerializeDFS(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	left := this.SerializeDFS(root.Left)
	right := this.SerializeDFS(root.Right)
	return strconv.Itoa(root.Val) + "," + left + "," + right
}

// Deserializes your encoded data to tree.
func (this *Codec) DeserializeDFS(data string) *TreeNode {
	arr := strings.Split(data, ",")
	var build func(*[]string) *TreeNode
	build = func(arr *[]string) *TreeNode {
		cur := (*arr)[0]
		*arr = (*arr)[1:]
		if cur == "nil" {
			return nil
		}
		val, _ := strconv.Atoi(cur)
		node := &TreeNode{Val: val}
		node.Left = build(arr)
		node.Right = build(arr)
		return node
	}
	return build(&arr)
}

func (this *Codec) SerializeBFS(root *TreeNode) string {
	res := []string{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur != nil {
			res = append(res, strconv.Itoa(cur.Val))
			q = append(q, cur.Left)
			q = append(q, cur.Right)
		} else {
			res = append(res, "nil")
		}
	}
	return strings.Join(res, ",")
}
func (this *Codec) DeserializeBFS(data string) *TreeNode {
	if data == "nil" {
		return nil
	}
	arr := strings.Split(data, ",")
	val, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: val}
	index, q := 1, []*TreeNode{root}
	for index < len(arr) {
		cur := q[0]
		q = q[1:]
		l, r := arr[index], arr[index+1]
		if l != "nil" {
			v, _ := strconv.Atoi(l)
			cur.Left = &TreeNode{Val: v}
			q = append(q, cur.Left)
		}
		if r != "nil" {
			v, _ := strconv.Atoi(r)
			cur.Right = &TreeNode{Val: v}
			q = append(q, cur.Right)
		}
		index += 2
	}
	return root
}

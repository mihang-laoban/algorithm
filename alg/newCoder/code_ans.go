package newCoder

import (
	"dp/ds/linkedList"
	"dp/ds/tree"
	"dp/tools"
	"math"
	"strconv"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r-l)>>1 + l
		if nums[m] == target {
			for m >= 0 {
				if nums[m] != target {
					break
				}
				m--
			}
			return m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

	}
	return -1
}

func ClimbStairsDP(n int) int {
	dp := make([]int, n+1)
	return ClimbStairsDP_(n, dp)
}

func ClimbStairsDP_(n int, dp []int) int {
	if dp[n] > 0 {
		return dp[n]
	}
	if n == 1 {
		dp[1] = 1
	} else if n == 2 {
		dp[2] = 2
	} else {
		dp[n] = ClimbStairsDP_(n-1, dp) + ClimbStairsDP_(n-2, dp)
	}
	return dp[n]
}

func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
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

func MinPathSum(matrix [][]int) int {
	col, row := len(matrix), len(matrix[0])
	dp := make([][]int, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < col; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < row; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			dp[i][j] = tools.Min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[col-1][row-1]
}

func OddEvenList(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	odd, even := head, head.Next
	for second != nil && second.Next != nil {
		first.Next = first.Next.Next
		first = first.Next
		second.Next = second.Next.Next
		second = second.Next
	}
	first.Next = even
	return odd
}

func GetLongestPalindrome(str string, size int) int {
	res := 1
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}
	for r := 1; r < size; r++ {
		for l := r - 1; l >= 0; l-- {
			if r-l < 2 {
				dp[l][r] = str[r] == str[l]
				res = tools.Max(res, r-l+1)
				continue
			}
			dp[l][r] = dp[l+1][r-1] && str[r] == str[l]
			if dp[l][r] {
				res = tools.Max(res, r-l+1)
			}
		}
	}
	return res
}

func MaxInWindows(num []int, size int) []int {
	res := []int{}
	if size == 0 || len(num) == 0 {
		return res
	}
	l, r := 0, len(num)
	for l+size <= r {
		max := find(num[l : l+size])
		res = append(res, max)
		l++
	}
	return res
}

func find(num []int) int {
	max := math.MinInt64
	for i := 0; i < len(num); i++ {
		if max < num[i] {
			max = num[i]
		}
	}
	return max
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

func LIS(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	maxLen := make([]int, len(arr))
	maxLen[0] = 1
	res := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		// 如果当前元素比结果集里面最后一个元素大，则当前元素也加入结果集
		if arr[i] > res[len(res)-1] {
			res = append(res, arr[i])
			maxLen[i] = len(res) // 记录当前的最大长度
		} else {
			l, r := 0, len(res)-1
			for l < r {
				mid := l + (r-l)>>1
				if res[mid] < arr[i] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			res[l] = arr[i]
			maxLen[i] = l + 1
		}
	}
	for i, j := len(maxLen)-1, len(res); j > 0; i-- {
		if maxLen[i] == j {
			j--
			res[j] = arr[i]
		}
	}
	return res
}

func MaxSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m < 2 {
		return m
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = tools.Min(tools.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			max = tools.Max(dp[i][j], max)
		}
	}
	return max * max
}

func KthNode(pRoot *tree.TreeNode, k int) *tree.TreeNode {
	// write code here
	stack := []*tree.TreeNode{}
	for pRoot != nil || len(stack) > 0 {
		for pRoot != nil {
			stack = append(stack, pRoot)
			pRoot = pRoot.Left
		}
		pRoot = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k-1 == 0 {
			return pRoot
		}
		k--
		pRoot = pRoot.Right
	}
	return nil
}

func Kmp(S string, T string) int {
	count, sSize, tSize := 0, len(S), len(T)
	if len(S) > len(T) {
		return 0
	}
	for i := 0; i < tSize-sSize+1; i++ {
		j := 0
		for ; j < sSize; j++ {
			if S[j] != T[i+j] {
				break
			}
		}
		if j == sSize {
			count++
		}
	}
	return count
}

func JumpFloor(number int) int {
	// write code here
	if number < 3 {
		return number
	}
	a, b := 1, 2
	for i := 3; i <= number+1; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	minH, minW := 0, 0
	maxH, maxW := len(matrix)-1, len(matrix[0])-1
	for minH <= maxH && minW <= maxW {
		for i := minW; i <= maxW; i++ {
			res = append(res, matrix[minH][i])
		}
		for i := minH + 1; i <= maxH; i++ {
			res = append(res, matrix[i][maxW])
		}
		for i := maxW - 1; i >= minW && maxH > minH; i-- {
			res = append(res, matrix[maxH][i])
		}
		for i := maxH - 1; i > minH && maxW > minW; i-- {
			res = append(res, matrix[i][minW])
		}
		minW++
		maxW--
		minH++
		maxH--
	}
	return res
}

func NumberOf1(n int) int {
	res := 0
	for n != 0 {
		n = n & int((int32(n) - 1))
		res++
	}
	return res
}

func FindElement(mat [][]int, n int, m int, x int) []int {
	i, j := 0, m-1
	for i < n && j >= 0 {
		if mat[i][j] == x {
			return []int{i, j}
		} else if mat[i][j] > x {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}

func RotateMatrix(mat [][]int, n int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := n - 1; j >= 0; j-- {
			tmp = append(tmp, mat[j][i])
		}
		res = append(res, tmp)
	}
	return res
}

func Intersect(a, b []int) []int {
	recorder := map[int]int{}
	for _, v := range a {
		if recorder[v] == 1 {
			continue
		}
		recorder[v]++
	}
	for _, v := range b {
		recorder[v]++
	}
	res := []int{}
	for key, value := range recorder {
		if value > 1 {
			res = append(res, key)
		}
	}
	return res
}

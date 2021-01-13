package dp

import (
	. "dp/tools"
	"fmt"
)

func Exchange(values []int, total int) int {
	dp := make([]int, total+1)
	for i := 0; i < total+1; i++ {
		dp[i] = total + 1
	}

	dp[0] = 0
	for i := 1; i < total+1; i++ {
		for _, value := range values {
			if i-value >= 0 {
				dp[i] = Min(dp[i-value]+1, dp[i])
			}
		}
	}
	if dp[total] == total+1 {
		return -1
	}
	return dp[total]
}

func Pack(weights []int, values []int, totalWeight int, totalCount int) {
	weights = append([]int{0}, weights...)
	values = append([]int{0}, values...)
	dp := make([][]int, totalCount+1)
	for i := 0; i < totalCount+1; i++ {
		dp[i] = make([]int, totalWeight+1)
	}
	for i := 0; i < totalCount+1; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < totalWeight+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < totalCount+1; i++ {
		for j := 1; j < totalWeight+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j-weights[i] >= 0 {
				dp[i][j] = Max(dp[i][j], dp[i][j-weights[i]]+values[i])
			}
		}
	}
	fmt.Println(dp[totalCount][totalWeight])
}

func PackImproved(weights []int, values []int, totalWeight int, totalCount int) {
	weights = append([]int{0}, weights...)
	values = append([]int{0}, values...)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, totalWeight+1)
	}
	for i := 0; i < 2; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < totalWeight+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < totalCount+1; i++ {
		for j := 1; j < totalWeight+1; j++ {
			dp[i%2][j] = dp[(i-1)%2][j]
			if j-weights[i] >= 0 {
				dp[i%2][j] = Max(dp[i%2][j], dp[i%2][j-weights[i]]+values[i]) // 可重复放入
				//dp[i%2][j] = Max(dp[(i-1)%2][j], dp[(i-1)%2][j-weights[i]] + values[i]) // 只可放入一次
			}
		}
	}
	fmt.Println(dp[totalCount%2][totalWeight])
}

func LargestSubArr(str string) int {
	size := len(str)
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}
	res := 0
	for i := 0; i < size; i++ {
		dp[i][i] = true
		res++
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			dp[j][i] = str[i] == str[j] && (i-j <= 3 || dp[j+1][i-1])
			if dp[j][i] {
				res++
			}
		}
	}
	return res
}

func MaxSubArray(nums []int) int {
	cur, res := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		tmp := Max(nums[i], nums[i]+cur)
		res = Max(res, tmp)
		cur = tmp
	}
	return res
}

func FindLargestSubSeq(str string) {
	size := len(str)
	dp := InitMemo(size, size)
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if str[i] == str[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[0][size-1])
}

func SharedSubSeq(text1 string, text2 string) {
	s1 := len(text1)
	s2 := len(text2)
	dp := make([][]int, s1+1)
	for i := 0; i < s1+1; i++ {
		dp[i] = make([]int, s2+1)
	}
	for i := 1; i < s1+1; i++ {
		for j := 1; j < s2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[s1][s2])
}

func LongestCommon(text1 string, text2 string) {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, 2)
	}

	for j := 1; j <= n; j++ {
		a := j % 2
		b := (j - 1) % 2
		for i := 1; i < m; i++ {
			if text2[j-1] == text1[i-1] {
				dp[i][a] = dp[i-1][b] + 1
			} else {
				dp[i][a] = Max(dp[i-1][a], dp[i][b])
			}
		}
	}
	fmt.Println(dp[m][n%2])
}

func RobotPathSum(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func RobotPathSum2(m int, n int) int {
	pre, cur := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = 1
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = pre[j] + cur[j-1]
		}
		copy(pre, cur)
	}
	return pre[n-1]
}

func RobotPathSum3(m int, n int) int {
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}

func InitPath() (arr [][]int) {
	m, n := 3, 3
	arr = make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	arr[1][1] = 1
	return
}

func WithObstacle(arr [][]int) {
	m := len(arr)
	n := len(arr[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if arr[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}

	for i := 0; i < n; i++ {
		if arr[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if arr[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[m-1][n-1])
}

func JumpGame(nums []int) bool {
	index := len(nums) - 1
	for i := index - 1; i >= 0; i-- {
		// 前一个位置可以抵达当前位置，则当前位置更新为前一个位置
		if index <= i+nums[i] {
			index = i
		}
	}
	if index == 0 {
		return true
	}
	return false
}

func FindContinuous(arr []int) {
	res := 0
	s1 := len(arr)
	dp := make([]int, s1)
	for i := 0; i < s1; i++ {
		dp[i] = 1
	}

	for i := 1; i < s1; i++ {
		if arr[i] > arr[i-1] {
			dp[i] = 1 + dp[i-1]
		}
		res = Max(dp[i], res)
	}

	fmt.Println(res)
}

func FindIncrement(arr []int) {
	s1 := len(arr)
	dp := make([]int, s1)
	for i := 0; i < s1; i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < s1; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = Max(dp[i], 1+dp[j])
				res = Max(res, dp[i])
			}
		}
	}
	fmt.Println(res)
}

func FindBySplit(arr []int) {
	length := 1
	n := len(arr)
	if n == 0 {
		fmt.Println(0)
	}
	d := make([]int, n+1)
	d[length] = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] > d[length] {
			length++
			d[length] = arr[i]
		} else {
			l := 1
			r := length
			pos := 0
			for l <= r {
				mid := (l + r) / 2
				if d[mid] < arr[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[pos+1] = arr[i]
		}
	}
	fmt.Println(length)
}

//最长上升子序列共有几个，你该怎么解呢？
func FindIncrementNum(arr []int) {
	s1 := len(arr)
	if s1 == 0 {
		fmt.Println(0)
	}
	dp := make([]int, s1)
	count := make([]int, s1)
	for i := 0; i < s1; i++ {
		dp[i] = 1
		count[i] = 1
	}
	for i := 1; i < s1; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
	}
	maxLen := 0
	for i := 0; i < s1; i++ {
		maxLen = Max(dp[i], maxLen)
	}
	res := 0
	for i := 0; i < s1; i++ {
		if dp[i] == maxLen {
			res += count[i]
		}
	}
	fmt.Println(res)
}

func FibDp(n int) {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[n])
}

func FibSingle(n int) {
	pre := 1
	cur := 1
	for i := 3; i <= n; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	fmt.Println(cur)
}

func ClimbRecur(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return ClimbRecur(n-1) + ClimbRecur(n-2)
}

func Climb(n int) int {
	a, b := 0, 1
	for i := 0; i < n+1; i++ {
		a, b = a+b, a
	}
	return a
}

func FindLargestKArray(nums []int, k int) {
	n := len(nums)

	m := make([][]int, n+1)
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		m[i] = make([]int, k+1)
		dp[i] = make([]int, k+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < k+1; j++ {
			m[i][j] = 0
			dp[i][j] = 0
		}
	}

	for i := 1; i < n+1; i++ {
		for j := Min(i, k); j > 0; j-- {
			if i == j {
				m[i][j] = m[i-1][j-1] + nums[i-1]
				dp[i][j] = dp[i-1][j-1] + nums[i-1]
			} else {
				m[i][j] = Max(m[i-1][j], dp[i-1][j-1]) + nums[i-1]
				dp[i][j] = Max(dp[i-1][j], m[i][j])
			}
		}
	}
	fmt.Println(dp[n][k])
}

func Product1(nums []int) {
	n := len(nums)
	ma := make([]int, n)
	mi := make([]int, n)
	ma[0] = nums[0]
	mi[0] = nums[0]

	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			temp := ma[i-1]
			ma[i-1] = mi[i-1]
			mi[i-1] = temp
		}
		ma[i] = Max(nums[i], ma[i-1]*nums[i])
		mi[i] = Min(nums[i], mi[i-1]*nums[i])
		res = Max(res, ma[i])
	}
	fmt.Println(res)
}

func Product2(nums []int) {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = nums[i]
		dp[i][1] = nums[i]
	}

	for i := 1; i < n; i++ {
		// 决策求解
		dp[i][0] = Max(dp[i-1][0]*nums[i], Max(nums[i], dp[i-1][1]*nums[i]))
		dp[i][1] = Min(dp[i-1][1]*nums[i], Min(nums[i], dp[i-1][0]*nums[i]))
	}

	res := dp[0][0]
	for i := 1; i < n; i++ {
		res = Max(res, dp[i][0])
	}
	fmt.Println(res)
}

func Rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	first, second := nums[0], Max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		first, second = second, Max(first+nums[i], second)
	}
	return second
}

func Rob2(nums []int) interface{} {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[size-1]
}

func MinCostClimbingStairs(cost []int) int {
	pre, cur := 0, 0
	for i := 2; i <= len(cost)+1; i++ {
		pre, cur = cur, Min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

/*
动态规划：
dp[i]表示i组括号的所有有效组合
dp[i] = "(dp[p]的所有有效组合)+【dp[q]的组合】"，其中 1 + p + q = i , p从0遍历到i-1, q则相应从i-1到0
*/
func DpGenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	dp := make([][]string, n+1)
	dp[0], dp[1] = []string{""}, []string{"()"}
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			// 正向遍历dp备忘录
			for _, p := range dp[j] {
				// 反向遍历dp备忘录
				for _, q := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+p+")"+q)
				}
			}
		}
	}
	return dp[n]
}

// https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
func GenerateParenthesis(n int) (res []string) {
	var GenBrackets func(int, int, int, string)

	GenBrackets = func(left int, right int, pairNum int, s string) {
		if left == pairNum && right == pairNum {
			res = append(res, s)
			return
		}
		if left < pairNum {
			GenBrackets(left+1, right, pairNum, s+"(")
		}
		if right < left {
			GenBrackets(left, right+1, pairNum, s+")")
		}
	}
	GenBrackets(0, 0, n, "")
	return
}

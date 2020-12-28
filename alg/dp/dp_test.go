package dp

import (
	. "dp/tools"
	"fmt"
	"testing"
)

func TestExchangeMinCount(t *testing.T) {
	values := []int{3, 5}
	total := 11

	dp := make([]int, total+1)
	for i := 1; i < total+1; i++ {
		dp[i] = total + 1
	}
	dp[0] = 0

	for i := 1; i < total+1; i++ {
		for j := 0; j < len(values); j++ {
			if i-values[j] < 0 {
				continue
			}
			dp[i] = Min(dp[i], dp[i-values[j]]+1)
		}
	}

	if dp[total] == total+1 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[total])
	}
}

func TestDp(t *testing.T) {
	str := "aaa"
	largestSubArr(str)
}

func largestSubArr(str string) {
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

	fmt.Println(res)
}

func TestFindLargestArrSum(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 3, -5, 1, 2}

	cur := 0
	res := cur
	for i := 0; i < len(arr); i++ {
		tmp := Max(arr[i], arr[i]+cur)
		res = Max(tmp, res)
		cur = tmp
	}
	fmt.Println(res)
}

func TestFindLargestSubSeq(t *testing.T) {
	str := "asssasms"

	FindLargestSubSeq(str)
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

func TestFindLongestSharedSubSeq(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"

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

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"
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
	//
	//int[][] dp = new int[m + 1][2];
	//for (int[] row: dp) { Arrays.fill(row, 0); }
	//
	//for (int j = 1; j <= n; j++) {
	//	int a = j % 2;
	//	int b = (j - 1) % 2;
	//	for (int i = 1; i <= m; i++) {
	//		if (text2.charAt(j - 1) == text1.charAt(i - 1)) {
	//			dp[i][a] = dp[i - 1][b] + 1;
	//		} else {
	//			dp[i][a] = Math.max(dp[i - 1][a], dp[i][b]);
	//		}
	//	}
	//}
	//
	//return dp[m][n%2];
}

func TestRobotPathSum(t *testing.T) {
	m := 4
	n := 3
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
	fmt.Println(dp[m-1][n-1])
}

func TestRobotPathSumWithObstacle(t *testing.T) {
	m := 3
	n := 3
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	arr[1][1] = 1

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

func TestJump(t *testing.T) {
	arr1 := []int{2, 3, 1, 1, 6}
	arr2 := []int{4, 2, 1, 0, 0, 6}

	JumpGame(arr1)
	JumpGame(arr2)
}

func JumpGame(arr []int) {
	index := len(arr) - 1
	for i := index - 1; i >= 0; i-- {
		if i+arr[i] >= index {
			index = i
		}
	}
	if index == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}

func TestContinuousIncrementalSubSeq(t *testing.T) {
	arr1 := []int{6, 6, 6, 6, 6, 6}
	arr2 := []int{1, 3, 5, 0, 7}

	findContinuous(arr1)
	findContinuous(arr2)
}

func findContinuous(arr []int) {
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

func TestIncrementalSubSeq(t *testing.T) {
	arr1 := []int{10, 9, 1, 5, 2, 6, 66, 18}
	findIncrement(arr1)
	findIncrementNum(arr1)
	findBySplit(arr1)
}

func findIncrement(arr []int) {
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

func findBySplit(arr []int) {
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

func findIncrementNum(arr []int) {
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

func TestFib(t *testing.T) {
	n := 6
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[n])
}

func TestFibSingle(t *testing.T) {
	n := 6
	pre := 1
	cur := 1
	for i := 3; i <= n; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	fmt.Println(cur)
}

func TestFindLargestSubArray(t *testing.T) {
	nums := []int{-1, 4, -2, 3, -2, 3}
	k := 2
	// 4 + (3 + -2 + 3) = 8
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

func TestLargestProductSubArr1(t *testing.T) {
	nums := []int{2, 8, -2, 4} // [-2, 0, -1] > 0
	//	16
	n := len(nums)
	dp_max := make([]int, n)
	dp_min := make([]int, n)
	dp_max[0] = nums[0]
	dp_min[0] = nums[0]

	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			temp := dp_max[i-1]
			dp_max[i-1] = dp_min[i-1]
			dp_min[i-1] = temp
		}
		dp_max[i] = Max(nums[i], dp_max[i-1]*nums[i])
		dp_min[i] = Min(nums[i], dp_min[i-1]*nums[i])
		res = Max(res, dp_max[i])
	}
	fmt.Println(res)
}

func TestLargestProductSubArr2(t *testing.T) {
	nums := []int{2, 8, -2, 4} // [-2, 0, -1] > 0
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

/*示例1：
输入：[3, 3, 5, 0, 0, 3, 1, 4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3 - 0 = 3 。随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4 - 1 = 3 。
示例2：
输入：[1, 2, 3, 4, 5]
输出：4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。需要注意的是，你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例3：
输入：[7, 6, 4, 3, 1]
输出：0
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。*/
func TestStock(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	length := len(prices)
	dp := make([][2][3]int, length)
	// 假设第一天没有买入
	dp[0][0][0] = 0
	dp[0][0][1] = 0
	dp[0][0][2] = 0
	// 第一天不可能已卖出
	dp[0][1][0] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][1][2] = -prices[0]
	// 处理后续日期
	for i := 1; i < length; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = Max(dp[i-1][1][0]+prices[i], dp[i-1][0][1])
		dp[i][0][2] = Max(dp[i-1][1][1]+prices[i], dp[i-1][0][2])
		dp[i][1][0] = Max(dp[i-1][0][0]-prices[i], dp[i-1][1][0])
		dp[i][1][1] = Max(dp[i-1][0][1]-prices[i], dp[i-1][1][1])
		dp[i][1][2] = 0
	}
	res := Max(dp[length-1][0][1], dp[length-1][0][2])
	fmt.Println(res)
}

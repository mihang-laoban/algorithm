package dp

import (
	"testing"
)

// Q1
func TestExchangeMinCount(t *testing.T) {
	Exchange([]int{3, 5}, 11)
}

func ex(values []int, total int) {

}

/*示例：

输入：W = 5, N = 3
w = [3, 2, 1], v = [5, 2, 3]
输出：15
解释：当 i = 2 时，选取 5 次，总价值为 5 * 3 = 15。
*/
// Q2
func TestPackage(t *testing.T) {
	weights, values := []int{3, 2, 1}, []int{5, 2, 3}
	totalWeight, totalCount := 5, 3

	Pack(weights, values, totalWeight, totalCount)
	PackImproved(weights, values, totalWeight, totalCount)
}

func pk(weights []int, values []int, totalWeight int, totalCount int) {

}

// Q3
//问题：给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
func TestDp(t *testing.T) {
	str := "aaa"
	largest(str)
	LargestSubArr(str)
}

func largest(str string) {

}

// Q4
func TestFindLargestArrSum(t *testing.T) {
	FindLargest([]int{-2, 1, -3, 4, -1, 3, -5, 1, 2})
}

// Q5
//问题：给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000。
func TestFindLargestSubSeq(t *testing.T) {
	FindLargestSubSeq("asssasms") // 5
}

// Q6
//问题：给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 若这两个字符串没有公共子序列，则返回 0。
// 其中：1 ≤ text1.length ≤ 1000；1 ≤ text2.length ≤ 1000；输入的字符串只含有小写英文字符。
func TestFindLongestSharedSubSeq(t *testing.T) {
	SharedSubSeq("abcde", "ade")
}

// Q7
func TestLongestCommonSubsequence(t *testing.T) {
	LongestCommon("abcde", "ade")
}

// Q8
func TestRobotPathSum(t *testing.T) {
	RobotPathSum(4, 3)
}

/*示例：

输入：
[
[0, 0, 0],
[0, 1, 0],
[0, 0, 0]
]
输出: 2
解释：3 * 3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右*/
// Q9
func TestRobotPathSumWithObstacle(t *testing.T) {
	WithObstacle(InitPath())
}

// Q10
func TestJump(t *testing.T) {
	JumpGame([]int{2, 3, 1, 1, 6})
	JumpGame([]int{4, 2, 1, 0, 0, 6})
}

/*示例2：

输入: nums = [1, 3, 5, 0, 7]
输出: 3
解释: 最长连续递增序列是 [1, 3, 5], 长度为 3。
你会发现 [1, 3, 5, 7] 也是升序的子序列, 但它不是连续的。
因为 5 和 7 在原数组中被 0 隔开。因此，这不是原问题的答案。*/
// Q11
func TestContinuousIncrementalSubSeq(t *testing.T) {
	FindContinuous([]int{6, 6, 6, 6, 6, 6})
	FindContinuous([]int{1, 3, 5, 0, 7})
}

/*
问题：给定一个无序的整数数组 nums，找到其中最长上升子序列的长度（Longest Increasing Subsequence，LIS）。
附加条件是：可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可；你算法的时间复杂度应该为 O(n2)。

示例：

输入: nums = [10, 9, 1, 5, 2, 6, 66, 18]
输出: 4
解释: 其中一个最长的上升子序列是 [1, 2, 6, 66]，它的长度是 4。*/
// Q12
func TestIncrementalSubSeq(t *testing.T) {
	arr1 := []int{10, 9, 1, 5, 2, 6, 66, 18}
	FindIncrement(arr1)
	FindIncrementNum(arr1)
	FindBySplit(arr1)
}

// Q13
func TestFib(t *testing.T) {
	FibDp(6)
}

func TestFibSingle(t *testing.T) {
	FibSingle(6)
}

// Q14
func TestFindLargestSubArray(t *testing.T) {
	// 4 + (3 + -2 + 3) = 8
	FindLargestKArray([]int{-1, 4, -2, 3, -2, 3}, 2)
}

// Q15
func TestLargestProductSubArr1(t *testing.T) {
	nums := []int{2, 8, -2, 4} // [-2, 0, -1] > 0
	//	16
	Product1(nums)
	Product2(nums)
}

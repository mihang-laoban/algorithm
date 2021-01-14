package dp

import (
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

// Q1
func TestExchangeMinCount(t *testing.T) {
	fmt.Println(Exchange([]int{3, 5}, 11))
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
	fmt.Println(lg(str))
	fmt.Println(LargestSubArr(str))
}

func lg(str string) int {
	n := len(str)
	return n
}

// Q4
func TestFindLargestArrSum(t *testing.T) {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 3, -5, 1, 2}))
	fdLg([]int{-2, 1, -3, 4, -1, 3, -5, 1, 2})
}

func fdLg(arr []int) {

}

// Q5
//问题：给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000。
//输入："asssasms"输出：5解释：一个可能的最长回文子序列为 "sssss"，另一种可能的答案是 "asssa"。
func TestFindLargestSubSeq(t *testing.T) {
	FindLargestSubSeq("asssasms") // 5
	findSubSeq("asssasms")
}

func findSubSeq(str string) {

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
/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

示例 1：
输入：m = 3, n = 7
输出：28

示例 2：
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右

示例 3：
输入：m = 7, n = 3
输出：28

示例 4：
输入：m = 3, n = 3
输出：6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestRobotPathSum(t *testing.T) {
	fmt.Println(RobotPathSum(4, 3))
	fmt.Println(RobotPathSum2(4, 3))
	fmt.Println(RobotPathSum3(4, 3))
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

/*你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。


示例 1：
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestRob(t *testing.T) {
	fmt.Println(Rob([]int{1, 2, 3, 1}))
	fmt.Println(Rob2([]int{1, 2, 3, 1}))
}

/*数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

示例 1：
输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。

示例 2：
输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出：6
解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestClimbMinCost(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	fmt.Println(MinCostClimbingStairs(cost))
	fmt.Println(MinCostClimbing(cost))
}

func MinCostClimbing(cost []int) interface{} {
	pre, cur := 0, 0
	for i := 2; i < len(cost)+1; i++ {
		pre, cur = cur, Min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

/*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶*/

func TestClimb(t *testing.T) {
	fmt.Println(ClimbRecur(3 + 1))
	fmt.Println(Climb(3))
}

/*数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例：
输入：n = 3
输出：[
"((()))",
"(()())",
"(())()",
"()(())",
"()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestGenBrackets(t *testing.T) {
	n := 3
	fmt.Println(GenerateParenthesis(n))
	fmt.Println(DpGenerateParenthesis(n))
}

// https://leetcode-cn.com/problems/unique-paths/

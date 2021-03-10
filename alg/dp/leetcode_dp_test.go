package dp

import (
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

/*给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

示例 3：
输入：s = "a"
输出："a"

示例 4：
输入：s = "ac"
输出："a"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(LongestPalindrome1("babad"))
	fmt.Println(LongestPalindrome2("babad"))
}

/*给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [0]
输出：0

示例 4：
输入：nums = [-1]
输出：-1

示例 5：
输入：nums = [-100000]
输出：-100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMaxSubArray(t *testing.T) {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSubArray([]int{1, 2}))
}

/*给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestTrap(t *testing.T) {
	fmt.Println(Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(Trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

/*给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3

解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

示例 2：
输入：word1 = "intention", word2 = "execution"
输出：5

解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMinDistance(t *testing.T) {
	fmt.Println(2 & 1)
	fmt.Println(MinDistance("horse", "ros"))
}

func MinDistance(word1 string, word2 string) int {
	s1, s2 := len(word1), len(word2)
	dp := make([][]int, s1+1)
	for i := 0; i < s1+1; i++ {
		dp[i] = make([]int, s2+1)
	}
	for i := 0; i < s1+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < s2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < s1+1; i++ {
		for j := 1; j < s2+1; j++ {
			dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = Min(dp[i][j], dp[i-1][j-1])
			}
		}
	}
	return dp[s1][s2]
}

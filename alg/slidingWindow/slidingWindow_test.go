package slidingWindow

import (
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

/*给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

示例 4:
输入: s = ""
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(LengthOfLongestSubstring("aaxqaamksoaqpsaaa"))
	fmt.Println(LengthOfLongestSubstring2("aaxqaamksoaqpsaaa"))
}

/*给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"
*/

func TestMinWindow(t *testing.T) {
	fmt.Println(MinWindow("ADOBECODEBANC", "ABC"))
}

func MinWindow2(s string, t string) string {
	return ""
}

/*给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

示例 1：
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例 2：
输入: s1= "ab" s2 = "eidboaoo"
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestCheckInclusion(t *testing.T) {
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ab", "eidbaooo"))
	fmt.Println(CheckInclusion("ba", "eidboaoo"))
	fmt.Println(CheckInclusion1("ba", "eidboaoo"))
}

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过 104。



示例 1：

输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：

输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestCharacterReplacement(t *testing.T) {
	fmt.Println(CharacterReplacement1("AABCABBB", 2))
	fmt.Println(CharacterReplacement1("ABAB", 2))
	fmt.Println(CharacterReplacement1("AABABBA", 1))
}

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。


示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

示例 3：
输入：nums = [1,-1], k = 1
输出：[1,-1]

示例 4：
输入：nums = [9,11], k = 2
输出：[11]

示例 5：
输入：nums = [4,-2], k = 2
输出：[4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSliding(t *testing.T) {
	fmt.Println(MaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(MyMaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(MyMaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(SlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func MyMaxSlidingWindow(nums []int, k int) (res []int) {
	queue := []int{}
	for i := 0; i < len(nums); i++ {
		queue = push(queue, nums[i])
		// K个元素之后开始添加最大值到结果集
		if i >= k-1 {
			res = append(res, queue[0])
			// 如果队列不为空，并且队列头部与数组当前元素相等，则弹出队头
			if len(queue) > 0 && queue[0] == nums[i-k+1] {
				queue = queue[1:]
			}
		}
	}
	return
}

func push(arr []int, v int) []int {
	for len(arr) > 0 && arr[len(arr)-1] < v {
		arr = arr[:len(arr)-1]
	}
	return append(arr, v)
}

func MaxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		// 队列不为空，并且当前元素大于等于队尾元素，则弹出最后一个元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 入队
		q = append(q, i)
	}

	// 前K个元素入队
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	// 前K个元素的最大值为队头
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 栈中最大元素小于等于元素，弹出队头元素
		for q[0] <= i-k {
			q = q[1:]
		}
		// 添加最大元素到结果集
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func SlidingWindow(nums []int, k int) (res []int) {
	queue := make([]int, 0, k)
	push := func(v int) {
		for len(queue) > 0 && v > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
	}

	pop := func(v int) {
		if len(queue) > 0 && queue[0] == v {
			queue = queue[1:]
		}
	}
	for i := 0; i < len(nums); i++ {
		push(nums[i])
		if i >= k-1 {
			res = append(res, queue[0])
			pop(nums[i-k+1])
		}
	}
	return
}

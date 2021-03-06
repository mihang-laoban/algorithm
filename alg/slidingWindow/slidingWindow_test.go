package slidingWindow

import (
	. "dp/tools"
	"fmt"
	"math"
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
	fmt.Println(LengthOfLongestSubstring("abcabcbb"))
}

func LengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	set := map[byte]bool{}
	right, ans, size := 0, 0, len(s)
	for left := 0; left < size; left++ {
		if left != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(set, s[left-1])
		}
		// 记录右指针走过的节点，直到当前字母已经出现过或者到达最后一个字母
		for right < size && !set[s[right]] {
			set[s[right]] = true
			right++
		}
		// 第 left 到 right 个字符是一个极长的无重复字符子串
		ans = Max(ans, right-left)
	}
	return ans
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

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	need := make([]int, 256)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	l, r, count, start, winSize := 0, 0, len(t), 0, math.MaxInt32
	for r < len(s) {
		c := s[r]
		if need[c] > 0 {
			count--
		}
		need[c]--
		if count == 0 {
			for l < r && need[s[l]] < 0 {
				need[s[l]]++
				l++
			}
			if r-l+1 < winSize {
				winSize = r - l + 1
				start = l
			}
			need[s[l]]++
			l++
			count++
		}
		r++
	}
	if winSize == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+winSize]
	}
}

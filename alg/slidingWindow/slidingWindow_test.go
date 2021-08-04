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

/*给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

示例1:
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
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

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
	fmt.Println(MinWindowOfficial2("ADOBECODEBANC", "ABC"))
	fmt.Println(MinWindow("ADOBECODEBANC", "ABC"))
}

func MinWindowOfficial2(s string, t string) string {
	var (
		sSize, tSize = len(s), len(t)
		l, r, start  = 0, 0, 0
		count        = tSize
		winSize      = math.MaxInt32
		need         = make([]int, 128)
	)
	if sSize == 0 || tSize == 0 {
		return ""
	}
	for i := 0; i < count; i++ {
		need[t[i]]++
	}
	// 右边界到达S字符串的末尾为止
	for r < sSize {
		// 如果S字符串中的字符在T字符串中存在，则计数器标记找到了一个字符
		if need[s[r]] > 0 {
			count--
		}
		// 标记S字符串中存在的字符，负数为不需要的字符
		need[s[r]]--
		// 如果所有的字符全部找到
		if count == 0 {
			// 右指针与左指针之间还存在空隙，并且当前左指针指向的字符串已经被标记为多余字符，则移动左边界
			for l < r && need[s[l]] < 0 {
				need[s[l]]++
				l++
			}
			// 如果当前窗口小于现有记录则更新窗口大小和左边界
			var curSize = r - l + 1
			if curSize < winSize {
				winSize = curSize
				start = l
			}
			// 移动左边界，需要重新寻找丢失的字符
			need[s[l]]++
			l++
			count++
		}
		// 每次都要移动右边界
		r++
	}
	if winSize == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+winSize]
	}
}

/*给定两个字符串s1和s2，写一个函数来判断 s2 是否包含 s1的排列。
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
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。

注意：字符串长度 和 k 不会超过104。



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
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的k个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。


示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7      3
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

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(MaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(MyMaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func MyMaxSlidingWindow(nums []int, k int) (res []int) {
	return
}

/*30. 串联所有单词的子串
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。

示例 1：
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。

示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]

示例 3：
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]

提示：
1 <= s.length <= 104
s 由小写英文字母组成
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] 由小写英文字母组成
*/

func TestFindSubstring(t *testing.T) {
	fmt.Println(FindSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(FindSubstring("barfoothefoobarman", []string{"bar", "foo"}))
}

func FindSubstring(s string, words []string) (res []int) {
	return
}

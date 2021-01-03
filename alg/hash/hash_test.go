package hash

import (
	"fmt"
	"testing"
)

/*给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(IsAnagram2("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	text1 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		text1[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := text1[t[i]]; ok {
			text1[t[i]]--
			if text1[t[i]] == 0 {
				delete(text1, t[i])
			}
		} else {
			return false
		}
	}

	return len(text1) == 0
}

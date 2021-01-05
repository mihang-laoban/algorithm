package stack

import (
	"fmt"
	"testing"
)

/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func Test(t *testing.T) {

	fmt.Println(valid("){"))
	fmt.Println(IsValid("([()()])"))
}

func valid(s string) interface{} {
	// 奇数个，匹配失败
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	// 下一个括号与前一个不匹配，匹配失败
	for i := 0; i < len(s); i++ {
		x := s[i]
		if pairs[x] > 0 {
			if len(stack) == 0 && stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	//  匹配完，成功
	return len(stack) == 0
}

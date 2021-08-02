package stack

import (
	"fmt"
	"testing"
)

/*给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例2:

输入: "()[]{}"
输出: true
示例3:

输入: "(]"
输出: false
示例4:

输入: "([)]"
输出: false
示例5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestIsValid(t *testing.T) {
	fmt.Println(valid("){"))
	fmt.Println(IsValid("([()()])"))
}

func valid(s string) interface{} {
	return nil
}

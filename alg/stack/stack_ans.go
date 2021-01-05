package stack

func IsValid(s string) bool {
	n := len(s)
	// 不是成对出现肯定不匹配
	if n%2 == 1 {
		return false
	}
	// 记录括号对
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 如果当前元素存在括号对中
			// 如果栈已经为空，或者栈中最后一个元素与当前元素不匹配，则匹配失败
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else { // 如果当前元素不在括号对中，则压栈
			stack = append(stack, s[i])
		}
	}
	// 栈空退出
	return len(stack) == 0
}

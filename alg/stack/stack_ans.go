package stack

func IsValid(s string) bool {
	n := len(s)
	// 不是成对出现肯定不匹配
	if n&1 == 1 {
		return false
	}
	// 记录括号对
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 正括号容器
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 反括号进这里
			// 如果进反括号的时候正括号为空，或正反括号不匹配，匹配失败
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 匹配成功，弹出匹配成功的正括号
			stack = stack[:len(stack)-1]
		} else { // 正括号进这里
			stack = append(stack, s[i])
		}
	}
	// 栈空退出
	return len(stack) == 0
}

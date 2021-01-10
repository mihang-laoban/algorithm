package backtrace

func Subset(nums []int) interface{} {
	trace, res := []int{}, [][]int{}

	var get func([]int, int, []int)
	get = func(nums []int, start int, trace []int) {
		// 创建路径副本
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			get(nums, i+1, trace)
			// 撤回上一层
			trace = trace[:len(trace)-1]
		}
	}

	get(nums, 0, trace)
	return res
}

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func PhoneNum(str string) interface{} {
	// 创建数字键盘映射表
	phoneMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res := []string{}
	var phoneBacktrace func(string, int, string)
	phoneBacktrace = func(str string, index int, trace string) {
		// 处理空字符串
		if len(str) == 0 {
			return
		}
		// 终止条件：数字达到两个
		if index == len(str) {
			res = append(res, trace)
		} else {
			// 找到数字对应的字母
			curStr := phoneMap[string(str[index])]
			// 遍历每一个当前数字代表的字母
			for i := 0; i < len(curStr); i++ {
				phoneBacktrace(str, index+1, trace+string(curStr[i]))
			}
		}
	}

	phoneBacktrace(str, 0, "")
	return res
}

//https://leetcode-cn.com/problems/n-queens/

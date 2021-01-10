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

func SolveNQueens(queueNum int) [][]string {
	solutions := [][]string{}
	queens := make([]int, queueNum)
	// 初始化皇后位置
	for i := 0; i < queueNum; i++ {
		queens[i] = -1
	}
	columns, main, sub := map[int]bool{}, map[int]bool{}, map[int]bool{}

	var backtrack func(int, int)
	backtrack = func(queueNum, row int) {
		if row == queueNum {
			board := generateBoard(queens, queueNum)
			solutions = append(solutions, board)
			return
		}
		for i := 0; i < queueNum; i++ {
			// 如果当前列已被占用，则继续
			if columns[i] {
				continue
			}
			// 如果
			mainIndex := row - i
			if main[mainIndex] {
				continue
			}
			subIndex := row + i
			if sub[subIndex] {
				continue
			}

			// 如果可用，标记占用，继续看下一行
			queens[row] = i
			columns[i], main[mainIndex], sub[subIndex] = true, true, true
			backtrack(queueNum, row+1)

			// 恢复状态
			queens[row] = -1
			delete(columns, i)
			delete(main, mainIndex)
			delete(sub, subIndex)
		}
	}

	backtrack(queueNum, 0)
	return solutions
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

/*

var solutions [][]string
func SolveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}

		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)

		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}
*/

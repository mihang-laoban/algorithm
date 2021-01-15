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
			// 如果当前列已被占用，则看下一个位置
			if columns[i] {
				continue
			}
			// 如果主对角线被占用，则看下一个位置
			mainIndex := row - i
			if main[mainIndex] {
				continue
			}
			// 如果副对角线被占用，则看下一个位置
			subIndex := row + i
			if sub[subIndex] {
				continue
			}

			// 记录当前皇后的行下表
			queens[row] = i
			// 当前位置皇后不存在冲突，占用当前位置，继续看下一行
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

// 生成棋盘
func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		// 每一行都初始化为....
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		// 设置皇后位置
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

func WordSearch(board [][]byte, word string) bool {
	// 节点待检查的四个方向
	directions := [][]int{
		[]int{0, -1},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{1, 0},
	}
	m, n := len(board), len(board[0])
	// 标记走过的格子
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int, int) bool
	var isValid func(int, int) bool

	// 确保数组不越界
	isValid = func(x int, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	dfs = func(i, j, start int) bool {
		// 如果单词查完，校验最后一个字母是否匹配
		if start == len(word)-1 {
			return board[i][j] == word[start]
		}

		if board[i][j] == word[start] {
			// 当前字母匹配，则标记为已访问
			visited[i][j] = true
			// 查看当前字母四个方向是否可以继续移动
			for _, direction := range directions {
				newX := i + direction[0]
				newY := j + direction[1]
				// 如果没越界，没有访问过，则继续看下一个字母
				if isValid(newX, newY) && !visited[newX][newY] && dfs(newX, newY, start+1) {
					return true
				}
			}
			// 如果没有完全匹配，则回溯时还原已访问的位置
			visited[i][j] = false
		}
		// 如果当前字母不匹配，则直接看矩阵里面下一元素
		return false
	}

	// 遍历矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func FindWords(board [][]byte, words []string) []string {
	type TrieNode struct {
		next [26]*TrieNode
		word string
	}
	var build func([]string) *TrieNode

	build = func(words []string) *TrieNode {
		root := &TrieNode{}
		for _, word := range words {
			// 每次从字典树初始位置开始遍历
			cur := root
			for _, w := range word {
				// 如果下一个字母不存在，则创建并移动到下一个节点
				if cur.next[w-'a'] == nil {
					cur.next[w-'a'] = &TrieNode{}
				}
				cur = cur.next[w-'a']
			}
			cur.word = word
		}
		return root
	}

	res := []string{}
	// 创建字典树
	root := build(words)
	var dfs func(int, int, *TrieNode)

	dfs = func(i, j int, root *TrieNode) {
		c := board[i][j]
		// 如果访问过或者当前位置没有字母退出
		if c == '#' || root.next[c-'a'] == nil {
			return
		}
		// 如果有字母或者未访问，则看下一个字母
		root = root.next[c-'a']
		// 如果当前位置存有单词，则记录结果
		if root.word != "" {
			res = append(res, root.word)
			root.word = ""
		}

		// 标记访问
		board[i][j] = '#'
		// 检查数组是否越界
		if i > 0 {
			dfs(i-1, j, root)
		}
		if j > 0 {
			dfs(i, j-1, root)
		}
		if i < len(board)-1 {
			dfs(i+1, j, root)
		}
		if j < len(board[i])-1 {
			dfs(i, j+1, root)
		}
		// 恢复未访问状态
		board[i][j] = c
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, root)
		}
	}
	return res
}

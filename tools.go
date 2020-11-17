package dp

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b{
		return a
	}
	return b
}

func InitMemo(col int, row int) (dp [][]int) {
	dp = make([][]int, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
	}
	return
}
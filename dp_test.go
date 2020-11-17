package dp

import (
	"fmt"
	"testing"
)

func TestDp(t *testing.T) {
	str := "aaa"
	largestSubArr(str)
}

func largestSubArr(str string)  {
	size := len(str)
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}
	res := 0
	for i := 0; i < size; i++ {
		dp[i][i] = true
		res++
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			dp[j][i] = str[i] == str[j] && (i - j <= 3 || dp[j+1][i-1])
			if dp[j][i] {
				res++
			}
		}
	}

	fmt.Println(res)
}

func TestFindLargestArrSum(t *testing.T){
    arr := []int{-2, 1, -3, 4, -1, 3, -5, 1, 2}

    cur := 0
    res := cur
	for i := 0; i < len(arr); i++ {
		tmp := Max(arr[i], arr[i] + cur)
		res = Max(tmp, res)
		cur = tmp
	}
    fmt.Println(res)
}

func TestFindLargestSubSeq(t *testing.T){
	str := "asssasms"

	FindLargestSubSeq(str)
}

func FindLargestSubSeq(str string) {
	size := len(str)
	dp := InitMemo(size, size)
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for i := size-1; i >= 0; i-- {
		for j := i+1; j < size; j++ {
			if str[i] == str[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			}else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[0][size-1])
}

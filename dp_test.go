package dp

import (
	"fmt"
	"testing"
)

func TestDp(t *testing.T) {
	str := "aaa"
	largestSubArr(str)
}

func largestSubArr(str string) {
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
			dp[j][i] = str[i] == str[j] && (i-j <= 3 || dp[j+1][i-1])
			if dp[j][i] {
				res++
			}
		}
	}

	fmt.Println(res)
}

func TestFindLargestArrSum(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 3, -5, 1, 2}

	cur := 0
	res := cur
	for i := 0; i < len(arr); i++ {
		tmp := Max(arr[i], arr[i]+cur)
		res = Max(tmp, res)
		cur = tmp
	}
	fmt.Println(res)
}

func TestFindLargestSubSeq(t *testing.T) {
	str := "asssasms"

	FindLargestSubSeq(str)
}

func FindLargestSubSeq(str string) {
	size := len(str)
	dp := InitMemo(size, size)
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if str[i] == str[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[0][size-1])
}

func TestFindLongestSharedSubSeq(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"
	size1 := len(text1)
	size2 := len(text2)

	dp := InitMemo(size1+1, size2+1)
	for j := 1; j <= size2; j++ {
		for i := 1; i <= size1; i++ {
			str1 := string(text1[i-1])
			str2 := string(text2[j-1])
			if str1 == str2 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[size1][size2])
}

func TestFindLongestSharedSubSeq2(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"
	size1 := len(text1)
	size2 := len(text2)

	dp := InitMemo(size1+1, size2+1)
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[size1][size2])
}

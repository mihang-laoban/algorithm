package test

import (
	"dp/tools"
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
		tmp := tools.Max(arr[i], arr[i]+cur)
		res = tools.Max(tmp, res)
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
	dp := tools.InitMemo(size, size)
	for i := 0; i < size; i++ {
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if str[i] == str[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = tools.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[0][size-1])
}

func TestFindLongestSharedSubSeq(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"

	s1 := len(text1)
	s2 := len(text2)
	dp := make([][]int, s1+1)
	for i := 0; i < s1+1; i++ {
		dp[i] = make([]int, s2+1)
	}
	for i := 1; i < s1+1; i++ {
		for j := 1; j < s2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = tools.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[s1][s2])
}

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ade"
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, 2)
	}

	for j := 1; j <= n; j++ {
		a := j % 2
		b := (j - 1) % 2
		for i := 1; i < m; i++ {
			if text2[j-1] == text1[i-1] {
				dp[i][a] = dp[i-1][b] + 1
			} else {
				dp[i][a] = tools.Max(dp[i-1][a], dp[i][b])
			}
		}
	}
	fmt.Println(dp[m][n%2])
	//
	//int[][] dp = new int[m + 1][2];
	//for (int[] row: dp) { Arrays.fill(row, 0); }
	//
	//for (int j = 1; j <= n; j++) {
	//	int a = j % 2;
	//	int b = (j - 1) % 2;
	//	for (int i = 1; i <= m; i++) {
	//		if (text2.charAt(j - 1) == text1.charAt(i - 1)) {
	//			dp[i][a] = dp[i - 1][b] + 1;
	//		} else {
	//			dp[i][a] = Math.max(dp[i - 1][a], dp[i][b]);
	//		}
	//	}
	//}
	//
	//return dp[m][n%2];
}

func TestRobotPathSum(t *testing.T) {

}

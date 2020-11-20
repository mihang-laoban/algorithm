package DP

import (
	"dp/tools"
	"fmt"
)

func FindLongestSharedSubSeq() {
	str1 := "abcde"
	str2 := "ade"

	size1 := len(str1)
	size2 := len(str2)

	dp := tools.InitMemo(size1+1, size2+1)
	for i := 1; i < size1+1; i++ {
		for j := 1; j < size2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = tools.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[size1][size2])
}

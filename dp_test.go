package dp

import (
	"fmt"
	"testing"
)

func TestDp(t *testing.T) {
	str := "aaa"
	find2(str)
}

func find2(s string) {
	if len(s) == 0 {
		fmt.Println(0)
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := 0
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		res++
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[j] == s[i] && (i-j < 3 || dp[i+1][j-1])
			if dp[i][j] {
				res++
			}
		}
	}
	fmt.Println(res)
}

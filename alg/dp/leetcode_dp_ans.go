package dp

import . "dp/tools"

func LongestPalindrome1(s string) string {
	start, maxLen, n := 0, 1, len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[j][i] = s[i] == s[j] && (i-j < 3 || dp[j+1][i-1])
			if dp[j][i] && maxLen < i-j+1 {
				maxLen = i - j + 1
				start = j
			}
		}
	}
	return s[start : start+maxLen]
}

func LongestPalindrome2(s string) (res string) {
	index, maxLen, size := 0, 0, len(s)
	for index < size {
		l, r := index, index
		for l >= 0 && s[l] == s[index] {
			l--
		}
		for r < size && s[r] == s[index] {
			r++
		}
		next := r
		for l >= 0 && r < size && s[r] == s[l] {
			l--
			r++
		}
		cur := r - l - 1
		if cur > maxLen {
			res = s[l+1 : r]
			maxLen = cur
		}
		index = next
	}
	return
}

func MaxSubArray(nums []int) int {
	pre, res := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		cur := Max(nums[i], pre+nums[i])
		res = Max(cur, res)
		pre = cur
	}
	return res
}

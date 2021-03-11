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

func Trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans, size := 0, len(height)
	lm, rm := make([]int, size), make([]int, size)
	lm[0] = height[0]
	for i := 1; i < size; i++ {
		lm[i] = Max(height[i], lm[i-1])
	}
	rm[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rm[i] = Max(height[i], rm[i+1])
	}
	for i := 1; i < size-1; i++ {
		ans += Min(lm[i], rm[i]) - height[i]
	}
	return ans
}

func Trap2(height []int) int {
	l, r, lm, rm, res := 0, len(height)-1, 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if lm < height[l] {
				lm = height[l]
			} else {
				res += lm - height[l]
			}
			l++
		} else {
			if rm < height[r] {
				rm = height[r]
			} else {
				res += rm - height[r]
			}
			r--
		}
	}
	return res
}

func MinDistance(word1 string, word2 string) int {
	s1, s2 := len(word1), len(word2)
	dp := make([][]int, s1+1)
	for i := 0; i < s1+1; i++ {
		dp[i] = make([]int, s2+1)
	}
	for i := 0; i < s1+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < s2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < s1+1; i++ {
		for j := 1; j < s2+1; j++ {
			dp[i][j] = Min(dp[i-1][j-1], Min(dp[i-1][j], dp[i][j-1])) + 1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[s1][s2]
}

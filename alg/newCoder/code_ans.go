package newCoder

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r-l)>>1 + l
		if nums[m] == target {
			for m >= 0 {
				if nums[m] != target {
					break
				}
				m--
			}
			return m + 1
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}

	}
	return -1
}

func ClimbStairsDP(n int) int {
	dp := make([]int, n+1)
	return ClimbStairsDP_(n, dp)
}

func ClimbStairsDP_(n int, dp []int) int {
	if dp[n] > 0 {
		return dp[n]
	}
	if n == 1 {
		dp[1] = 1
	} else if n == 2 {
		dp[2] = 2
	} else {
		dp[n] = ClimbStairsDP_(n-1, dp) + ClimbStairsDP_(n-2, dp)
	}
	return dp[n]
}

func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

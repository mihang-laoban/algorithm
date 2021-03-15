package newCoder

import (
	"dp/ds/linkedList"
	"dp/tools"
)

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

func AddTwoNumbers(head1 *linkedList.ListNode, head2 *linkedList.ListNode) *linkedList.ListNode {
	nums1, nums2 := []int{}, []int{}
	for head1 != nil {
		nums1 = append(nums1, head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		nums2 = append(nums2, head2.Val)
		head2 = head2.Next
	}
	var add int
	var target *linkedList.ListNode
	for len(nums1) > 0 || len(nums2) > 0 || add > 0 {
		var a, b int
		if len(nums1) > 0 {
			a = nums1[len(nums1)-1]
			nums1 = nums1[:len(nums1)-1]
		}
		if len(nums2) > 0 {
			b = nums2[len(nums2)-1]
			nums2 = nums2[:len(nums2)-1]
		}
		res := a + b + add
		add = res / 10
		rest := res % 10
		target = &linkedList.ListNode{Val: rest, Next: target}
	}
	return target
}

func MinPathSum(matrix [][]int) int {
	col, row := len(matrix), len(matrix[0])
	dp := make([][]int, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < col; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < row; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			dp[i][j] = tools.Min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[col-1][row-1]
}

func OddEvenList(head *linkedList.ListNode) *linkedList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	odd, even := head, head.Next
	for second != nil && second.Next != nil {
		first.Next = first.Next.Next
		first = first.Next
		second.Next = second.Next.Next
		second = second.Next
	}
	first.Next = even
	return odd
}

func GetLongestPalindrome(str string, size int) int {
	res := 1
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}
	for r := 1; r < size; r++ {
		for l := r - 1; l >= 0; l-- {
			if r-l < 2 {
				dp[l][r] = str[r] == str[l]
				res = tools.Max(res, r-l+1)
				continue
			}
			dp[l][r] = dp[l+1][r-1] && str[r] == str[l]
			if dp[l][r] {
				res = tools.Max(res, r-l+1)
			}
		}
	}
	return res
}

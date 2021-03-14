package newCoder

import "dp/ds/linkedList"

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

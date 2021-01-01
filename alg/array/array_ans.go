package array

import (
	. "dp/tools"
	"sort"
)

func threesome(nums []int) (res [][]int) {
	n := len(nums)
	sort.Ints(nums)
	third := n - 1
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[third]+nums[second] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[third]+nums[second] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}

func MoveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func LargeCon(nums []int) (ma int) {
	j := len(nums) - 1
	for i := 0; i < j; {
		short := 0
		if nums[i] < nums[j] {
			short = nums[i]
			i++
		} else {
			short = nums[j]
			j--
		}
		ma = Max((j-i+1)*short, ma)
	}
	return
}

func Climb(n int) int {
	i, j := 0, 1
	for k := 0; k < n+1; k++ {
		i, j = i+j, i
	}
	return i
}

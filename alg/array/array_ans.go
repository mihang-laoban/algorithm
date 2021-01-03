package array

import (
	. "dp/ds/linkedList"
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

func LargestRectangleArea(heights []int) (res int) {
	// extend the length of the array to add sentinel
	size := len(heights) + 2
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := []int{0} // create a stack with 0 as the base
	for i := 1; i < size; i++ {
		// try to find the index whose height is lower than the one at the stack tail
		//- the element is the right boundary of the rectangle
		for heights[i] < heights[stack[len(stack)-1]] {
			curHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			curWidth := i - stack[len(stack)-1] - 1
			res = Max(res, curHeight*curWidth)
		}
		// if the current height is higher that of the stack tail,
		// add the index to the stack to continue finding the right boundary
		stack = append(stack, i)
	}
	return
}

func MyLargestRectangleArea(heights []int) (res int) {
	heights = PreInt(0, append(heights, 0))
	size, stack := len(heights), Deque{}
	stack.Append(0)
	for i := 1; i < size; i++ {
		for heights[i] < heights[stack.PeekLast().(int)] {
			curHeight := heights[stack.Pop().(int)]
			curWidth := i - stack.PeekLast().(int) - 1
			res = Max(res, curWidth*curHeight)
		}
		stack.Append(i)
	}
	return res
}

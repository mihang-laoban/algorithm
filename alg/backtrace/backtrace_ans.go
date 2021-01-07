package backtrace

func Subset(nums []int) interface{} {
	res := [][]int{}
	stack := []int{}

	var back func([]int, int, []int)
	back = func(nums []int, start int, stack []int) {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			stack = append(stack, nums[i])
			back(nums, i+1, stack)
			stack = stack[:len(stack)-1]
		}
	}

	back(nums, 0, stack)
	return res
}

package backtrace

func Subset(nums []int) interface{} {
	trace, res := []int{}, [][]int{}

	var get func([]int, int, []int)
	get = func(nums []int, start int, trace []int) {
		tmp := make([]int, len(trace))
		copy(tmp, trace)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			get(nums, i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}

	get(nums, 0, trace)
	return res
}

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
//https://leetcode-cn.com/problems/n-queens/

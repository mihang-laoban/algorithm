package sorting

import . "dp/tools"

func Bubble(nums []int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func Select(nums []int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

func Insert1(nums []int) []int {
	size := len(nums)
	for i := 1; i < size; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}

func Shell(nums []int) []int {
	size := len(nums)
	for gap := size >> 1; gap > 0; gap >>= 1 {
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < size; i++ {
			j := i
			for j-gap >= 0 && nums[j] < nums[j-gap] {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
				j -= gap
			}
		}
	}
	return nums
}

func Quick1(nums []int) []int {
	return QuickSort(nums, 0, len(nums)-1)
}

func QuickSort(nums []int, start, end int) []int {
	if start > end {
		return nums
	}
	low, high, mid := start, end, nums[(start+end)>>1]
	for low <= high {
		for nums[low] < mid {
			low++
		}
		for nums[high] > mid {
			high--
		}
		if low <= high {
			nums[low], nums[high] = nums[high], nums[low]
			low++
			high--
		}
	}
	if start < high {
		QuickSort(nums, start, high)
	}
	if end > low {
		QuickSort(nums, low, end)
	}
	return nums
}

func Merge1(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	mid := length >> 1
	return subMerge(Merge1(nums[:mid]), Merge1(nums[mid:]))
}

func subMerge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func Heap1(nums []int) []int {
	size := len(nums)
	// 创建大根堆
	for i := size >> 1; i >= 0; i-- {
		heapifyL(nums, i, size)
	}
	// 从最后一个元素开始排序，把最后一个节点拿到堆顶
	for i := size - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		size--
		heapifyR(nums, 0, size)
	}
	return nums
}

func heapifyR(nums []int, index, size int) {
	// 当前元素的左右孩子的位置，并假定当前节点就是最大值的位置
	left, right, largest := index<<1+1, index<<1+2, index
	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}
	// 如果更新过最大节点后和原有最大节点不一致，则交换最大节点和当前节点的位置，并继续寻找
	if largest != index {
		nums[index], nums[largest] = nums[largest], nums[index]
		heapifyR(nums, largest, size)
	}
}

func heapifyL(lists []int, index, size int) {
	cur := lists[index]
	for i := index<<1 + 1; i < size; i = i<<1 + 1 {
		// 找到最小的那个节点下标
		if i+1 < size && lists[i] < lists[i+1] {
			i++
		}
		if lists[index] < lists[i] {
			// 如果当前元素大于最小元素的值，则当前元素更新为最小元素
			// 标记被更新的元素下标，已备后续更新
			lists[index], index = lists[i], i
		}
	}
	if cur != lists[index] {
		// 如果迭代过后下标被更新了，则发生位置交换，否则无交换
		lists[index] = cur
		heapifyL(lists, index, size)
	}
}

func Counting(nums []int) []int {
	bucketLen := FindX(nums, Max) + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		bucket[nums[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			nums[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return nums
}

func Bucket(nums []int) []int {
	tmp := make([]int, 100)
	for _, v := range nums {
		tmp[v] += 1
	}
	var result []int
	for k, v := range tmp {
		if v > 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}

func Radix(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	max := nums[0]
	numsLen := len(nums)
	for i := 1; i < numsLen; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// 计算最大值的位数
	maxDigit := 0
	for max > 0 {
		max = max / 10
		maxDigit++
	}
	// 定义每一轮的除数，1,10,100...
	divisor := 1
	// 定义了10个桶，为了防止每一位都一样所以将每个桶的长度设为最大,与原数组大小相同
	bucket := [10][20]int{{0}}
	// 统计每个桶中实际存放的元素个数
	count := [10]int{0}
	// 获取元素中对应位上的数字，即装入那个桶
	var digit int
	// 经过maxDigit+1次装通操作，排序完成
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < numsLen; j++ {
			tmp := nums[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		// 被排序数组的下标
		k := 0
		// 从0到9号桶按照顺序取出
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				nums[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}
	return nums
}

func GetLeastNumbers(arr []int, k int) []int {
	QuickSelect(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func QuickSelect(nums []int, start, end, k int) {
	for start < end {
		//findMidInTree(nums, start, end)
		pivot := nums[start]
		l, r := start, end
		for l <= r {
			for l <= r && nums[l] < pivot {
				l++
			}
			for l <= r && nums[r] > pivot {
				r--
			}
			if l <= r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}
		if l < k {
			start = l
		} else {
			end = l - 1
		}
	}
}

func findMidInTree(nums []int, start, end int) {
	mid := start + (end-start)>>1
	if nums[start] > nums[end] {
		nums[start], nums[end] = nums[end], nums[start]
	}
	if nums[mid] > nums[end] {
		nums[mid], nums[end] = nums[end], nums[mid]
	}
	if nums[mid] > nums[start] {
		nums[mid], nums[start] = nums[start], nums[mid]
	}
}

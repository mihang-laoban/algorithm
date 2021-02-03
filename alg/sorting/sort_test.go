package sorting

import (
	. "dp/tools"
	"testing"
)

var x = []int{6, 4, 2, 8, 9, 5, 1, 3, 7, 6, 8}
var length = len(x)

type Sorting interface {
	Insert()
	Select()
	Quick()
	Merge()
	Stack()
	Bubble()
	Shell()
	Bucket()
	Counting()
	Radix()
}

func TestBubble(t *testing.T)   { LoopPrint(Bubble(x)) }
func TestSelect(t *testing.T)   { LoopPrint(Select(x)) }
func TestInsert(t *testing.T)   { LoopPrint(Insert(x)) }
func TestShell(t *testing.T)    { LoopPrint(Shell(x)) }
func TestQuick(t *testing.T)    { LoopPrint(Quick(x)) }
func TestMerge(t *testing.T)    { LoopPrint(Merge(x)) }
func TestHeap(t *testing.T)     { LoopPrint(Heap(x)) }
func TestCounting(t *testing.T) { LoopPrint(Counting(x)) }
func TestBucket(t *testing.T)   { LoopPrint(Bucket(x)) }
func TestRadix(t *testing.T)    { LoopPrint(Radix(x)) }

func Bubble(arr []int) []int {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func Select(arr []int) []int {
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func Insert(arr []int) []int {
	for i := 1; i < length; i++ { //把第一个元素作为有序序列
		j := i
		if arr[i] < arr[i-1] {
			for j > 0 && arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				j--
			}
		}
	}
	return arr
}

func Shell(arr []int) []int {
	for gap := length >> 1; gap > 0; gap >>= 1 {
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < length; i++ {
			j := i
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
	}
	return arr
}

func Quick(arr []int) []int {
	return QuickSort(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		i, j := start, end
		key := arr[(start+end)>>1]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
	return arr
}

func Merge(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length >> 1
	return merge(Merge(arr[0:middle]), Merge(arr[middle:]))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
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

func Heap(nums []int) []int {
	size := len(nums)
	// 创建大根堆
	for i := size >> 1; i >= 0; i-- {
		heapify(nums, i, size)
	}
	// 从最后一个元素开始排序，把最后一个节点拿到堆顶
	for i := size - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		size--
		heapify(nums, 0, size)
	}
	return nums
}

func heapify(nums []int, i, size int) {
	// 当前元素的左右孩子的位置，并假定当前节点就是最大值的位置
	left, right, largest := i<<1+1, i<<1+2, i
	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}
	// 如果更新过最大节点后和原有最大节点不一致，则交换最大节点和当前节点的位置，并继续寻找
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, size)
	}
}

func Counting(arr []int) []int {
	bucketLen := FindX(arr, Max) + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}

func Bucket(arr []int) []int {
	tmp := make([]int, 100)
	for _, v := range arr {
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

func Radix(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	max := arr[0]
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		if arr[i] > max {
			max = arr[i]
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
		for j := 0; j < arrLen; j++ {
			tmp := arr[j]
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
				arr[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}
	return arr
}

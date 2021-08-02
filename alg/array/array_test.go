package array

import (
	. "dp/ds/linkedList"
	. "dp/tools"
	"fmt"
	"sort"
	"testing"
)

/*给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
[-1, 0, 1],
[-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func init() {
	Max(1, 2)
	_ = Deque{}
}

func TestThree(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//res := threesome(nums)
	fmt.Println(three(nums))
	fmt.Println(threesome(nums))
}

func three(nums []int) (res [][]int) {
	size := len(nums)
	sort.Ints(nums)
	for first := 0; first < size; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := size - 1
		target := -nums[first]
		for second := first + 1; second < size; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}

/*给定一个整数数组 nums和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestTwoSum(t *testing.T) {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) (res []int) {
	return
}

/*给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func TestMoveZero(t *testing.T) {
	fmt.Println(MoveZeroes([]int{1, 0, 0, 3, 12}))
	fmt.Println(move0([]int{1, 0, 0, 3, 12}))
}

func move0(nums []int) []int {
	return nums
}

/*给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
示例 2：

输入：height = [1,1]
输出：1
示例 3：

输入：height = [4,3,2,1,4]
输出：16
示例 4：

输入：height = [1,2,1]
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestLargestContainer(t *testing.T) {
	fmt.Println(LargeCon([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/*假设你正在爬楼梯。需要 n阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestClimb(t *testing.T) {
	fmt.Println(Climb(2))
}

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
示例:

输入: [2,1,5,6,2,3]
输出: 10
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/*/

func TestLargestRectangle(t *testing.T) {
	fmt.Println(LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(MyLargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

/*实现pow(x, n)，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例2:

输入: 2.10000, 3
输出: 9.26100
示例3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/powx-n
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPow(t *testing.T) {
	fmt.Println(MyPow(2.00, 10))
	fmt.Println(MyPow1(2.00, 10))
	fmt.Println(MyPowNew(2.00, 10))
}

func MyPow1(f float64, i int) float64 {
	if i < 0 {
		return 1 / q(f, -i)
	}
	return q(f, i)
}

func q(f float64, i int) float64 {
	if i == 0 {
		return 1
	}
	y := q(f, i>>1)
	if i&1 == 0 {
		return y * y
	} else {
		return y * y * f
	}
}

/*
实现int sqrt(int x)函数。

计算并返回x的平方根，其中x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
输入: 4
输出: 2

示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
    由于返回类型是整数，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSqrt(t *testing.T) {
	fmt.Println(MySqrt(8))
}

func TestBinarySearch(t *testing.T) {
	target, nums := 9, []int{1, 2, 3, 4, 5, 6, 8, 9}
	fmt.Println(BinarySearch(target, nums))
}

/*
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为[4,5,6,7,0,1,2] ）。
请你在数组中搜索target ，如果数组中存在这个目标值，则返回它的索引，否则返回-1。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestSearchRotatedArray(t *testing.T) {
	//target, nums := 1, []int{4, 5, 6, 7, 0, 1, 2}
	target, nums := 1, []int{3, 2, 1}
	//target, nums := 2, []int{1, 1, 1, 1, 1, 2, 1, 1, 1} // 5
	fmt.Println(SearchRotatedArray(target, nums))
	fmt.Println(SearchRotatedArray2(target, nums))
}

func SearchRotatedArray2(target int, nums []int) int {
	return -1
}

func TestRemoveDuplicate(t *testing.T) {
	fmt.Println(RemoveDuplicate([]int{0, 0, 1, 2, 3, 3, 2}))
}

func RemoveDuplicate(nums []int) interface{} {
	sort.Ints(nums)
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return nums[:slow+1]
}

/*数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？

注意：本题相对书上原题稍作改动

示例 1：

输入：[3,0,1]
输出：2


示例 2：

输入：[9,6,4,2,3,5,7,0,1]
输出：8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMissingNumber(t *testing.T) {
	fmt.Println(MissingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(MissingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func MissingNumber1(nums []int) int {
	a, i := 0, 0
	for ; i < len(nums); i++ {
		a = a ^ i ^ nums[i]
	}
	return a ^ i
}

func MissingNumber2(nums []int) int {
	n := len(nums) + 1
	sum := (n - 1) * n >> 1
	for k := range nums {
		sum -= nums[k]
	}
	return sum
}

/*给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

以任意顺序返回这两个数字均可。

示例 1:

输入: [1]
输出: [2,3]
示例 2:

输入: [2,3]
输出: [1,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-two-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestMissingTwo(t *testing.T) {
	fmt.Println(MissingTwo1([]int{1, 3, 4, 6}))
	fmt.Println(MissingTwo2([]int{1, 3, 4, 6}))
}

func MissingTwo1(nums []int) []int {
	sum, n := 0, len(nums)+2
	for _, x := range nums {
		sum += x
	}
	sumTwo := n*(n+1)>>1 - sum
	limits := sumTwo >> 1
	sum = 0
	for _, x := range nums {
		if x <= limits {
			sum += x
		}
	}
	one := limits*(limits+1)>>1 - sum
	return []int{one, sumTwo - one}
}

func MissingTwo2(nums []int) []int {
	ans, n := 0, len(nums)
	for i := 1; i <= n+2; i++ {
		ans ^= i
	}
	for _, value := range nums {
		ans ^= value
	}
	one, diff := 0, ans&-ans
	for i := 1; i <= n+2; i++ {
		if diff&i != 0 {
			one ^= i
		}
	}
	for _, value := range nums {
		if diff&value != 0 {
			one ^= value
		}
	}
	return []int{one, one ^ ans}
}

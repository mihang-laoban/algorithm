package array

import (
	. "dp/ds/linkedList"
	. "dp/tools"
	"fmt"
	"sort"
	"testing"
)

/*给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

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
			for second != third && nums[second]+nums[third] > target {
				third--
			}
			//if second == third {
			//	continue
			//}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}

/*给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	return nil
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
}

/*给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
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

/*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

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

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。


示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

示例 3：
输入：nums = [1,-1], k = 1
输出：[1,-1]

示例 4：
输入：nums = [9,11], k = 2
输出：[11]

示例 5：
输入：nums = [4,-2], k = 2
输出：[4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSliding(t *testing.T) {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(MyMaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func push(arr []int, v int) []int {
	for len(arr) > 0 && arr[len(arr)-1] < v {
		arr = arr[:len(arr)-1]
	}
	return AppInt(arr, v)
}

func MyMaxSlidingWindow(nums []int, k int) (res []int) {
	queue := []int{}
	for i := 0; i < len(nums); i++ {
		// 前K个元素加入队列
		if i < k-1 {
			queue = push(queue, nums[i])
		} else {
			queue = push(queue, nums[i])
			res = AppInt(res, queue[0])
			// 如果队列不为空，队列头部与数组
			if len(queue) > 0 && queue[0] == nums[i-k+1] {
				queue = queue[:len(queue)-1]
			}
		}
	}
	return
}

func MaxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		// 队列不为空，并且当前元素大于等于队尾元素，则弹出最后一个元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 入队
		q = append(q, i)
	}

	// 前K个元素入队
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	// 前K个元素的最大值为队头
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		// 栈中最大元素小于等于元素，弹出队头元素
		for q[0] <= i-k {
			q = q[1:]
		}
		// 添加最大元素到结果集
		ans = append(ans, nums[q[0]])
	}
	return ans
}

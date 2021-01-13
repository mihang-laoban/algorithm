package array

import (
	. "dp/ds"
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
	fmt.Println(MaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(MyMaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(MyMaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(SlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func SlidingWindow(nums []int, k int) (res []int) {
	queue := make([]int, 0, k)
	push := func(v int) {
		for len(queue) > 0 && v > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
	}

	pop := func(v int) {
		if len(queue) > 0 && queue[0] == v {
			queue = queue[1:]
		}
	}
	for i := 0; i < len(nums); i++ {
		push(nums[i])
		if i >= k-1 {
			res = append(res, queue[0])
			pop(nums[i-k+1])
		}
	}
	return
}

func push(arr []int, v int) []int {
	for len(arr) > 0 && arr[len(arr)-1] < v {
		arr = arr[:len(arr)-1]
	}
	return append(arr, v)
}

func MyMaxSlidingWindow(nums []int, k int) (res []int) {
	queue := []int{}
	for i := 0; i < len(nums); i++ {
		queue = push(queue, nums[i])
		// K个元素之后开始添加最大值到结果集
		if i >= k-1 {
			res = append(res, queue[0])
			// 如果队列不为空，并且队列头部与数组当前元素相等，则弹出队头
			if len(queue) > 0 && queue[0] == nums[i-k+1] {
				queue = queue[1:]
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

/*实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

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

func MyPow1(f float64, i int) interface{} {
	return nil
}

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

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
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。


示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
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
	target, nums := 2, []int{1, 1, 1, 1, 1, 2, 1, 1, 1} // 5
	fmt.Println(SearchRotatedArray(target, nums))
	fmt.Println(SearchRotatedArray2(target, nums))
}

func SearchRotatedArray2(target int, nums []int) interface{} {
	return -1
}

/*给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
["1","1","0","0","0"],
["1","1","0","0","0"],
["0","0","1","0","0"],
["0","0","0","1","1"]
]
输出：3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestIsland(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// 空地的数量
	spaces := 0
	unionFind := UnionFind{}
	unionFind.UnionFind(rows * cols)
	directions := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				spaces++
			} else {
				// 此时 grid[i][j] == '1'
				for _, direction := range directions {
					newX := i + direction[0]
					newY := j + direction[1]
					// 先判断坐标合法，再检查右边一格和下边一格是否是陆地
					if newX < rows && newY < cols && grid[newX][newY] == '1' {
						unionFind.Union(i*cols+j, newX*cols+newY)
					}
				}
			}
		}
	}

	return unionFind.Count - spaces
}

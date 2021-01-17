package bit

import (
	"fmt"
	"testing"
)

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a >> 1
	a = a << 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}

func TestTest(t *testing.T) {
	fmt.Printf("even: %v\n", even(100))
	a, b := swap(100, 200)
	fmt.Printf("swap: %d\t%d\n", a, b)
	fmt.Printf("shifting: %d\n", shifting(100))
	fmt.Printf("nagation: %d\n", nagation(100))
	fmt.Println(13 & 1)
	fmt.Println(12 & 0)
	fmt.Println(12 & (12 - 1))
	fmt.Println(12 & -12)
	fmt.Println(12 ^ 0)
	fmt.Println(12 ^ -1)
	fmt.Println(11 & 0)
	fmt.Println(11 & 1)

}

/*输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestBit(t *testing.T) {
	n := 0101011

	fmt.Println(Bits1(uint32(n)))
	fmt.Println(Bits2(uint32(n)))
}

// best
func Bits1(n uint32) int {
	sum := 0
	for n > 0 {
		sum++
		// 清零最低位的1
		n &= (n - 1)
	}
	return sum
}

func Bits2(n uint32) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

/*给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:
输入: 1
输出: true
解释: 20 = 1

示例 2:
输入: 16
输出: true
解释: 24 = 16

示例 3:
输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/********************************************
* ^(XOR) 在go语言中XOR是作为二元运算符存在的：
* 但是如果是作为一元运算符出现，他的意思是按位取反，相当于 ~
*********************************************/

func Test2Mi(t *testing.T) {
	fmt.Println(IsPowerOfTwo(16))
	fmt.Println(IsPowerOfTwo2(16))
	fmt.Println(4 & -4)       // 4
	fmt.Println(15 & -15)     // 1
	fmt.Println(^8 + 1)       // -8
	fmt.Println(8 ^ 0)        // 8
	fmt.Println(8 ^ 1)        // 9
	fmt.Println(^8)           // -9
	fmt.Println(8 ^ -1)       // -9
	fmt.Println(8 ^ (8 ^ ^8)) // -9
	fmt.Println(8 ^ 8)        // 0
	fmt.Println(8 ^ ^8)       // -1
	fmt.Println(^0)           // -1
}

func IsPowerOfTwo(n int) bool {
	// n&(n-1) 清零最低位的1
	return n > 0 && n&(n-1) == 0
}

func IsPowerOfTwo2(n int) bool {
	// x & (-x) 可以获取到二进制中最右边的 1，且其它位设置为 0
	// -x = ^x + 1
	return n > 0 && n&-n == n
}

/*颠倒给定的 32 位无符号整数的二进制位。

示例 1：
输入: 00000010100101000001111010011100
输出: 00111001011110000010100101000000
解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。

示例 2：
输入：11111111111111111111111111111101
输出：10111111111111111111111111111111
解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
     因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestReverseBits(t *testing.T) {
	// 956301312 (00111001000000000000000000000000)
	fmt.Println(reverseBits(010011100))
}

func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

/*给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
输入: [2,2,1]
输出: 1

示例 2:
输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSingleNumber(t *testing.T) {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println(SingleNumber(nums1))
	fmt.Println(SingleNumber(nums2))
}

func SingleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

/*给你一个整数数组 nums ，返回该数组所有可能的子集（幂集）。解集不能包含重复的子集。


示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSubSet(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(SubSet(nums))
}

func SubSet(nums []int) [][]int {
	size := len(nums)
	n := 1 << size
	res := [][]int{}
	for i := 0; i < n; i++ {
		cur := []int{}
		for j := 0; j < size; j++ {
			if i>>j&1 == 1 {
				//if i&(1<<j) != 0 {
				//tmp := 1 << j
				//one := i & tmp
				//if one != 0 {
				cur = append(cur, nums[j])
			}
		}
		res = append(res, cur)
	}
	return res
}

/*给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :
输入: [1,2,1,3,2,5]
输出: [3,5]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSingeNumber2(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(SingleNumber2(nums))
}

func SingleNumber2(nums []int) []int {
	bitmask := 0
	for _, num := range nums {
		bitmask ^= num
	}
	// 确定第一个差异位
	diff := bitmask & (-bitmask)

	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x ^= num
		}
	}
	return []int{x, bitmask ^ x}
}

/*给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestSingeNumber3(t *testing.T) {
	nums := []int{0, 1, 99, 0, 1, 0, 1}
	fmt.Println(SingleNumber3(nums))
}

func SingleNumber3(nums []int) interface{} {
	one, two := 0, 0
	for _, n := range nums {
		one = one ^ n & ^two
		two = two ^ n & ^one
	}
	return one
}

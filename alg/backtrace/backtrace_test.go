package backtrace

import (
	"fmt"
	"testing"
)

/*给你一个整数数组 nums ，返回该数组所有可能的子集（幂集）。解集不能包含重复的子集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]*/
func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(Subset(nums))
	fmt.Println(subset(nums))
}

func subset(nums []int) interface{} {
	res := [][]int{}
	trace := []int{}
	var get func([]int, int, []int)
	get = func(nums []int, start int, trace []int) {
		res = append(res, trace)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			get(nums, i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}
	get(nums, 0, trace)
	return res
}

func TestN(t *testing.T) {
	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	slice := make([]int, 5, 5)
	slice1 := slice
	slice = append(slice, 1)
	slice[0] = 1
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice1) //[0 0 0 0 0]
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	slice2 := make([]int, 4, 4)
	slice3 := make([]int, 5, 5)
	fmt.Println(copy(slice2, slice)) //4
	fmt.Println(copy(slice3, slice)) //5
	slice2[1] = 2
	slice3[1] = 3
	fmt.Println(slice)  //[1 0 0 0 0 1]
	fmt.Println(slice2) //[1 2 0 0]
	fmt.Println(slice3) //[1 3 0 0 0]
}

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

phoneMap := map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestPhoneNum(t *testing.T) {
	fmt.Println(PhoneNum("23"))
	fmt.Println(Phone("23"))
}

func Phone(str string) interface{} {
	res := []string{}
	return res
}

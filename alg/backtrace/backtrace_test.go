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

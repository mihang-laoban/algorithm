package sorting

import (
	. "dp/tools"
	"fmt"
	"testing"
)

var x = []int{6, 4, 2, 5, 9, 3, 10, 7, 6, 8, 1}

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

func Insert(nums []int) []int {
	return nums
}

func Quick(nums []int) []int {
	return []int{}
}

func Merge(nums []int) []int {
	return nums
}

func Heap(nums []int) []int {
	return nums
}

/*输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]

示例 2：
输入：arr = [0,1,2,1], k = 1
输出：[0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func TestGetLeastNumbers(t *testing.T) {
	//fmt.Println(GetLeastNumbers([]int{3, 2, 1}, 1))
	fmt.Println(GetLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4))
}

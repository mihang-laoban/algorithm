package sorting

import (
	"testing"
)

var x  = []int {6,4,2,8,9,5,1,3,7,6,8}
var size = len(x)

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

func TestBubble(t *testing.T)   { Bubble(x) }
func TestSelect(t *testing.T)   { Select(x) }
func TestShell(t *testing.T)    { Shell(x) }
func TestInsert(t *testing.T)   { Insert(x) }
func TestQuick(t *testing.T)    { Quick(x) }
func TestMerge(t *testing.T)    { Merge(x) }
func TestHeap(t *testing.T)     { Heap(x) }
func TestBucket(t *testing.T)   { Bucket(x) }
func TestCounting(t *testing.T) { Counting(x) }
func TestRadix(t *testing.T)    { Radix(x) }



func Bubble(nums []int){
	//sort.Sort(nums)
}

func Select(nums []int){

}

func Shell(nums []int){

}

func Insert(nums []int) {

}

func Quick(nums []int){

}

func Merge(nums []int){

}

func Heap(nums []int){

}

func Bucket(nums []int){

}

func Counting(nums []int){

}

func Radix(nums []int){

}

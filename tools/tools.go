package tools

import (
	"fmt"
	. "reflect"
	"strconv"
	"time"
)

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindX(a []int, x func(a, b int) int) (res int) {
	for _, value := range a {
		res = x(res, value)
	}
	return
}

func IsSameType(a, b interface{}) bool {
	return TypeOf(a) == TypeOf(b)
}

func Swap(a, b interface{}) {
	if IsSameType(a, b) {
		switch a.(type) {
		case *int:
			*a.(*int), *b.(*int) = *b.(*int), *a.(*int)
		case *string:
			*a.(*string), *b.(*string) = *b.(*string), *a.(*string)
		default:
			fmt.Println("unknown type")
		}
	}
}

func InitMemo(col, row int) (dp [][]int) {
	dp = make([][]int, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
	}
	return
}

func LoopPrint(a interface{}) {
	switch a.(type) {
	case []int:
		for _, value := range a.([]int) {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	case []string:
		for _, value := range a.([]string) {
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	default:
		fmt.Println("unknown type")
	}
}

func Maxi(a ...int) (res int) {
	res = a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func Mini(a ...int) (res int) {
	res = a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AbsFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func PreInt(val int, arr []int) []int {
	return append([]int{val}, arr...)
}

func ToInt(str string) int {
	if res, err := strconv.Atoi(str); err == nil {
		return res
	} else {
		panic(err)
	}
}

func ReverseArr2(nums []int) {
	for i, size := 0, len(nums); i < size>>1; i++ {
		nums[i], nums[size-1-i] = nums[size-1-i], nums[i]
	}
}

func ReverseArr1(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func Timer(test func(), desc string) {
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("\n[%s] takes: %d\n", desc, end-start)
}

func ReverseStr(str string) string {
	arr := []rune(str)
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}

func TimeStr2Unix(str string) int64 {
	res, err := time.ParseInLocation(TIME_TEMPLATE, str, time.Local)
	if err != nil {
		fmt.Println(err)
	}
	return res.Unix()
}

func Unix2TimeStr(unix int64) string {
	return time.Unix(unix, 0).Format(TIME_TEMPLATE)
}

func Int2bit(n int) (res string) {
	for ; n > 0; n >>= 1 {
		cur := n % 2
		res = strconv.Itoa(cur) + res
	}
	return
}

func Bit2Int(n string) (res int) {
	recorder := []rune{}
	for _, value := range n {
		recorder = append(recorder, value-'0')
	}
	for len(recorder) > 0 {
		size := len(recorder) - 1
		cur := recorder[0]
		recorder = recorder[1:]
		cur <<= uint(size)
		res += int(cur)
	}
	return
}

func Room(n int) {
	unit := 1024
	//bit := n * 8
	//fmt.Printf("%d bits\n", bit)
	bytes := n
	if bytes != 0 {
		fmt.Printf("%d bytes\n", bytes)
	}
	kb := bytes / unit
	if kb != 0 {
		fmt.Printf("%d KB\n", kb)
	}
	mb := kb / unit
	if mb != 0 {
		fmt.Printf("%d MB\n", mb)
	}
	gb := mb / unit
	if gb != 0 {
		fmt.Printf("%d GB\n", gb)
	}
	tb := gb / unit
	if tb != 0 {
		fmt.Printf("%d TB\n", tb)
	}
	pb := tb / unit
	if pb != 0 {
		fmt.Printf("%d PB\n", pb)
	}
	fmt.Println()
}

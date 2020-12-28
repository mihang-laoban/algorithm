package tools

import (
	"fmt"
	. "reflect"
)

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func IsSameType(a interface{}, b interface{})  bool{
	return TypeOf(a) == TypeOf(b)
}


func Swap(a interface{}, b interface{}) {
	if IsSameType(a,  b) {
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

func InitMemo(col int, row int) (dp [][]int) {
	dp = make([][]int, col)
	for i := 0; i < col; i++ {
		dp[i] = make([]int, row)
	}
	return
}

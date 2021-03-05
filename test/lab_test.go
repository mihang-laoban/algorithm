package test

import (
	"strconv"
	"testing"
)

//func TestSwap(t *testing.T) {
//	a := 1
//	b := 2
//	fmt.Println(a, b)
//	SwapInt(&a, &b)
//	fmt.Println(a, b)
//
//	c := "a"
//	d := "b"
//	fmt.Println(c, d)
//	SwapStr(&c, &d)
//	fmt.Println(c, d)
//}
//
//func SwapInt(a *int, b *int) {
//	*a, *b = *b, *a
//}
//
//func SwapStr(a *string, b *string) {
//	*a, *b = *b, *a
//}

func MyItoa(i int) string {
	//return fmt.Sprint(i) //版本一
	return strconv.Itoa(i) //版本二
}

var r string

func BenchmarkMyItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = MyItoa(i)
	}
}

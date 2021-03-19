package alg

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2}
	b := a
	d := make([]int, 2)
	copy(d, a)
	a[0] = 3
	fmt.Println(b)
	fmt.Println(d)

	e := make([]int, 2, 3)
	e[1] = 1
	fmt.Println(len(e), cap(e))
	e = append(e, 5)
	e = append(e, 5)
	fmt.Println(len(e), cap(e))
	fmt.Println(e)

	x, y := [2]int{1, 2}, [2]int{1, 2}
	fmt.Println(x == y)
	fmt.Println(cap(x), len(y))
}

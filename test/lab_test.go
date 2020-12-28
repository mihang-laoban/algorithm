package test

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	fmt.Println(a, b)
	SwapInt(&a, &b)
	fmt.Println(a, b)

	c := "a"
	d := "b"
	fmt.Println(c, d)
	SwapStr(&c, &d)
	fmt.Println(c, d)
}

func SwapInt(a *int, b *int) {
	*a, *b = *b, *a
}

func SwapStr(a *string, b *string) {
	*a, *b = *b, *a
}

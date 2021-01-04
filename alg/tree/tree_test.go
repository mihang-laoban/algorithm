package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"testing"
)

func init() {
	Max(1, 2)
}

func TestGo(t *testing.T) {
	arr := []int{15, 6, 18, 3, 7, 17, 20, 2, 4, 13, 9, 14}
	tmp := InitTree(arr)
	tmp.InOrderTravel()
}

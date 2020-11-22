package main

import (
	"dp/ds"
)

func main() {
	tmp := ds.LinkedList{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		tmp.Append(arr[i])
	}
	tmp.Display()
	tmp.Reverse1()
	tmp.Display()
	tmp.Reverse2()
	tmp.Display()
}

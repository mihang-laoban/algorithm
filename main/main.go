package main

import (
	"dp/ds"
	"fmt"
)

func main() {
	node := ds.LinkedList{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(node.IsEmpty())
	for i := 0; i < len(arr); i++ {
		node.Append(arr[i])
	}
	node.Display()
	node.Reverse1()
	node.Display()
	node.Reverse2()
	node.Display()
	node.Prepend(0)
	node.DisplayR()
	node.Reverse2()
	node.Display()
	fmt.Println(node.IsEmpty())
	fmt.Println(node.Size())
}

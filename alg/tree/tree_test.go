package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"testing"
)

const PRE = 0
const IN = 1
const POST = 2

func init() {
	Max(1, 2)
}

func run(root *TreeNode, functions []func(*TreeNode) []int) {
	for _, fun := range functions {
		fmt.Println(fun(root))
	}
}

func traverse(order int) {
	roots := GetRootExample()
	inFunc := []func(*TreeNode) []int{InOrderLabel, InOrderRecur, InOrderLoop}
	preFunc := []func(*TreeNode) []int{PreOrderLabel, PreOrderRecur, PreOrderLoop}
	postFunc := []func(*TreeNode) []int{PostOrderLabel, PostOrderRecur, PostOrderLoop}

	for _, root := range roots {
		switch order {
		case PRE:
			run(root, preFunc)
		case IN:
			run(root, inFunc)
		case POST:
			run(root, postFunc)
		}
	}
}

func TestGo(t *testing.T) {
	fmt.Println("Pre:")
	traverse(PRE)
	fmt.Println("In:")
	traverse(IN)
	fmt.Println("Post:")
	traverse(POST)
}

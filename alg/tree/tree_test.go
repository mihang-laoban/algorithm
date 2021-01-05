package main

import (
	. "dp/ds/tree"
	. "dp/tools"
	"fmt"
	"testing"
)

func init() {
	Max(1, 2)
}

func traverse(order int) {
	preFunc := []func(*TreeNode) []int{PreOrderLabel, PreOrderRecur, PreOrderLoop, MyPreOrder}
	inFunc := []func(*TreeNode) []int{InOrderLabel, InOrderRecur, InOrderLoop, MyInOrder}
	postFunc := []func(*TreeNode) []int{PostOrderLabel, PostOrderRecur, PostOrderLoop, MyPostOrder}

	for _, root := range GetRootForTraverse() {
		switch order {
		case PRE:
			fmt.Println("Pre:")
			Ordering(root, preFunc)
		case IN:
			fmt.Println("In:")
			Ordering(root, inFunc)
		case POST:
			fmt.Println("Post:")
			Ordering(root, postFunc)
		}
	}
}

func MyPreOrder(root *TreeNode) (res []int)  { return }
func MyInOrder(root *TreeNode) (res []int)   { return }
func MyPostOrder(root *TreeNode) (res []int) { return }

func AllOrderTraverse() {
	for _, v := range []int{PRE, IN, POST} {
		traverse(v)
	}
}

func SerAndDes() {
	// "[1,2,3,nil,nil,4,5,nil,nil,nil,nil]"
	root := GetRootForSerialize()
	str := Serialize(root)
	res := Deserialize(str)
	fmt.Println(Serialize(res))
}

func TestGo(t *testing.T) {

}

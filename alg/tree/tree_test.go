package tree

import (
	"dp/ds/tree"
	"testing"
)

func TestBalanced(t *testing.T) {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	one := tree.NewTree()
	for _, value := range arr {
		if value > 0 {
			one.Insert(value)
		}
	}
	one.Traverse(tree.IN)
	one.Traverse(tree.PRE)
	one.Traverse(tree.POST)
}

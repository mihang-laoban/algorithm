package ds

import (
	"fmt"
	"testing"
)

// TreeNode 二叉树节点，采用链表的形式表示
type TreeNode struct {
	Left  *TreeNode
	Data  interface{}
	Right *TreeNode
}

type Operations interface {
	PreOrder()
	InOrder()
	PostOrder()
	IsEmpty() bool
	Insert()
	Search()
	Delete()
}

func (node *TreeNode) IsEmpty() bool { return node.Left == nil && node.Right == nil }
func (node *TreeNode) Insert()       {}
func (node *TreeNode) PreOrder()     {}
func (node *TreeNode) InOrder()      {}
func (node *TreeNode) PostOrder()    {}
func (node *TreeNode) Delete()       {}
func (node *TreeNode) Search()       {}

func TestTree(t *testing.T) {
	tree := TreeNode{}
	fmt.Println(tree.IsEmpty())
}

package tree

import "fmt"

type Node struct {
	Value  int		// 节点值
	left  *Node 	// 节点的左子叶引用
	right *Node 	// 节点的右子叶引用
}

type Tree struct {
	root *Node 		// 根节点引用
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Find(target int) *Node {
	curNode := t.root // 从根节点开始查找
	if curNode == nil {
		return nil
	}
	for curNode.Value != target {
		if curNode.Value < target {
			curNode = curNode.right
		}else {
			curNode = curNode.left
		}

		if curNode == nil {
			return nil
		}
	}
	return curNode
}

func (t *Tree) Insert(target int) {
	node := &Node{Value: target}

	if t.root == nil {
		t.root = node
		return
	}

	curNode := t.root

	for {
		if node.Value < curNode.Value {
			if curNode.left == nil {
				curNode.left = node
				return
			}
			curNode = curNode.left
		} else {
			if curNode.right == nil {
				curNode.right = node
				return
			}
			curNode = curNode.right
		}
	}
}

func (t *Tree) Delete(target int) {
	// 太复杂，暂不研究
}

const (
	PRE = iota
	IN
	POST
)

func (t *Tree) Traverse(sortType int) {
	switch sortType {
	case PRE:
		fmt.Print("前序遍历:")
		t.PreOrder(t.root)
	case IN:
		fmt.Print("中序遍历:")
		t.InOrder(t.root)
	case POST:
		fmt.Print("后序遍历:")
		t.PostOrder(t.root)
	}
}

func (t *Tree) PreOrder(node *Node) {
	if node != nil {
		fmt.Print(node.Value, " ")
		t.PreOrder(node.left)
		t.PreOrder(node.right)
	}
}
func (t *Tree) InOrder(node *Node) {
	if node != nil {
		t.InOrder(node.left)
		fmt.Print(node.Value, " ")
		t.InOrder(node.right)
	}
}
func (t *Tree) PostOrder(node *Node) {
	if node != nil {
		t.PostOrder(node.left)
		t.PostOrder(node.right)
		fmt.Print(node.Value, " ")
	}
}

func InitBinaryTree() {
	tree := NewTree()

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(9)

	tree.Traverse(PRE)
	fmt.Print("\n")
	tree.Traverse(IN)
	fmt.Print("\n")
	tree.Traverse(POST)
	fmt.Print("\n")
}

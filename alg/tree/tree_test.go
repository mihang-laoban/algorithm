package main

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	data   int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

type BiSearchTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func (bst *BiSearchTree) Add(data int) {
	bst.create = new(TreeNode)
	bst.create.data = data

	if !bst.IsEmpty() {
		bst.cur = bst.root
		for {
			if data < bst.cur.data {
				//如果要插入的值比当前节点的值小，则当前节点指向当前节点的左孩子，如果
				//左孩子为空，就在这个左孩子上插入新值
				if bst.cur.left == nil {
					bst.cur.left = bst.create
					bst.create.parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.left
				}

			} else if data > bst.cur.data {
				//如果要插入的值比当前节点的值大，则当前节点指向当前节点的右孩子，如果
				//右孩子为空，就在这个右孩子上插入新值
				if bst.cur.right == nil {
					bst.cur.right = bst.create
					bst.create.parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.right
				}

			} else {
				//如果要插入的值在树中已经存在，则退出
				return
			}
		}

	} else {
		bst.root = bst.create
		bst.root.parent = nil
	}
}

func (bst *BiSearchTree) Delete(data int) {
	var (
		deleteNode func(node *TreeNode)
		node       = bst.Search(data)
	)

	deleteNode = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.left == nil && node.right == nil {
			//如果要删除的节点没有孩子，直接删掉它就可以(毫无挂念~.~!)
			if node == bst.root {
				bst.root = nil
			} else {
				if node.parent.left == node {
					node.parent.left = nil
				} else {
					node.parent.right = nil
				}
			}

		} else if node.left != nil && node.right == nil {
			//如果要删除的节点只有左孩子或右孩子，让这个节点的父节点指向它的指针指向它的
			//孩子即可
			if node == bst.root {
				node.left.parent = nil
				bst.root = node.left
			} else {
				node.left.parent = node.parent
				if node.parent.left == node {
					node.parent.left = node.left
				} else {
					node.parent.right = node.left
				}
			}

		} else if node.left == nil && node.right != nil {
			if node == bst.root {
				node.right.parent = nil
				bst.root = node.right
			} else {
				node.right.parent = node.parent
				if node.parent.left == node {
					node.parent.left = node.right
				} else {
					node.parent.right = node.right
				}
			}

		} else {
			//如果要删除的节点既有左孩子又有右孩子，就把这个节点的直接后继的值赋给这个节
			//点，然后删除直接后继节点即可
			successor := bst.GetSuccessor(node.data)
			node.data = successor.data
			deleteNode(successor)
		}
	}
	deleteNode(node)
}
func (bst BiSearchTree) GetRoot() *TreeNode {
	if bst.root != nil {
		return bst.root
	}
	return nil
}

func (bst BiSearchTree) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

func (bst BiSearchTree) InOrderTravel() {
	var inOrderTravel func(node *TreeNode)

	inOrderTravel = func(node *TreeNode) {
		if node != nil {
			inOrderTravel(node.left)
			fmt.Printf("%d ", node.data)
			inOrderTravel(node.right)
		}
	}

	inOrderTravel(bst.root)
}

func (bst BiSearchTree) PreOrderTravel() {
	var PreOrderTravel func(node *TreeNode)

	PreOrderTravel = func(node *TreeNode) {
		if node != nil {
			fmt.Printf("%d ", node.data)
			PreOrderTravel(node.left)
			PreOrderTravel(node.right)
		}
	}

	PreOrderTravel(bst.root)
}

func (bst BiSearchTree) PostOrderTravel() {
	var PostOrderTravel func(node *TreeNode)

	PostOrderTravel = func(node *TreeNode) {
		if node != nil {
			PostOrderTravel(node.left)
			PostOrderTravel(node.right)
			fmt.Printf("%d ", node.data)
		}
	}

	PostOrderTravel(bst.root)
}

func (bst BiSearchTree) Search(data int) *TreeNode {
	//和Add操作类似，只要按照比当前节点小就往左孩子上拐，比当前节点大就往右孩子上拐的思路
	//一路找下去，知道找到要查找的值返回即可
	bst.cur = bst.root
	for {
		if bst.cur == nil {
			return nil
		}

		if data < bst.cur.data {
			bst.cur = bst.cur.left
		} else if data > bst.cur.data {
			bst.cur = bst.cur.right
		} else {
			return bst.cur
		}
	}
}

func (bst BiSearchTree) GetDeepth() int {
	var getDepth func(node *TreeNode) int

	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.left == nil && node.right == nil {
			return 1
		}
		var (
			lDepth = getDepth(node.left)
			rDepth = getDepth(node.right)
		)
		if lDepth > rDepth {
			return lDepth + 1
		} else {
			return rDepth + 1
		}
	}

	return getDepth(bst.root)
}

func (bst BiSearchTree) GetMin() int {
	//根据二叉查找树的性质，树中最左边的节点就是值最小的节点
	if bst.root == nil {
		return -1
	}
	bst.cur = bst.root
	for {
		if bst.cur.left != nil {
			bst.cur = bst.cur.left
		} else {
			return bst.cur.data
		}
	}
}

func (bst BiSearchTree) GetMax() int {
	//根据二叉查找树的性质，树中最右边的节点就是值最大的节点
	if bst.root == nil {
		return -1
	}
	bst.cur = bst.root
	for {
		if bst.cur.right != nil {
			bst.cur = bst.cur.right
		} else {
			return bst.cur.data
		}
	}
}

func (bst BiSearchTree) GetPredecessor(data int) *TreeNode {
	getMax := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.right != nil {
				node = node.right
			} else {
				return node
			}
		}
	}

	node := bst.Search(data)
	if node != nil {
		if node.left != nil {
			//如果这个节点有左孩子，那么它的直接前驱就是它左子树的最右边的节点，因为比这
			//个节点值小的节点都在左子树，而这些节点中值最大的就是这个最右边的节点
			return getMax(node.left)
		} else {
			//如果这个节点没有左孩子，那么就沿着它的父节点找，知道某个父节点的父节点的右
			//孩子就是这个父节点，那么这个父节点的父节点就是直接前驱
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.right {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (bst BiSearchTree) GetSuccessor(data int) *TreeNode {
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.left != nil {
				node = node.left
			} else {
				return node
			}
		}
	}

	//参照寻找直接前驱的函数对比着看
	node := bst.Search(data)
	if node != nil {
		if node.right != nil {
			return getMin(node.right)
		} else {
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.left {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (bst *BiSearchTree) Clear() {
	bst.root = nil
	bst.cur = nil
	bst.create = nil
}

/*
			15
		   /  \
		  6    18
		 / \   / \
		3   7 17  20
	   / \   \
	  2   4   13
			  / \
			 9   14
*/

func Test(t *testing.T) {
	var bst BiSearchTree
	arr := []int{15, 6, 18, 3, 7, 17, 20, 2, 4, 13, 9, 14}
	for _, v := range arr {
		bst.Add(v)
	}

	fmt.Println("[in]The nodes of the BiSearchTree is: ")
	bst.InOrderTravel()
	fmt.Println()
	fmt.Println("[post]The nodes of the BiSearchTree is: ")
	bst.PostOrderTravel()
	fmt.Println()
	fmt.Println("[pre]The nodes of the BiSearchTree is: ")
	bst.PreOrderTravel()
	fmt.Println()

	fmt.Printf("The deepth of the tree is: %d\n", bst.GetDeepth())
	fmt.Printf("The min is: %d\n", bst.GetMin())
	fmt.Printf("The max is: %d\n", bst.GetMax())

	if bst.Search(17) != nil {
		fmt.Printf("The 17 exists.\n")
	}

	fmt.Printf("root node's predecessor is: %d\n", bst.GetPredecessor(bst.GetRoot().data).data)
	fmt.Printf("root node's successor is: %d\n", bst.GetSuccessor(bst.GetRoot().data).data)

	bst.Delete(13)
	fmt.Printf("Nodes after delete the 13: ")
	bst.InOrderTravel()
	fmt.Printf("\n")
}

package tree

import "fmt"

type HoffNode struct {
	Value int   // 节点值
	left  *HoffNode // 节点的左子叶引用
	right *HoffNode // 节点的右子叶引用

	Weight   int   // 链表排序值
	next   *HoffNode // 链表下个节点引用
}

type List struct {
	first  *HoffNode // 链表首节点
	length int   // 链表长度
}

func NewList() *List {
	return &List{}
}

func (l *List) Insert(node *HoffNode) {
	if l.first == nil {
		l.first = node
	}else {
		var preHoffNode *HoffNode
		var curHoffNode = l.first

		for curHoffNode.Weight < node.Weight  {
			preHoffNode = curHoffNode
			if curHoffNode.next == nil {
				curHoffNode = nil
				break
			}else {
				curHoffNode = curHoffNode.next
			}
		}

		if preHoffNode == nil {
			node.next = l.first
			l.first = node
		} else if curHoffNode == nil {
			preHoffNode.next = node
		} else {
			preHoffNode.next = node
			node.next = curHoffNode
		}
	}

	l.length++
}

func (l *List) PopFront() *HoffNode{
	node := l.first
	l.first = l.first.next
	l.length--
	return node
}

type HoffTree struct {
	root *HoffNode 		// 根节点引用
}

func NewHoffTree(l *List) *HoffTree {
	for l.length > 1  {
		node := &HoffNode{}
		node.left, node.right = l.PopFront(), l.PopFront()
		node.Weight = node.left.Weight + node.right.Weight
		l.Insert(node)
	}
	return &HoffTree{l.first}
}

func (t *HoffTree) MiddleWeight(node *HoffNode) {
	if node != nil {
		t.MiddleWeight(node.left)
		if node.Value > 0 {
			fmt.Print(node.Value, " ")
		}
		t.MiddleWeight(node.right)
	}
}

func InitHoffmanTree() {
	list := NewList()

	list.Insert(&HoffNode{Value:10, Weight:2})
	list.Insert(&HoffNode{Value:20, Weight:5})
	list.Insert(&HoffNode{Value:30, Weight:9})
	list.Insert(&HoffNode{Value:40, Weight:3})

	tree := NewHoffTree(list)

	tree.MiddleWeight(tree.root)
}
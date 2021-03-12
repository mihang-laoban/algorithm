package newCoder

import (
	"fmt"
	"testing"
)

func TestLRU_coder(t *testing.T) {
	fmt.Println(LRU([][]int{[]int{1, 1, 1}, []int{1, 2, 2}, []int{1, 3, 2}, []int{2, 1}, []int{1, 4, 4}, []int{2, 2}}, 3))
}

//type node struct {
//	val  int
//	pre  *node
//	next *node
//}
//
//type Lru struct {
//	recorder map[int]*node
//	head     *node
//	tail     *node
//	size     int
//	cap      int
//}
//
//func (l *Lru) set(k, v int) {
//	// 判断节点是否存在，如果存在则更新节点，并把节点位置移动到链表头部
//	if one, ok := l.recorder[k]; ok {
//		one.val = v
//		l.toFirst(one)
//	} else {
//		// 如果节点不存在，则判断是否超出容量，超出则删除最后一个节点，没有则直接添加
//		cur := &node{val: v}
//
//		if l.size > len(l.recorder) {
//			l.recorder[k] = cur
//			l.toFirst(cur)
//			return
//		}
//
//		if l.head == nil {
//			l.head = cur
//		} else {
//			l.head.pre = cur
//			cur.next = l.head
//		}
//	}
//}
//
//func (l *Lru) get(k int) int {
//	if one, ok := l.recorder[k]; ok {
//		one.pre.next = one.next
//		one.next.pre = one.pre
//		l.toFirst(one)
//		return one.val
//	}
//	return -1
//}
//
//func (l *Lru) toFirst(n *node) {
//	n.next = l.head.next
//	l.head.next.pre = n
//	l.head.next = n
//	n.pre = l.head
//}
//
func LRU(operators [][]int, k int) []int {
	res := []int{}
	lru := Constructor(k)
	for _, op := range operators {
		if op[0] == 1 {
			lru.set(op[1], op[2])
		} else if op[0] == 2 {
			res = append(res, lru.get(op[1]))
		}
	}
	return res
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{key: 0, value: 0},
		tail:     &DLinkedNode{key: 0, value: 0},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) set(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &DLinkedNode{key: key, value: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

package ds

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

type DLinkedNode struct {
	Key, value int
	pre, next  *DLinkedNode
}

func LRUConstructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{Key: 0, value: 0},
		tail:     &DLinkedNode{Key: 0, value: 0},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	cur := this.cache[key]
	this.MoveToHead(cur)
	return cur.value
}

func (this *LRUCache) Set(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &DLinkedNode{Key: key, value: value}
		this.cache[key] = node
		this.AddToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.RemoveTail()
			delete(this.cache, removed.Key)
			this.size--
		}
	} else {
		cur := this.cache[key]
		cur.value = value
		this.MoveToHead(cur)
	}
}

func (this *LRUCache) AddToHead(cur *DLinkedNode) {
	cur.pre = this.head
	cur.next = this.head.next
	this.head.next.pre = cur
	this.head.next = cur
}

func (this *LRUCache) RemoveNode(cur *DLinkedNode) {
	cur.pre.next = cur.next
	cur.next.pre = cur.pre
}

func (this *LRUCache) MoveToHead(cur *DLinkedNode) {
	this.RemoveNode(cur)
	this.AddToHead(cur)
}

func (this *LRUCache) RemoveTail() *DLinkedNode {
	cur := this.tail.pre
	this.RemoveNode(cur)
	return cur
}

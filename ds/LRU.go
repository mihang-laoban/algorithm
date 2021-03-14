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
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if cur, ok := this.cache[key]; ok {
		this.MoveToHead(cur)
		return cur.value
	}
	return -1
}

func (this *LRUCache) Set(key int, value int) {
	if cur, ok := this.cache[key]; ok {
		cur.value = value
		this.MoveToHead(cur)
	} else {
		node := &DLinkedNode{Key: key, value: value}
		this.cache[key] = node
		this.AddToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.RemoveTail()
			delete(this.cache, removed.Key)
			this.size--
		}
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

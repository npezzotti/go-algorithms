package lru_cache

type doublyLinkedNode struct {
	key  int
	val  int
	prev *doublyLinkedNode
	next *doublyLinkedNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*doublyLinkedNode
	head     *doublyLinkedNode
	tail     *doublyLinkedNode
}

func Constructor(capacity int) LRUCache {
	head := &doublyLinkedNode{key: 0, val: 0}
	tail := &doublyLinkedNode{key: 0, val: 0}
	head.next = tail
	tail.prev = head

	return LRUCache{
		cache:    make(map[int]*doublyLinkedNode, capacity),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.removeNode(node)
		this.moveToHead(node)
	} else {
		if len(this.cache) == this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
		}
		newNode := &doublyLinkedNode{key: key, val: value}
		this.addToHead(newNode)
		this.cache[newNode.key] = newNode
	}
}

func (this *LRUCache) moveToHead(node *doublyLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *doublyLinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *doublyLinkedNode) {
	next := node.next
	prev := node.prev
	next.prev = prev
	prev.next = next
}

func (this *LRUCache) removeTail() *doublyLinkedNode {
	tail := this.tail.prev
	this.removeNode(tail)
	return tail
}

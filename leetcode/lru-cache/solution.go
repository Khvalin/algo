package main

type node struct {
	key  int
	prev *node
	next *node
}

type LRUCache struct {
	cache    map[int]int
	capacity int
	keys     map[int]*node
	size     int
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: map[int]int{}, keys: map[int]*node{}, capacity: capacity}
}

func (this *LRUCache) bumpKey(key int) {

	if n, f := this.keys[key]; f {
		if n.prev != nil {
			n.prev.next = n.next
		}

		if n == this.head {
			this.head = this.head.next
			this.head.next.prev = nil
		}

		if n.next != nil {
			n.next.prev = n.prev
		}

		if n == this.head {
			this.head = this.head.next
		}

		n.prev = this.tail
		this.tail.next = n
		this.tail = n

		return
	}

	keyNode := &node{key: key}
	this.keys[key] = keyNode

	if this.tail == nil {
		this.head = keyNode
		this.tail = keyNode

		return
	}

	this.size++
	this.tail.next = keyNode
	keyNode.prev = this.tail
	this.tail = keyNode

	if this.size == this.capacity {
		v := this.head.key

		delete(this.keys, v)
		delete(this.cache, v)

		this.head.next.prev = nil
		this.head = this.head.next
		this.size--
	}
}

func (this *LRUCache) Get(key int) int {
	if v, f := this.cache[key]; f {
		this.bumpKey(key)
		return v
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.bumpKey(key)
	this.cache[key] = value
}

func main() {
	/*

	   ["LRUCache","put","put","get","get","put","get","get","get"]
	   [[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]

	*/

	cache := Constructor(2 /* capacity */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // returns 1
	cache.Put(3, 3) // evicts key 2
	// cache.Get(2)    // returns -1 (not found)
	// cache.Put(4, 4) // evicts key 1
	// cache.Get(1)    // returns -1 (not found)
	// cache.Get(3)    // returns 3
	// cache.Get(4)    // returns 4
}

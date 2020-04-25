package main

import (
	"fmt"
)

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
	headNode := &node{key: -1}
	return LRUCache{cache: map[int]int{}, keys: map[int]*node{}, capacity: capacity, head: headNode, tail: headNode}
}

func (this *LRUCache) bumpKey(key int) {
	if n, f := this.keys[key]; f {
		if n == this.tail {
			return
		}

		keyNode := &node{key: key, prev: this.tail}
		this.keys[key] = keyNode

		this.tail.next = keyNode
		this.tail = keyNode

		fmt.Println(n.prev.key)
		n.next.prev = n.prev
		n.prev.next = n.next

		return
	}

	keyNode := &node{key: key}
	this.keys[key] = keyNode

	this.size++
	this.tail.next = keyNode
	keyNode.prev = this.tail
	this.tail = keyNode

	if this.size > this.capacity {
		v := this.head.next.key

		delete(this.keys, v)
		delete(this.cache, v)

		this.head.next = this.head.next.next
		this.head.next.prev = this.head

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

func (this *LRUCache) String() string {
	r := ""
	n := this.head
	for n != nil {
		r += fmt.Sprintf(" %d", n.key)
		n = n.next
	}

	return r
}

func main() {
	/*

	   ["LRUCache","put","put","get","get","put","get","get","get"]
	   [[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]

	*/

	cache := Constructor(3 /* capacity */)

	for i := 1; i < 5; i++ {
		cache.Put(i, i)
		fmt.Println(cache.String())
	}

	for i := 4; i > 0; i-- {
		cache.Get(2)
		fmt.Println(cache.String())
	}
}

package lru

import (
	"fmt"
	"testing"
)

type LRUCache struct {
	size       int
	cache      map[int]*DLinkList
	head, tail *DLinkList
}

type DLinkList struct {
	key, value int
	prev, next *DLinkList
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:  capacity,
		cache: map[int]*DLinkList{},
		head:  initDLinKList(0, 0),
		tail:  initDLinKList(0, 0),
	}
	// 设置首位
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func initDLinKList(key, value int) *DLinkList {
	return &DLinkList{
		key:   key,
		value: value,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	// moveToHead
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) moveToHead(node *DLinkList) {
	// removeNode
	this.removeNode(node)
	// addToHead
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := &DLinkList{key: key, value: value}
		this.cache[key] = node
		this.addToHead(node)
		if len(this.cache) > this.size {
			removed := this.removeTail()
			delete(this.cache, removed.key)
		}
	} else {
		//  添加新值，移动到 head
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeNode(node *DLinkList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkList) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeTail() *DLinkList {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func TestLRU(t *testing.T) {
	capacity := 2
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得关键字 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}

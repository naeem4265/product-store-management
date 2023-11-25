package main

import "fmt"

type LRUCache struct {
	head, tail *node
	ref        [102]*node
	capacity   int
}

type node struct {
	name   string
	id     int
	parent *node
	child  *node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	for i := 0; i < 102; i++ {
		lru.ref[i] = &node{id: -1}
	}
	return lru
}

func (lru *LRUCache) Get(key int) string {
	if lru.ref[key].id == -1 {
		return "Empty"
	}
	if lru.tail == lru.ref[key] {
		return lru.ref[key].name
	}
	if lru.head == lru.ref[key] {
		lru.head = lru.head.child
		lru.ref[key].parent = lru.tail
		lru.tail.child = lru.ref[key]
		lru.ref[key].child = nil
		lru.tail = lru.ref[key]
	} else {
		lru.ref[key].parent.child = lru.ref[key].child
		lru.ref[key].child.parent = lru.ref[key].parent
		lru.tail.child = lru.ref[key]
		lru.ref[key].parent = lru.tail
		lru.ref[key].child = nil
		lru.tail = lru.ref[key]
	}
	return lru.ref[key].name
}

func (lru *LRUCache) Listen(id int, name string) {
	if lru.head == nil {
		lru.head = lru.ref[id]
		lru.tail = lru.head
		lru.ref[id].id = id
		lru.ref[id].name = name
		lru.ref[id].parent = nil
		lru.ref[id].child = nil
		lru.capacity--
	} else if lru.ref[id].id != -1 {
		lru.ref[id].id = id
		lru.ref[id].name = name
		if lru.tail == lru.ref[id] {
			return
		}
		if lru.head == lru.ref[id] {
			lru.head = lru.head.child
			lru.ref[id].parent = lru.tail
			lru.tail.child = lru.ref[id]
			lru.ref[id].child = nil
			lru.tail = lru.ref[id]
		} else {
			lru.ref[id].parent.child = lru.ref[id].child
			lru.ref[id].child.parent = lru.ref[id].parent
			lru.tail.child = lru.ref[id]
			lru.ref[id].parent = lru.tail
			lru.ref[id].child = nil
			lru.tail = lru.ref[id]
		}
	} else {
		lru.ref[id].id = id
		lru.ref[id].name = name
		lru.ref[id].parent = lru.tail
		lru.ref[id].child = nil
		lru.tail.child = lru.ref[id]
		lru.tail = lru.ref[id]
		if lru.capacity == 0 {
			lru.head.id = -1
			temp := lru.head
			lru.head = lru.head.child
			temp.child = nil
			lru.head.parent = nil
		} else {
			lru.capacity--
		}
	}
}

func main() {
	obj := Constructor(100)
	obj.Listen(1, "song1")
	obj.Listen(2, "song2")
	obj.Listen(3, "song3")
	fmt.Println(obj.Get(1))
	obj.Listen(4, "song4")
	fmt.Println(obj.Get(3))
}

package main

import (
	"container/list"
	"fmt"
)

type Pair struct {
	key, value int
}

type LRUCache struct {
	cache map[int]*list.Element
	l     *list.List
	size  int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.size = capacity
	lru.cache = make(map[int]*list.Element)
	lru.l = list.New()
	return lru
}

func (this *LRUCache) Get(key int) int {
	if res, found := this.cache[key]; found {
		cur := res.Value.(Pair)
		key := cur.key
		val := cur.value
		this.l.Remove(res)
		this.l.PushBack(Pair{key, val})
		this.cache[key] = this.l.Back()
		return cur.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	if res, found := this.cache[key]; found {
		this.l.Remove(res)
	}
	this.l.PushBack(Pair{key, value})
	bk := this.l.Back()

	this.cache[key] = bk
	if len(this.cache) > this.size {
		ft := this.l.Front()
		p := ft.Value.(Pair)
		this.l.Remove(ft)
		delete(this.cache, p.key)
	}
}

func show(root *list.Element) {
	for root != nil {
		fmt.Printf("(%d , %d)--", root.Value.(Pair).key, root.Value.(Pair).value)
		root = root.Next()
	}
	fmt.Println()
}

func main() {
	cache := Constructor(3)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	show(cache.l.Front())
	cache.Put(5, 6)
	show(cache.l.Front())
	fmt.Println(cache.Get(1))
	show(cache.l.Front())
	fmt.Println(cache.Get(2))
	show(cache.l.Front())
	fmt.Println(cache.Get(4))
	show(cache.l.Front())
	cache.Put(2, 4)
	show(cache.l.Front())
	fmt.Println(cache.Get(5))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
}

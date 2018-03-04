package main

// 借由自带的heap实现
import (
	"container/heap"
	"fmt"
)

type PriorityQueue struct {
	itemQueue ItemListInter
}

func NewPrioirtyQueue(ilt ItemListInter) PriorityQueue {
	pq := PriorityQueue{}
	pq.itemQueue = ilt
	heap.Init(pq.itemQueue)
	return pq
}

func (pq *PriorityQueue) PushItem(it *Item) {
	//pq.itemQueue.Show()
	heap.Push(pq.itemQueue, it)
}

func (pq *PriorityQueue) PopItem() *Item {
	res := heap.Pop(pq.itemQueue).(*Item)
	return res
}

func (pq *PriorityQueue) FrontItem() *Item {
	res := pq.itemQueue.Front()
	return res.(*Item)
}

func (pq *PriorityQueue) Len() int {
	return pq.itemQueue.Len()
}

type Item struct {
	Value interface{}
	Score float64
	Index int
}

type ItemList []*Item

func (s ItemList) Len() int { return len(s) }

func (s ItemList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].Index = i
	s[j].Index = j
}

func (s *ItemList) Push(x interface{}) {
	n := len(*s)
	item := x.(*Item)
	(*item).Index = n
	*s = append(*s, item)
}

func (s *ItemList) Pop() interface{} {
	old := *s
	n := len(old)
	item := old[n-1]
	item.Index = n
	*s = old[0 : n-1]
	return item
}

func (s ItemList) Show() {
	fmt.Println("show:")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func (s ItemList) Front() {
	return s[0]
}

type ItemListInter interface {
	Len() int
	Swap(int, int)
	Push(interface{})
	Pop() interface{}
	Less(int, int) bool
	Show()
	Front() interface{}
}

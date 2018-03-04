package main

import "fmt"

type QueryMid struct {
	pq_max PriorityQueue
	pq_min PriorityQueue
	mid    int
}

type MaxDataList struct {
	ItemList
}

func (s MaxDataList) Less(i, j int) bool {
	return s.ItemList[i].Score > s.ItemList[j].Score
}

type MinDataList struct {
	ItemList
}

func (s MinDataList) Less(i, j int) bool {
	return s.ItemList[i].Score < s.ItemList[j].Score
}

func NewMaxDataList() MaxDataList {
	dl := MaxDataList{}
	dl.ItemList = make(ItemList, 0)
	return dl
}

func NewMinDataList() MinDataList {
	dl := MinDataList{}
	dl.ItemList = make(ItemList, 0)
	return dl
}

func NewWorker() QueryMid {
	nadl := NewMaxDataList()
	nidl := NewMinDataList()
	pq_max := NewPrioirtyQueue(&nadl)
	pq_min := NewPrioirtyQueue(&nidl)
	qm := QueryMid{pq_max: pq_max, pq_min: pq_min}
	return qm
}

func (s *QueryMid) Insert(x int) {
	if x >= s.mid {
		s.pq_min.PushItem(&Item{Value: x, Score: float64(x)})
	} else {
		s.pq_max.PushItem(&Item{Value: x, Score: float64(x)})
	}
	sub := s.pq_max.Len() - s.pq_min.Len()
	fmt.Println(sub)
	if sub <= -2 {
		res := s.pq_min.PopItem()
		resValue := res.Value.(int)
		s.pq_max.PushItem(&Item{Value: s.mid, Score: float64(s.mid)})
		s.mid = resValue
	} else if sub >= 2 {
		res := s.pq_max.PopItem()
		resValue := res.Value.(int)
		s.pq_min.PushItem(&Item{Value: s.mid, Score: float64(s.mid)})
		s.mid = resValue
	}
}

func (s *QueryMid) InsertFirst(x int) {
	s.mid = x
}

func (s QueryMid) Query() int {
	if s.pq_min.Len() == s.pq_max.Len() {
		return s.mid
	} else if s.pq_min.Len() < s.pq_max.Len() {

	}

}

func main() {
	qm := NewWorker()
	qm.InsertFirst(1)
	qm.Insert(2)
	qm.Insert(5)
	fmt.Println(qm.Query())
	qm.Insert(3)
	qm.Insert(4)
	fmt.Println(qm.Query())
}

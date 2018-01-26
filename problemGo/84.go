package main

import (
	"container/list"
	"errors"
	"fmt"
)

//Given n non-negative integers representing the histogram's bar height where the width of each bar
// is 1, find the area of largest rectangle in the histogram.
//
//
//Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
//
//The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
//For example,
//Given heights = [2,1,5,6,2,3],
//return 10.

type pair struct {
	v, idx int
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	ls := list.New()
	return &Stack{ls}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			s.list.Remove(e)
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}

}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

type OrderStack struct {
	Stack
}

func (s *OrderStack) Push(value pair) *pair {
	var last *pair
	last = nil
	for !s.Empty() {
		t, _ := s.Top()
		eTop := t.(pair)
		last = &eTop
		if eTop.v >= value.v {
			s.Pop()
		} else {
			break
		}
	}

	s.Stack.Push(value)
	return last
}

func NewOrderStack() *OrderStack {
	o := OrderStack{*NewStack()}
	return &o
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	//right := make([]int, n)
	ostk := NewOrderStack()
	for idx, v := range heights {
		fmt.Println(idx, v)
		res := ostk.Push(pair{idx: idx, v: v})
		if res == nil {
			left[idx] = 1
		} else {
			fmt.Println(res)
			left[idx] = left[res.idx] + idx - res.idx
		}
	}
	fmt.Println(left)
	return 0
}

func main() {
	a := []int{9, 1, 4, 5, 6, 5, 1, 4, 2}
	largestRectangleArea(a)
	//st := NewOrderStack()
	//for idx, i := range(a) {
	//	st.Push(pair{v:i, idx:idx})
	//	fmt.Println(idx, i)
	//}
	//for !st.Empty() {
	//	r, _ := st.Pop()
	//	fmt.Println(r.(pair))
	//}

}

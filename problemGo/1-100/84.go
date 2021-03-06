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
		if eTop.v >= value.v {
			last = &eTop
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

func work(height []int, st, step int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	n := len(height)
	ans := make([]int, n)
	ostk := NewOrderStack()
	for i := st; i >= 0 && i < n; i += step {
		res := ostk.Push(pair{idx: i, v: height[i]})
		if res == nil {
			ans[i] = 1
		} else {
			ans[i] = ans[res.idx] + abs(i-res.idx)
		}
	}
	fmt.Println(ans)
	return ans
}

func largestRectangleArea(heights []int) int {
	left := work(heights, 0, 1)
	right := work(heights, len(heights)-1, -1)
	maxx := 0
	for i := 0; i < len(heights); i++ {
		s := (right[i] + left[i] - 1) * heights[i]
		if s > maxx {
			maxx = s
		}
	}
	return maxx
}

func main() {
	a := []int{9, 1, 4, 5, 6, 5, 1, 4, 2}
	res := largestRectangleArea(a)
	fmt.Println(res)
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

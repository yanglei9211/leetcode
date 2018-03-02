package main

import (
	"container/list"
	"errors"
	"fmt"
)

//Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's
// and return its area.
//
//For example, given the following matrix:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//Return 6.

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
	return ans
}

func gt(heights []int) int {
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

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	down := make([][]int, n)
	for i := 0; i < n; i++ {
		down[i] = make([]int, m)
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				down[i][j] = 0
			} else {
				if i+1 < n {
					down[i][j] = down[i+1][j] + 1
				} else {
					down[i][j] = 1
				}
			}
		}
	}
	ans := 0
	for _, x := range down {
		res := gt(x)
		if res > ans {
			ans = res
		}
	}
	return ans
}

func main() {
	//[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	a := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

	res := maximalRectangle(a)
	fmt.Println(res)
}

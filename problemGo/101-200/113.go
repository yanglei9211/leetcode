package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type ss struct {
	cur []int
	ans [][]int
	rt  *TreeNode
}

func (s *ss) addAns() {
	tp := make([]int, 0, len(s.cur))
	for _, x := range s.cur {
		tp = append(tp, x)
	}
	s.ans = append(s.ans, tp)
}

func (s *ss) do(rt *TreeNode, sum int) {
	if rt == nil {
		return
	}
	fmt.Println(rt.Val, sum)
	if sum == rt.Val && rt.Left == nil && rt.Right == nil {
		s.cur = append(s.cur, rt.Val)
		s.addAns()
		s.cur = s.cur[:len(s.cur)-1]
		return
	}
	s.cur = append(s.cur, rt.Val)
	s.do(rt.Left, sum-rt.Val)
	s.do(rt.Right, sum-rt.Val)
	s.cur = s.cur[:len(s.cur)-1]
	return
}

func pathSum(root *TreeNode, sum int) [][]int {
	s := ss{}
	s.rt = root
	s.do(root, sum)
	fmt.Println(s.ans)
	return s.ans
}

func inorderTraversal(r *TreeNode) {
	if r == nil {
		return
	}
	inorderTraversal(r.Left)
	fmt.Printf("%d ", r.Val)
	inorderTraversal(r.Right)
}

func main() {
	r7 := TreeNode{7, nil, nil}
	r2 := TreeNode{2, nil, nil}
	r11 := TreeNode{11, &r7, &r2}
	r4 := TreeNode{4, &r11, nil}
	r5 := TreeNode{5, nil, nil}
	r1 := TreeNode{1, nil, nil}
	r42 := TreeNode{4, &r5, &r1}
	r13 := TreeNode{13, nil, nil}
	r8 := TreeNode{8, &r13, &r42}
	r52 := TreeNode{5, &r4, &r8}
	//inorderTraversal(&r52)
	pathSum(&r52, 22)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type S struct {
	l1 []int
	l2 []int
}

func (s *S) LRTraversal(rt *TreeNode) {
	if rt == nil {
		s.l1 = append(s.l1, -1)
		return
	}
	s.LRTraversal(rt.Left)

	s.LRTraversal(rt.Right)
	s.l1 = append(s.l1, rt.Val)
}

func (s *S) RLTraversal(rt *TreeNode) {
	if rt == nil {
		s.l2 = append(s.l2, -1)
		return
	}
	s.RLTraversal(rt.Right)

	s.RLTraversal(rt.Left)
	s.l2 = append(s.l2, rt.Val)
}

func isSymmetric(root *TreeNode) bool {
	s := S{}
	s.LRTraversal(root)
	s.RLTraversal(root)
	fmt.Println(s.l1)
	fmt.Println(s.l2)
	for i := 0; i < len(s.l1); i++ {
		if s.l1[i] != s.l2[i] {
			return false
		}
	}

	return true
}

func main() {
	t1 := TreeNode{Val: 1}
	t21 := TreeNode{Val: 2}
	t22 := TreeNode{Val: 2}
	t31 := TreeNode{Val: 3}
	t32 := TreeNode{Val: 3}
	//t41 := TreeNode{Val:4}
	//t42 := TreeNode{Val:4}
	//t1.Left = &t21
	//t21.Left = &t31
	//t21.Right = &t41
	//t1.Right = &t22
	//t22.Left = &t42
	//t22.Right = &t32
	t1.Left = &t21
	t21.Left = &t31
	t1.Right = &t32
	t32.Left = &t22
	isSymmetric(&t1)
}

package main

import "fmt"

//Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
//For example,
//If n = 4 and k = 2, a solution is:
//
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
type sss struct {
	n, k int
	ans  [][]int
	a    []int
	vis  []bool
}

func (s *sss) dfs(i int) {
	if i == s.k {
		tmp := make([]int, len(s.a))
		copy(tmp, s.a)
		s.ans = append(s.ans, tmp)
		return
	}
	var st int
	if i > 0 {
		st = s.a[i-1]
	} else {
		st = 1
	}
	for j := st; j <= s.n; j++ {
		if !s.vis[j] {
			s.a[i] = j
			s.vis[j] = true
			s.dfs(i + 1)
			s.vis[j] = false
		}
	}
}

func (s *sss) init(n, k int) {
	s.n = n
	s.k = k
	s.ans = make([][]int, 0)
	s.a = make([]int, k)
	s.vis = make([]bool, n+1)
}

//var ans [][]int
//
//func dfs(n, k,  i int, a []int, vis []bool) {
//	if i == k {
//		tmp := make([]int, len(a))
//		copy(tmp, a)
//		ans = append(ans, tmp)
//		return
//	}
//	var st int
//	if i > 0 {
//		st = a[i-1]
//	} else {
//		st = 1
//	}
//	for j := st; j <= n; j++ {
//		if !vis[j] {
//			a[i] = j
//			vis[j] = true
//			dfs(n, k, i+1, a, vis)
//			vis[j] = false
//		}
//	}
//}

func combine(n int, k int) [][]int {
	worker := sss{}
	worker.init(n, k)
	worker.dfs(0)
	return worker.ans
}

func main() {
	res := combine(4, 2)
	fmt.Println(res)
}

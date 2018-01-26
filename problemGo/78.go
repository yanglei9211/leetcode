package main

import "fmt"

//Given a set of distinct integers, nums, return all possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//For example,
//If nums = [1,2,3], a solution is:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]

type ss struct {
	n    int
	ans  [][]int
	a    []int
	vis  []bool
	nums []int
}

func (s *ss) init(nums []int) {
	s.n = len(nums)
	s.a = make([]int, s.n)
	s.vis = make([]bool, s.n)
	s.nums = nums
	s.ans = make([][]int, 0)
	s.ans = append(s.ans, []int{})
}

func (s *ss) dfs(i int) {
	if i > 0 {
		tmp := make([]int, 0)
		for x := 0; x < i; x++ {
			tmp = append(tmp, s.nums[s.a[x]])
		}
		s.ans = append(s.ans, tmp)
	}
	var st int
	if i == 0 {
		st = 0
	} else {
		st = s.a[i-1]
	}
	for j := st; j < s.n; j++ {
		if !s.vis[j] {
			s.vis[j] = true
			s.a[i] = j
			s.dfs(i + 1)
			s.vis[j] = false
		}
	}

}

func subsets(nums []int) [][]int {
	worker := ss{}
	worker.init(nums)
	worker.dfs(0)
	return worker.ans
}

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

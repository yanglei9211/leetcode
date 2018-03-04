package main

import "fmt"

//Given a collection of integers that might contain duplicates, nums, return all
// possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//For example,
//If nums = [1,2,2], a solution is:
//
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//]

type worker struct {
	nums []int
	pool map[int]int
	cur  []int
	cn   int
	ans  [][]int
}

func (s *worker) init(nums []int) {

	s.pool = make(map[int]int)
	for _, x := range nums {
		s.pool[x]++
	}
	s.nums = make([]int, 0, len(s.pool))
	for x := range s.pool {
		s.nums = append(s.nums, x)
	}
	s.cn = len(s.nums)
	s.cur = make([]int, len(s.nums))
}

func (s *worker) save() {
	tmp := make([]int, 0)
	for i, j := range s.cur {
		x := s.nums[i]
		for k := 0; k < j; k++ {
			tmp = append(tmp, x)
		}
	}
	s.ans = append(s.ans, tmp)
}

func (s *worker) dfs(x int) {
	if x == s.cn {
		s.save()
		return
	}
	y := s.nums[x]
	cnt := s.pool[y]
	for j := 0; j <= cnt; j++ {
		s.cur[x] = j
		s.dfs(x + 1)
	}
}

func subsetsWithDup(nums []int) [][]int {
	wk := worker{}
	wk.init(nums)
	wk.dfs(0)
	return wk.ans
}

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}

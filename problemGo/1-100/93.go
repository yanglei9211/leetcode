package main

import "fmt"

//Given a string containing only digits, restore it by returning all possible valid IP address combinations.
//
//For example:
//Given "25525511135",
//
//return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)

type Solve struct {
	s   string
	ans []string
	cur string
}

func (s *Solve) dfs(i, j int) {
	if j == 4 {
		if i == len(s.s) {
			s.ans = append(s.ans, s.cur[0:len(s.cur)-1])
		}
		return
	}
	var d int
	for k := 1; k <= 3; k++ {
		if i+k > len(s.s) {
			continue
		}
		old := s.cur
		fmt.Sscanf(s.s[i:i+k], "%d", &d)
		if d >= 0 && d < 256 && checkPreZero(s.s[i:i+k], d) {
			s.cur = s.cur + s.s[i:i+k]
			s.cur = s.cur + "."
			s.dfs(i+k, j+1)
			s.cur = old
		}
	}
}

func checkPreZero(s string, d int) bool {
	if s == "0" && d == 0 {
		return true
	}
	tp := []int{1, 10, 100}
	l := len(s) - 1
	return d >= tp[l]
}

func NewSolve(s string) Solve {
	so := Solve{}
	so.s = s
	so.ans = make([]string, 0)
	return so
}

func restoreIpAddresses(s string) []string {
	so := NewSolve(s)
	so.dfs(0, 0)
	return so.ans
}

func main() {
	s := "010010"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}

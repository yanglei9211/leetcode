package main

import (
	"fmt"
)

type ss struct {
	dp  [][]float64
	op1 []int
	op2 []int
	n   int
	ans float64
}

type qitem struct {
	n1, n2 int
	p      float64
}

func (s *ss) init(n int) {
	s.dp = make([][]float64, n+2)
	for i := 0; i <= n+1; i++ {
		s.dp[i] = make([]float64, n+2)
	}
	s.dp[n][n] = 1
	s.op1 = []int{100, 75, 50, 25}
	s.op2 = []int{0, 25, 50, 75}
	s.n = n
}

//func (s *ss)bfs(n int) float64 {
//	q := list.New()
//	q.PushBack(qitem{n, n, 1})
//	ct := 0
//	for q.Len() > 0 {
//		ct ++
//		if ct % 1000 == 0 {
//			fmt.Println(ct)
//		}
//		cur := q.Front()
//		citem := cur.Value.(qitem)
//		s.dp[citem.n1][citem.n2] = citem.p
//		q.Remove(cur)
//		for o := 0; o < 4; o++ {
//			nn1 := citem.n1 - s.op1[o]
//			nn2 := citem.n2 - s.op2[o]
//			np := citem.p * 0.25
//			if nn1 <= 0 && nn2 > 0 {
//				s.ans += np
//			} else if nn1 <= 0 && nn2 <= 0 {
//				s.ans += np * 0.5
//			} else if nn1 > 0 && nn2 > 0 {
//				q.PushBack(qitem{nn1, nn2, np})
//			}
//		}
//	}
//	fmt.Println(s.ans)
//	return s.ans
//}

func (s *ss) get(n1, n2 int) float64 {
	if n1 >= s.n && n2 < s.n {
		return 1
	} else if n1 >= s.n && n2 >= s.n {
		return 0.5
	} else if n2 >= s.n {
		return 0
	} else {
		if s.dp[n1][n2] > 0 {
			return s.dp[n1][n2]
		}
		cur := 0.0
		for o := 0; o < 4; o++ {
			r := s.get(n1+s.op1[o], n2+s.op2[o])
			cur += 0.25 * r
		}
		s.dp[n1][n2] = cur
		return cur
	}
}

func soupServings(n int) float64 {
	if n > 5000 {
		return 1.0
	} else {
		s := ss{}
		s.init(n)
		r := s.get(0, 0)
		return r
	}
}

func main() {
	fmt.Println(soupServings(10000))
}

package main

import "fmt"

//The set [1,2,3,…,n] contains a total of n! unique permutations.
//
//By listing and labeling all of the permutations in order,
//We get the following sequence (ie, for n = 3):
//
//"123"
//"132"
//"213"
//"231"
//"312"
//"321"
//Given n and k, return the kth permutation sequence.
//
//Note: Given n will be between 1 and 9 inclusive.

type KT struct {
	fac []int
	N   int
}

func (s *KT) init(n int) {
	s.N = n
	s.fac = make([]int, 10)
	s.fac[0] = 1
	for i := 1; i < s.N; i++ {
		s.fac[i] = s.fac[i-1] * i
	}
}

func (s *KT) kangtuo(n int, a []int) int {
	sum := 0
	for i := 0; i < n; i++ {
		t := 0
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				t++
			}
		}
		sum += t * s.fac[n-i-1]
	}
	return sum + 1
}

func (s *KT) reverse_kangtuo(n, k int) []int {
	vis := make([]bool, s.N)
	res := make([]int, n)
	var i, j int
	for i = 0; i < n; i++ {
		t := k / s.fac[n-i-1]
		for j = 1; j < n; j++ {
			if !vis[j] {
				if t == 0 {
					break
				}
				t--
			}
		}
		res[i] = j
		vis[j] = true
		k %= s.fac[n-i-1]
	}
	return res
}

func getPermutation(n int, k int) string {
	kt := KT{}
	k -= 1
	kt.init(10)
	res := kt.reverse_kangtuo(n, k)
	fmt.Println(res)
	tmp := 0
	for i := 0; i < len(res); i++ {
		tmp += res[i]
		if i < len(res)-1 {
			tmp *= 10
		}
	}
	return fmt.Sprintf("%d", tmp)
}

func main() {
	fmt.Println(getPermutation(2, 0))
}

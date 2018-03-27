package main

import "fmt"

func merge(l, r []int) ([]int, int) {
	i := 0
	j := 0
	res := make([]int, 0, len(l)+len(r))
	cnt := 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			cnt += j
			res = append(res, l[i])
			i++
		} else if l[i] > r[j] {
			cnt += len(l) - i
			res = append(res, r[j])
			j++
		}
	}
	for i < len(l) {
		cnt += len(r)
		res = append(res, l[i])
		i++
	}

	for j < len(r) {
		res = append(res, r[j])
		j++
	}
	return res, cnt >> 1
}

func count_inversion(a []int) ([]int, int) {
	if len(a) <= 1 {
		return a, 0
	}
	mid := len(a) >> 1
	left, cl := count_inversion(a[:mid])
	right, cr := count_inversion(a[mid:])
	res, cnt := merge(left, right)
	cnt += cl + cr
	return res, cnt
}

func main() {
	a := []int{9, 15, 7, 20, 3}
	fmt.Println(count_inversion(a))
}

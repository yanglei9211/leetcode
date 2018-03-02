package main

import "fmt"

func countTree(l, r int) int {
	if l >= r {
		return 1
	}
	res := 0
	for k := l; k <= r; k++ {
		lefts := countTree(l, k-1)
		rights := countTree(k+1, r)
		qaq := lefts * rights
		res += qaq
	}
	return res
}

func numTrees(n int) int {

	a := []int{1, 1, 2}
	for i := 3; i <= n; i++ {
		res := 0
		for j := i - 1; j >= 0; j-- {
			res += a[j] * a[i-j-1]
		}
		a = append(a, res)
	}
	return a[n]
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(numTrees(i))
	}
}

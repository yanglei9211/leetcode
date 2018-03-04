package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	f := [][]bool{}
	fx := []int{0, 1, 0, -1}
	fy := []int{1, 0, -1, 0}
	dir := 0
	cx := 0
	cy := 0
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	f = make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, m)
	}
	res := make([]int, 0, n*m)
	for cnt := 0; cnt < n*m; cnt++ {
		res = append(res, matrix[cx][cy])
		f[cx][cy] = true
		if n > cx+fx[dir] && cx+fx[dir] >= 0 && m > cy+fy[dir] && cy+fy[dir] >= 0 && !f[cx+fx[dir]][cy+fy[dir]] {
			cx += fx[dir]
			cy += fy[dir]
		} else {
			dir++
			dir %= len(fx)
			cx += fx[dir]
			cy += fy[dir]
		}
	}
	return res
}

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(a))
}

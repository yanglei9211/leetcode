package main

import "fmt"

func generateMatrix(n int) [][]int {
	f := [][]bool{}
	fx := []int{0, 1, 0, -1}
	fy := []int{1, 0, -1, 0}
	dir := 0
	cx := 0
	cy := 0
	res := make([][]int, n)
	f = make([][]bool, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
		f[i] = make([]bool, n)
	}
	cnt := n * n
	for i := 1; i <= cnt; i++ {
		res[cx][cy] = i
		nx := fx[dir] + cx
		ny := fy[dir] + cy
		f[cx][cy] = true
		if nx >= 0 && nx < n && ny >= 0 && ny < n && !f[nx][ny] {
			cx = nx
			cy = ny
		} else {
			dir++
			dir %= len(fx)
			cx = fx[dir] + cx
			cy = fy[dir] + cy
		}
	}
	return res
}

func main() {
	n := 5
	fmt.Println(generateMatrix(n))
}

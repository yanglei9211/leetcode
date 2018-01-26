package main

import "fmt"

//Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	if m == 0 {
		return
	}

	fx := make([]bool, n)
	fy := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				fx[i] = true
				fy[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if fx[i] || fy[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(a)
	fmt.Println(a)
}

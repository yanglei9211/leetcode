package main

import "fmt"

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(a)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := [][]int{}
	for i := 0; i < n; i++ {
		col := []int{}
		for j := n - 1; j >= 0; j-- {
			col = append(col, matrix[j][i])
		}
		tmp = append(tmp, col)
	}
	copy(matrix, tmp)
}

/*
1 2
3 4

3 1
4 2

1 2 3
4 5 6
7 8 9

7 4 1
8 5 2
9 6 3


*/

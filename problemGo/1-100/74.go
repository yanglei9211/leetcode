package main

import "fmt"

//Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//
//Integers in each row are sorted from left to right.
//The first integer of each row is greater than the last integer of the previous row.
//For example,
//
//Consider the following matrix:
//
//[
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//Given target = 3, return true.

func searchMatrix(matrix [][]int, target int) bool {
	var lx, ly, cx, cy int
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	lx = 0
	ly = n
	for lx+1 < ly {
		lmid := (lx + ly) >> 1
		if matrix[lmid][0] == target {
			return true
		} else if matrix[lmid][0] > target {
			ly = lmid
		} else {
			lx = lmid
		}
	}
	cx = 0
	cy = m
	for cx+1 < cy {
		cmid := (cx + cy) >> 1
		if matrix[lx][cmid] == target {
			return true
		} else if matrix[lx][cmid] > target {
			cy = cmid
		} else {
			cx = cmid
		}
	}
	if matrix[lx][cx] == target {
		return true
	}
	return false
}

func main() {
	a := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	res := searchMatrix(a, 12)
	fmt.Println(res)
}

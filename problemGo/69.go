package main

import "fmt"

//Implement int sqrt(int x).
//
//Compute and return the square root of x.
//
//x is guaranteed to be a non-negative integer.

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l := 1
	r := x
	for l+1 < r {
		mid := (l + r) >> 1
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}

func main() {
	fmt.Println(mySqrt(145))
}

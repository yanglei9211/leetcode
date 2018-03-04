package main

import "fmt"

//The gray code is a binary numeral system where two successive values differ in only one bit.
//
//Given a non-negative integer n representing the total number of bits in the code, print the
// sequence of gray code. A gray code sequence must begin with 0.
//
//For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:
//
//00 - 0
//01 - 1
//11 - 3
//10 - 2
//Note:
//For a given n, a gray code sequence is not uniquely defined.
//
//For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.
//
//For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.

func main() {
	res := grayCode(4)
	fmt.Println(res)
}

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	} else if n == 2 {
		return []int{0, 1, 3, 2}
	} else {
		b := []int{0, 1, 3, 2}
		res := make([]int, len(b))
		copy(res, b)
		fmt.Println(res)
		for i := 2; i < n; i++ {
			for j := len(b) - 1; j >= 0; j-- {
				res = append(res, b[j]+(1<<uint(i)))
			}
			b = make([]int, len(res))
			copy(b, res)
		}
		return res
	}
}

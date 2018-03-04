package main

import "fmt"

//Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
//
//You may assume the integer do not contain any leading zero, except the number 0 itself.
//
//The digits are stored such that the most significant digit is at the head of the list.

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1]++
	for i := n - 1; i >= 0; i-- {
		if i > 0 && digits[i] > 9 {
			digits[i] -= 10
			digits[i-1]++
		}
	}
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	a := []int{9, 9, 9}
	res := plusOne(a)
	fmt.Println(res)
}

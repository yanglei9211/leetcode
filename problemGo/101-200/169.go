package main

import "fmt"

func majorityElement(nums []int) int {
	var m1, c1 int
	for _, x := range nums {
		if c1 == 0 {
			m1 = x
			c1 = 1
		} else {
			if x == m1 {
				c1++
			} else {
				c1--
			}
		}
	}
	cc1 := 0
	for _, x := range nums {
		if x == m1 {
			cc1++
		}
	}
	return m1
}

func main() {
	a := []int{1, 1, 2, 1, 1, 1, 2, 1}
	fmt.Println(majorityElement(a))
}

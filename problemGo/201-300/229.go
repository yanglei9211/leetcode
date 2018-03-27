package main

import "fmt"

func majorityElement(nums []int) []int {
	var m1, m2, c1, c2 int
	for _, x := range nums {
		if c1 == 0 && x != m2 {
			m1 = x
			c1 = 0
		} else if c2 == 0 {
			m2 = x
			c2 = 0
		}
		fmt.Println("**********")
		fmt.Println(x)
		if m1 == x {
			c1++
		} else if m2 == x {
			c2++
		} else {
			c1--
			c2--
		}
		fmt.Println(m1, c1)
		fmt.Println(m2, c2)
		fmt.Println("-----")
	}
	var cc1, cc2 int
	line := len(nums) / 3
	for _, x := range nums {
		if x == m1 {
			cc1++
		}
		if x == m2 {
			cc2++
		}
	}
	ans := []int{}
	fmt.Println(m1, m2)
	if cc1 > line {
		ans = append(ans, m1)
	}
	if cc2 > line && m2 != m1 {
		ans = append(ans, m2)
	}
	return ans
}

func main() {
	a := []int{1, 2, 2, 3, 2, 1, 1, 3}
	fmt.Println(majorityElement(a))
}

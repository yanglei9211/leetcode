package main

import (
	"fmt"
)

//Given two binary strings, return their sum (also a binary string).
//
//For example,
//a = "11"
//b = "1"
//Return "100".

func addBinary(a string, b string) string {
	formatNum := func(s string) []byte {
		l := len(s)
		res := make([]byte, l)
		for i := 0; i < l; i++ {
			res[l-1-i] = byte(s[i]) - byte('0')
		}
		return res
	}
	maxx := func(x, y int) int {
		if x < y {
			return y
		} else {
			return x
		}
	}

	minn := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	num1 := formatNum(a)
	num2 := formatNum(b)
	l1 := len(num1)
	l2 := len(num2)
	if l1 < l2 {
		l1, l2 = l2, l1
		num1, num2 = num2, num1
	}
	l := minn(l1, l2)
	lm := maxx(l1, l2)

	for i := 0; i < l; i++ {
		num1[i] += num2[i]
	}
	for i := 0; i < lm-1; i++ {
		if num1[i] > 1 {
			num1[i+1]++
			num1[i] -= 2
		}
	}
	fmt.Println(num1)
	if num1[lm-1] > 1 {
		num1[lm-1] -= 2
		num1 = append(num1, 1)
		lm++
	}

	temp := make([]byte, lm)
	for i := 0; i < lm; i++ {
		temp[i] = num1[lm-1-i] + byte('0')
	}
	return string(temp)
}

func main() {
	s := "11"
	s2 := "1"
	fmt.Println(len(s2), len(s))
	////s1 := "101"
	//for _, i := range s {
	//	fmt.Println(i, reflect.TypeOf(i))
	//}
	res := addBinary(s, s2)
	fmt.Println(res)
}

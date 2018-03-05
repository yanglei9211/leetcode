package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "0"
	b := "0"
	fmt.Println(multiply(a, b))
}

func multiply(num1 string, num2 string) string {
	x := string2arr(num1)
	y := string2arr(num2)
	xMy := make([]int, len(x)+len(y)+1)
	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			xMy[i+j-1] += x[i] * y[j]
		}
	}
	for i := 1; i < len(xMy)-1; i++ {
		if xMy[i] > 9 {
			xMy[i+1] += xMy[i] / 10
			xMy[i] %= 10
		}
	}
	xMy = reverse(xMy[1:])
	res := ""
	i := 0
	for i < len(xMy) && xMy[i] == 0 {
		i++
	}
	for i < len(xMy) {
		res += strconv.Itoa(xMy[i])
		i++
	}
	if len(res) == 0 {
		res = "0"
	}
	return res
}

func string2arr(s string) []int {
	// 0 : 48
	res := []int{}
	for i := 0; i < len(s); i++ {
		res = append(res, int(s[i])-48)
	}
	res = append(res, 0)
	res = reverse(res)
	return res
}

func reverse(x []int) []int {
	l := 0
	r := len(x) - 1
	mid := r >> 1
	for i := 0; i <= mid; i++ {
		if l+i < r-i {
			swap(&x[l+i], &x[r-i])
		}
	}
	return x
}

func swap(x, y *int) {
	*x ^= *y
	*y ^= *x
	*x ^= *y
}

package main

import "fmt"

func main() {
	s1 := "ababaab"
	s2 := ""
	res := strStr(s1, s2)
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if n == 0 && m > 0 {
		return -1
	}
	if n == 0 || m == 0 {
		return 0
	}
	f := getFail(needle)
	//fmt.Println(f)
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = f[j]
		}
		if needle[j] == haystack[i] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func getFail(needle string) []int {
	//m = len([]rune(needle))
	var f = []int{}
	for i := 0; i < len(needle)+2; i++ {
		f = append(f, 0)
	}
	f[0] = 0
	f[1] = 0
	for i := 1; i < len(needle); i++ {
		j := f[i]
		for j > 0 && needle[i] != needle[j] {
			j = f[j]
		}
		if needle[i] == needle[j] {
			f[i+1] = j + 1
		} else {
			f[i+1] = 0
		}
	}
	return f
}

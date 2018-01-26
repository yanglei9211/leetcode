package main

import (
	"fmt"
)

//Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
//For example,
//S = "ADOBECODEBANC"
//T = "ABC"
//Minimum window is "BANC".
//
//Note:
//If there is no such window in S that covers all characters in T, return the empty string "".
//
//If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.

func mapCopy(s, t map[byte]int) {
	for k, v := range s {
		t[k] = v
	}
}

func minWindow(s string, t string) string {
	check := func(ct []int) bool {
		for i := 0; i < 256; i++ {
			if ct[i] > 0 {
				return false
			}
		}
		return true
	}
	if s == "" || t == "" {
		return ""
	}
	ct := make([]int, 256)
	for idx, _ := range t {
		ct[t[idx]]++
	}

	lx := 0
	ly := 0
	res := ""
	for lx <= ly && lx < len(s) && ly <= len(s) {
		if check(ct) {
			if (res != "" && len(res) > ly-lx) || (res == "") {
				res = s[lx:ly]
			}
			ct[s[lx]]++
			lx++
		} else if ly < len(s) {
			ct[s[ly]]--
			ly++
		} else if ly == len(s) {
			break
		}
	}
	return res
}

func main() {
	s := "abcabc"
	t := "ca"
	fmt.Println(len(s))
	res := minWindow(s, t)
	fmt.Println(res)
	//ss := "ABabab"
	//a := make([]int, 256)
	//for idx, _ := range(ss) {
	//	c := ss[idx]
	//	a[c]++
	//}
	//fmt.Println(a)
}

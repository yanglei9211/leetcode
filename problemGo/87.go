package main

import "fmt"

func hasCheck(s1, s2 string) bool {
	ct1 := make(map[rune]int)
	ct2 := make(map[rune]int)
	for _, x := range s1 {
		ct1[x]++
	}
	for _, x := range s2 {
		ct2[x]++
	}
	for k := range ct1 {
		if ct1[k] != ct2[k] {
			return false
		}
	}
	return true
}

func simstr(s1, s2 string) bool {
	if !hasCheck(s1, s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	if len(s1) == len(s2) {
		if len(s1) == 1 {
			return s1[0] == s2[0]
		}
		for cut := 1; cut < len(s1); cut++ {
			if simstr(s1[:cut], s2[:cut]) && simstr(s1[cut:], s2[cut:]) {
				return true
			}
			fcut := len(s1) - cut
			if simstr(s1[:cut], s2[fcut:]) && simstr(s1[cut:], s2[:fcut]) {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

func isScramble(s1 string, s2 string) bool {
	return simstr(s1, s2)
}

func main() {
	s1 := "bab"
	s2 := "bab"
	fmt.Println(isScramble(s1, s2))
	//fmt.Println(hasCheck(s1, s2))
}

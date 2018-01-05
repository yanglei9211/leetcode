package main

import (
	"strings"
	"fmt"
)
//Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
// return the length of last word in the string.
//
//If the last word does not exist, return 0.
//
//Note: A word is defined as a character sequence consists of non-space characters only.
//
//Example:
//
//Input: "Hello World"
//Output: 5

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
    ss :=  strings.Split(s, " ")
	for i := len(ss)-1; i >=0; i-- {
		if len(ss[i]) > 0 {
			return len(ss[i])
		}
	}
	return 0
}

func main() {
	s := "aadasf gvadsfga fads adb   "
	fmt.Println(lengthOfLastWord(s))
}
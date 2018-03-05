package main

import (
	"fmt"
	"regexp"
)

//Validate if a given string is numeric.
//
//Some examples:
//"0" => true
//" 0.1 " => true
//"abc" => false
//"1 a" => false
//"2e10" => true
//Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one.
//
//Update (2015-02-10):
//The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.

func isNumber(s string) bool {
	reg1 := regexp.MustCompile(`^\s*[-|+]?\d+\s*$`)
	reg2 := regexp.MustCompile(`^\s*[-|+]?\d*\.\d+\s*$`)
	reg3 := regexp.MustCompile(`^\s*[-|+]?\d+\.\d*\s*$`)
	reg4 := regexp.MustCompile(`^\s*[-|+]?\d+e[-|+]?\d+\s*$`)
	reg5 := regexp.MustCompile(`^\s*[-|+]?\d*\.\d+e[-|+]?\d+\s*$`)
	reg6 := regexp.MustCompile(`^\s*[-|+]?\d+\.\d*e[-|+]?\d+\s*$`)
	if reg1.MatchString(s) {
		fmt.Println(1)
		return true
	}
	if reg2.MatchString(s) {
		fmt.Println(2)
		return true
	}
	if reg3.MatchString(s) {
		fmt.Println(3)
		return true
	}
	if reg4.MatchString(s) {
		fmt.Println(4)
		return true
	}
	if reg5.MatchString(s) {
		fmt.Println(5)
		return true
	}
	if reg6.MatchString(s) {
		fmt.Println(6)
		return true
	}
	return false
}

func main() {
	s := []string{"e10"}
	for _, r := range s {
		fmt.Println(isNumber(r))
	}
}

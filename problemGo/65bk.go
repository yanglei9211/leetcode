package main

import (
	"fmt"
	"strconv"
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
	//defer func() {
	//	if err := recover(); err != nil {
	//		return false
	//	}
	//}()
	_, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return true
	} else {
		return false
	}

}

func main() {
	s := []string{"e10"}
	for _, r := range s {
		fmt.Println(isNumber(r))
	}
}

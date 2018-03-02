package main

import "fmt"

//A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//Given an encoded message containing digits, determine the total number of ways to decode it.
//
//For example,
//Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).
//
//The number of ways decoding "12" is 2.

func numDecodings(s string) int {

	if len(s) == 0 {
		return 0
	}
	//dp := make([]int, len(s))
	//dp[0] = 1
	//var d int
	//fmt.Sscanf(s[0:2], "%d", &d)
	//if d < 27 {
	//	dp[1] = 2
	//} else {
	//	dp[1] = 1
	//}
	//for i := 2; i < len(s); i ++ {
	//	dp[i] = dp[i-1]
	//	fmt.Sscanf(s[i-1:i+1], "%d", &d)
	//	if d < 27 {
	//		dp[i] += dp[i-2]
	//	}
	//}
	//return dp[len(s)-1]
	dp := make([]int, len(s)+1)
	dp[0] = 1
	var d int
	for i := 1; i <= len(s); i++ {
		if i-1 >= 0 {
			if s[i-1] != '0' {
				dp[i] += dp[i-1]
			}
		}
		if i-2 >= 0 {
			fmt.Sscanf(s[i-2:i], "%d", &d)
			if 9 < d && d < 27 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(s)]
	//if s[len(s)-1] == '0' {
	//	if len(s) > 2 && s[len(s)-2] == '0' {
	//		return 0
	//	} else {
	//		return dp[len(s)-2]
	//	}
	//} else {
	//	return dp[len(s)]
	//}
}

func main() {
	fmt.Println(numDecodings("110"))

}

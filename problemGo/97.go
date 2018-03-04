package main

import "fmt"

//Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
//
//For example,
//Given:
//s1 = "aabcc",
//s2 = "dbbca",
//
//When s3 = "aadbbcbcac", return true.
//When s3 = "aadbbbaccc", return false.

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	n := len(s3)
	dp := make([][]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	//if s1[0] == s3[0] {
	//	dp[1][1] = 1
	//}
	//if s2[0] == s3[0] {
	//	dp[1][0] == 1
	//}
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j > 0 && j-1 < len(s1) && s3[i-1] == s1[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}

			if i-j > 0 && i-j-1 < len(s2) && s3[i-1] == s2[i-j-1] {
				dp[i][j] += dp[i-1][j]
			}
		}
	}
	for j := 0; j <= n; j++ {
		if dp[n][j] > 0 {
			return true
		}
	}
	return false
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}

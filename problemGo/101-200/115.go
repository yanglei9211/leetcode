package main

import "fmt"

func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
    dp := make([][]int, len(s)+5)
	for i := 0; i < len(s)+5; i++ {
		dp[i] = make([]int, len(t)+5)
		dp[i][0] = 1
	}
	dp[0][0] = 1

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i && j <= len(t); j++ {
			if i > j {
				fmt.Println(i,j)
				dp[i][j] += dp[i-1][j]
			}
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	for i := 0; i <= len(s); i++ {
		fmt.Println(dp[i])
	}
	return dp[len(s)][len(t)]
}


func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
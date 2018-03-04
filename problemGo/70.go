package main

import "fmt"

//You are climbing a stair case. It takes n steps to reach to the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//Note: Given n will be a positive integer.

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] += dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(10))
}

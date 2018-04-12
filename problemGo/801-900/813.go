package main

import (
	"fmt"
)

func largestSumOfAverages(a []int, k int) float64 {
	ave := func(x, y int) float64 {
		var r float64 = 0
		for i := x; i <= y; i++ {
			r += float64(a[i])
		}
		return r / float64(y-x+1)
	}
	n := len(a)
	dp := make([][]float64, len(a)+2)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]float64, k+2)
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = ave(0, i-1)
	}
	for j := 2; j <= k; j++ {
		for i := j; i <= n; i++ {
			var cur float64 = 0
			for p := j - 1; p < i; p++ {
				s := dp[p][j-1] + ave(p, i-1)
				if s > cur {
					cur = s
				}
			}
			dp[i][j] = cur
		}
	}
	//for i := 1; i <= n; i ++ {
	//	fmt.Println(dp[i])
	//}
	return dp[n][k]

}

func main() {
	a := []int{4, 1, 7, 5, 6, 2, 3}
	fmt.Println(largestSumOfAverages(a, 4))
}

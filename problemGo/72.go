package main

import "fmt"

//Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)
//
//You have the following 3 operations permitted on a word:
//
//a) Insert a character
//b) Delete a character
//c) Replace a character

func minDistance(word1 string, word2 string) int {
	minn := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	n := len(word1)
	m := len(word2)
	if n == 0 && m == 0 {
		return 0
	} else if n == 0 {
		return m
	} else if m == 0 {
		return n
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	var t int
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i > 0 && j > 0 {
				dp[i][j] = minn(dp[i-1][j]+1, dp[i][j-1]+1)
				if word1[i-1] == word2[j-1] {
					t = 0
				} else {
					t = 1
				}

				dp[i][j] = minn(dp[i][j], dp[i-1][j-1]+t)
			} else if i > 0 {
				dp[i][j] = dp[i-1][j] + 1
			} else if j > 0 {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		fmt.Printf("%d ", dp[i][j])
	//	}
	//	fmt.Println()
	//}
	return dp[n][m]

}

func main() {
	word1 := "ababa"
	word2 := "babab"
	res := minDistance(word1, word2)
	fmt.Println(res)
}

package main

import "fmt"

func splitArraySameAverage(a []int) bool {
	n := len(a)
	if n == 0 {
		return true
	}
	maxSize := 10000 * (n + 1)

	dp := make([][][]bool, n+2)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]bool, n+2)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]bool, maxSize+2)
		}
	}

	sum := 0
	for _, x := range a {
		sum += x
	}
	for i := 0; i <= n; i++ {
		dp[i][0][0] = true
	}
	//dp[0][0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for k := a[i-1]; k <= maxSize; k++ {
				if dp[i-1][j][k] {
					dp[i][j][k] = true
				}

				if dp[i-1][j][k-a[i-1]] {
					//fmt.Println(i, j+1, k)
					dp[i][j+1][k] = true
					if i < n {
						dp[i+1][j+1][k] = true
					}
					if (k*n == (j+1)*sum) && j+1 < n {
						//fmt.Println(j, k-a[i-1])
						//fmt.Println(j+1, k)
						//fmt.Println(n, sum)
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	//a := []int{3479,5179,5278,3264,6947,7122,5273,9586,4349,3631,5715,5632,7137,3705,4761,4569,5337,1968,2}
	a := []int{3, 1}
	fmt.Println(splitArraySameAverage(a))
}

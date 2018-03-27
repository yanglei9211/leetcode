package main

import "fmt"

func splitArraySameAverage(a []int) bool {
	n := len(a)
	if n == 0 {
		return true
	}
	maxSize := 10000 * (n + 1)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, maxSize+5)
	}
	sum := 0
	for _, x := range a {
		sum += x
	}

	dp[0][0] = true
	upSize := maxSize >> 1
	for i := 1; i <= n; i++ {
		//for j := 0; j < i; j++ {
		up := i - 1
		if up > (n >> 1) {
			up = n >> 1
		}
		for j := up; j >= 0; j-- {
			for k := upSize; k >= a[i-1]; k-- {
				if dp[j][k-a[i-1]] {
					dp[j+1][k] = true
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
	a := []int{3479, 5179, 5278, 3264, 6947, 7122, 5273, 9586, 4349, 3631, 5715, 5632, 7137, 3705, 4761, 4569, 5337, 1968, 2}
	fmt.Println(splitArraySameAverage(a))
}

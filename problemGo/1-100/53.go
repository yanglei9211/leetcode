package main

import "fmt"

func maxSubArray(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	n := len(nums)
	dp := make([]int, n)
	res := nums[0]
	if nums[0] > 0 {
		dp[0] = nums[0]
	}
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		res = max(res, nums[i])
	}
	if res < 0 {
		return res
	}
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func main() {
	//[-2,1,-3,4,-1,2,1,-5,4]
	//	s := []int{-2,1,-3,4,-1,2,1,-5,4}
	s := []int{-1, -2}
	fmt.Println(maxSubArray(s))
}

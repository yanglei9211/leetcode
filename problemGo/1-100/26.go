package main

import "fmt"

func main() {
	var nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	var res = []int{}
	var ss = 1
	if len(nums) > 0 {
		res = append(res, nums[0])
	} else {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		if nums[i] != nums[j] {
			res = append(res, nums[j])
			ss++
		}
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
	//fmt.Println(nums)
	return ss
}

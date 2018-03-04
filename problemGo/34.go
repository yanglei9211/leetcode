package main

import "fmt"

func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	x := 8
	// fmt.Println(searchRange(a, 1))
	fmt.Println(leftSearch(a, x))
	fmt.Println(rightSearch(a, x))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return nosol()
	}
	x := leftSearch(nums, target)
	if x == -1 {
		return nosol()
	} else {
		y := rightSearch(nums, target)
		return []int{x, y}
	}
}

func leftSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var res = -1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			res = mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func rightSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var res = -1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			res = mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func nosol() []int {
	return []int{-1, -1}
}

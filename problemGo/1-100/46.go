package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	a := sort.IntSlice(nums)
	a.Sort()
	x := jc(len(nums))
	tmp := make([]int, len(nums))
	copy(tmp, a)
	ans = append(ans, tmp)
	for i := 0; i < x-1; i++ {
		nextPermutation(a)
		tp := make([]int, len(nums))
		copy(tp, a)
		ans = append(ans, tp)
	}
	return ans
}

func jc(x int) int {
	if x == 1 {
		return 1
	}
	return x * jc(x-1)
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	if len(nums) == 2 {
		swap(nums, 0, 1)
		return
	}

	index := len(nums) - 2
	for index >= 0 {
		if nums[index] < nums[index+1] {
			break
		}
		index--
	}

	if index < 0 {
		reverse(nums, 0, len(nums)-1)
	} else {
		for i := len(nums) - 1; i > index; i-- {
			if nums[index] < nums[i] {
				swap(nums, index, i)
				break
			}
		}
		reverse(nums, index+1, len(nums)-1)
	}
}

func swap(nums []int, x int, y int) {
	nums[x] ^= nums[y]
	nums[y] ^= nums[x]
	nums[x] ^= nums[y]
}

func reverse(nums []int, start int, end int) {
	mid := (start + end) >> 1
	for i := 0; i <= mid; i++ {
		if start+i < end-i {
			swap(nums, start+i, end-i)
		}
	}
}

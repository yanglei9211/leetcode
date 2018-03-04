package main

import (
	"fmt"
	"sort"
)

//Follow up for "Remove Duplicates":
//What if duplicates are allowed at most twice?
//
//For example,
//Given sorted array nums = [1,1,1,2,2,3],
//
//Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3.
//It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	var i, j int
	cur := 0
	for j < n {
		if j == 0 || cur < 2 {
			if j > 0 && nums[j] == nums[j-1] {
				cur++
			} else {
				cur = 1
			}
			nums[i] = nums[j]
			i++
			j++
		} else {

			if cur == 2 {
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				if j < n {
					cur = 0
				} else {
					break
				}
			}
		}
	}
	return i
}
func main() {
	a := []int{1, 2, 3, 1, 2, 1}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
}

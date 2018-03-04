package main

import "fmt"

//Given an array with n objects colored red, white or blue, sort them so that objects of the same
// color are adjacent, with the colors in the order red, white and blue.
//
//Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
//
//Note:
//You are not suppose to use the library's sort function for this problem.
//
//click to show follow up.
//
//Follow up:
//A rather straight forward solution is a two-pass algorithm using counting sort.
//First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total
// number of 0's, then 1's and followed by 2's.
//
//Could you come up with an one-pass algorithm using only constant space?

func sortColors(nums []int) {
	count := make([]int, 3)
	for _, x := range nums {
		count[x]++
	}
	cnt := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < count[i]; j++ {
			nums[cnt] = i
			cnt++
		}
	}
}

func main() {
	a := []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 0, 1, 1, 2, 2}
	sortColors(a)
	fmt.Println(a)
}

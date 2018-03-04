package main

import "fmt"

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//Note:
//You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold
//additional elements from nums2. The number of elements initialized in nums1 and nums2 are
// m and n respectively.

func merge(nums1 []int, m int, nums2 []int, n int) {
	tm := make([]int, 0, m)
	for i := 0; i < m; i++ {
		tm = append(tm, nums1[i])
	}
	i := 0
	j := 0
	idx := 0
	for i < m && j < n {
		if tm[i] < nums2[j] {
			nums1[idx] = tm[i]
			i++
		} else {
			nums1[idx] = nums2[j]
			j++
		}
		idx++
		fmt.Println(nums1)
	}
	for i < m {
		nums1[idx] = tm[i]
		i++
		idx++
	}
	for j < n {
		nums1[idx] = nums2[j]
		j++
		idx++
	}
}

func main() {
	a := []int{1, 5, 9, 12, 19}
	b := []int{3, 8, 20}
	nums1 := make([]int, len(a)+len(b))
	for i, x := range a {
		nums1[i] = x
	}
	merge(nums1, 5, b, 3)
	fmt.Println(nums1)
}

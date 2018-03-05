package main

import (
	"math"
	"fmt"
)

func minn(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxx(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	n := len(nums1)
	m := len(nums2)
	l := 0
	r := 2 * n
	var l1, l2, r1, r2 int
	for l <= r {
		c1 := (l+r) >> 1
		c2 := n + m - c1
		if c1 == 0 {
			l1 = math.MinInt32
		} else {
			l1 = nums1[(c1-1)>>1]
		}
		if c1 == 2 * n {
			r1 = math.MaxInt32
		} else {
			r1 = nums1[c1>>1]
		}

		if c2 == 0 {
			l2 = math.MinInt32
		} else {
			l2 = nums2[(c2-1)>>1]
		}

		if c2 == 2 * m {
			r2 = math.MaxInt32
		} else {
			r2 = nums2[c2>>1]
		}
		if l1 > r2 {
			r = c1-1
		} else if l2 > r1 {
			l = c1+1
		} else {
			break
		}
	}
	sum := minn(r1,r2) + maxx(l1, l2)
	return float64(sum) / 2.0
}

func main() {
	a := []int{1,5,11,16,19}
	b := []int{3,7,28}
	fmt.Println(findMedianSortedArrays(a,b))
}
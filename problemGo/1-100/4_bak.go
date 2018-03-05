package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	//b := []int{6, 7, 8, 9, 10}
	b := []int{4, 5}
	//fmt.Println(findMedianSortedArrays(a, b))
	fmt.Println(findCnt(a,b,4))
	//fmt.Println(findCnt(a,b,1))
}

func cntLess(a []int, x int) int {
	l := 0
	r := len(a)
	for l+1 < r {
		mid := (l + r) >> 1
		if a[mid] >= x {
			r = mid
		} else {
			l = mid
		}
	}
	if l >= len(a) {
		return len(a)
	}
	if a[l] >= x {
		return l
	} else {
		return l + 1
	}

}

func findCntOnePlus(nums []int, x int, b int) int {
	c := cntLess(nums, b)
	if c == x {
		return b
	}
	l := 0
	r := len(nums)
	for l+1 < r {
		mid := (l + r) >> 1
		c = cntLess(nums, nums[mid])
		if nums[mid] > b {
			c++
		}
		if c > x {
			r = mid
		} else {
			l = mid
		}
	}
	return nums[l]
}

func findCntOne(nums []int, x int) int {
	l := 0
	r := len(nums)
	for l+1 < r {
		mid := (l + r) >> 1
		c := cntLess(nums, nums[mid])
		if c > x {
			r = mid
		} else {
			l = mid
		}
	}
	return nums[l]
}

func findCnt(nums1, nums2 []int, x int) int {

	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 && n2 == 0 {
		return 0
	}

	if n1 > 1 && n2 > 1 {
		l1 := 0
		r1 := n1
		l2 := 0
		r2 := n2
		for l1+1 < r1 && l2+1 < r2 {
			mid1 := (l1 + r1) >> 1
			mid2 := (l2 + r2) >> 1
			cnt11 := cntLess(nums1, nums1[mid1])
			cnt12 := cntLess(nums2, nums1[mid1])
			cnt22 := cntLess(nums2, nums2[mid2])
			cnt21 := cntLess(nums1, nums2[mid2])
			if cnt11+cnt12 == x {
				return nums1[mid1]
			}
			if cnt21+cnt22 == x {
				return nums2[mid2]
			}
			if cnt11+cnt12 > x {
				r1 = mid1
			} else {
				l1 = mid1
			}

			if cnt21+cnt22 > x {
				r2 = mid2
			} else {
				l2 = mid2
			}
		}

		var tp0, sub0, tp1, sub1 int
		sub1 = 0x3f3f3f3f
		sub0 = 0x3f3f3f3f
		if l1+1 == r1 {
			d := nums1[l1]
			c1 := cntLess(nums1, d)
			c2 := cntLess(nums2, d)
			if c1+c2 <= x {
				//return d
				tp0 = d
				sub0 = x - c1 - c2
			}
		}

		if l2+1 == r2 {
			d := nums2[l2]
			c1 := cntLess(nums1, d)
			c2 := cntLess(nums2, d)
			if c1+c2 <= x {
				tp1 = d
				sub1 = x - c1 - c2
			}
		}

		if sub0 == 0 || sub0 <= sub1 {
			return tp0
		}
		if sub1 == 0 || sub1 <= sub0 {
			return tp1
		}
	}

	if n1 == 0 {
		if n2 == 1 {
			return nums2[0]
		} else if n2 > 1 {
			r := findCntOne(nums2, x)
			return r
		}
	} else if n1 == 1 {
		if n2 > 1 {
			r := findCntOnePlus(nums2, x, nums1[0])
			return r
		}
	}

	if n2 == 0 {
		if n1 == 1 {
			return nums1[0]
		} else if n1 > 1 {
			r := findCntOne(nums1, x)
			return r
		}
	} else if n2 == 1 {
		if n1 > 1 {
			r := findCntOnePlus(nums1, x, nums2[0])
			return r
		}
	}
	return -1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n == 0 {
		return 0
	}
	v := n >> 1
	if (n & 1) == 1 {
		x := findCnt(nums1, nums2, v)
		return float64(x)
	} else {
		if len(nums1) == 1 && len(nums2) == 1 {
			z := float64(nums1[0]) + float64(nums2[0])
			z = z / 2.0
			return z
		}
		x := findCnt(nums1, nums2, v-1)
		y := findCnt(nums1, nums2, v)
		z := float64(x) + float64(y)
		z = z / 2.0
		return z
	}
	return 0.0
}

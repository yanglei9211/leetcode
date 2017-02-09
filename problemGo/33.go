package main

import "fmt"

func main() {
    x := []int{3,1}
    t := 1
    fmt.Println(search(x, t))
}

func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    l := 0
    r := len(nums)-1
    for l < r {
        mid := (l+r)>>1
        if nums[mid] == target {
            return mid
        } else {
            if nums[l] <= nums[mid] {
                if target >= nums[l] && target < nums[mid] {
                    r = mid-1
                } else {
                    l = mid+1
                }
            } else {
                if target > nums[mid] && target <= nums[r] {
                    l = mid+1
                } else {
                    r = mid-1
                }
            }
        }
    }
    if target == nums[l]{
        return l
    } else {
        return -1
    }
}

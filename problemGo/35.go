package main

import "fmt"

func main() {
    x := []int{1,3,4,6}
    y := 0
    fmt.Println(searchInsert(x,y))
}

func searchInsert(nums []int, target int) int {
    l := 0
    r := len(nums)-1
    for l < r {
        mid := (l+r)>>1
        if nums[mid] == target{
            return mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    if nums[l] < target {
        return l + 1
    } else {
        return l
    }
}
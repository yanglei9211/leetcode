package main

import "fmt"

func main(){
    var nums = []int{3,2,2,3}
    x := removeElement(nums, 3)
    fmt.Println(x)
    fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
    var res = []int{}
    begin := 0
    end := 0
    for i, x := range nums {
        if x == val{
            if begin < end {
                res = append(res, nums[begin:end]...)
            }
            begin = i + 1
            end = begin
        } else {
            end++
        }
    }
    if begin < end {
        res = append(res, nums[begin:end]...)
    }
    for i := 0; i < len(res); i++ {
        nums[i] = res[i]
    }
    return len(res)
}

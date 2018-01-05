package main

import "fmt"
//A = [2,3,1,1,4], return true
//A = [3,2,1,0,4], return false


func canJump(nums []int) bool {
    minn := len(nums)-1
	for i := minn-1; i >=0; i-- {
		if nums[i] + i >= minn {
			minn = i
		}
	}
	fmt.Println(minn)
	return minn == 0
}

func main() {
	a := []int{3,2,1,0,4}
	fmt.Println(canJump(a))
}
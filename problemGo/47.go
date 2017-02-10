package main

import (
    "fmt"
)

func main() {
    a := []int{-1,2,0,-1,1,0,1}
    fmt.Println(permuteUnique(a))
}

func permuteUnique(nums []int) [][]int {
    mp := make(map[int]int)
    for _, x := range nums {
        mp[x]++
    }
    ans := [][]int{}
    arr := []int{}
    dfs(len(nums), arr, mp, &ans)
    return ans
}

func dfs(n int, arr []int, mp map[int]int, ans *[][]int){
    if n==0 {
        tmp := make([]int, len(arr))
        copy(tmp, arr)
        *ans = append(*ans, tmp)
        return
    }
    for key, value := range mp {
        if value > 0 {
            mp[key]--
            arr = append(arr, key)
            dfs(n-1, arr, mp, ans)
            arr = arr[:len(arr)-1]
            mp[key]++
        }
    }

}


func dupAns(arr []int, ans [][]int)bool {
    for i := 0; i < len(ans); i++ {
        tmp := make([]int, len(arr))
        copy(tmp, ans[i])
        if same(arr, tmp){
            return true
        }
    }
    return false
}

func same(x []int, y []int)bool {
    if len(x) != len(y) {
        return false
    }
    for i := 0; i < len(x); i++ {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}
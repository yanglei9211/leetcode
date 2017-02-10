package main

import (
    "fmt"
    "sort"
)

func main() {
    x := []int{4,4,2,1,4,2,2,1,3}
    a := sort.IntSlice(x)
    a.Sort()
    fmt.Println(a)
    fmt.Println(x)
    y := 6
    fmt.Println(combinationSum2(x, y))
}

func combinationSum2(candidates []int, target int) [][]int {

    ans := [][]int{}
    arr :=[]int{}
    dfs(target, -1, candidates, arr, &ans)
    return ans
}

func dfs(res, last int, data []int, arr []int, ans *[][]int){
    if res == 0 {
        var tmp = []int{}
        //fmt.Println(arr)
        for _, x := range arr{
            tmp = append(tmp, x)
        }
        if !dupAns(tmp, *ans){
            *ans = append(*ans, tmp)
        }
        return
    }

    for i := last+1; i < len(data); i++ {
        if res >= data[i]{
            arr = append(arr, data[i])
            dfs(res-data[i], i, data, arr, ans)
            arr = arr[:len(arr)-1]
        }
    }
    return
}

func dupAns(arr []int, ans [][]int)bool {
    for i := 0; i < len(ans); i++ {
        if same(arr, ans[i]){
            return true
        }
    }
    return false
}

func same(x []int, y []int)bool {
    if len(x) != len(y) {
        return false
    }
    a := sort.IntSlice(x)
    b := sort.IntSlice(y)
    a.Sort()
    b.Sort()
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
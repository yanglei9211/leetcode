package main

import "fmt"

func main() {
    a := [][]int{}
    tt(&a)
    fmt.Println(a)
}

func tt(ans *[][]int){
    *ans = append(*ans, []int{1,2,3})
    *ans = append(*ans, []int{3,4,5})
}
package main

import "fmt"

func main() {
    fmt.Println(divide(-2147483648,-1))
}

func divide(x int, y int) int {
    int_max := (1<<31) - 1
    int_min := -(1<<31)
    xx := x
    yy := y
    x = abs(x)
    y = abs(y)
    res := 0
    for x >= y {
        sub := y
        add := 1
        for x >= sub<<1 {
            sub <<= 1
            add <<= 1
        }
        x -= sub
        res += add
    }
    if xx ^ yy < 0 {
        res = -res
    }
    if res > int_max {
        res = int_max
    }
    if res < int_min {
        res = int_min
    }
    return res
}

func abs(x int)int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}

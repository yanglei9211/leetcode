package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(countAndSay(5))
}


func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    s := "1"
    for i := 1; i < n; i++ {
        s = genNext(s)
    }
    return s
}

func genNext(s string) string {
    bs := s[0]
    cnt := 1
    res := ""
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1]{
            cnt++
        } else {
            res += strconv.Itoa(cnt) + string(bs)
            bs = s[i]
            cnt = 1
        }
    }
    res += strconv.Itoa(cnt) + string(bs)
    return res
}

/*
1
11
21
1211
111221
312211
13112221
1113213211
*/
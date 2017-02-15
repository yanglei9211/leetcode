package main

import "fmt"

func main() {
    s := "aab"
    p := "c*a*b"
    var save byte
    save = p[0]
    for i := 0; i < len(p); i++ {
        if sa

    }
    fmt.Println(save)
    fmt.Println(s, p)
}

func isMatch(s string, p string) bool {
    return false
}

func tryMath(s, p string, x, y int) bool {
    // var save byte
    if p[y] >= 'a' && p[y] <= 'z'{
        if s[x] != p[y]{
            return false
        } else {
            for x < len(s) && s[x] == p[y] && y < len(p) - 1 && p[y+1] == '*' {
                x++
            }
        }
    }
    return false
}

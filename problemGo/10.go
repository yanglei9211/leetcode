package main

import "fmt"

func main() {
    s := "aab"
    p := "c*a*b"
    fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
    for i := 0; i < len(p); i++ {
        fmt.Println(i)
        if p[i] != '*' && tryMatch(s, p, 0, i) {
            return true
        }
    }
    return false
}

func tryMatch(s, p string, x, y int) bool {
    var save byte
    for x < len(p) && y < len(p){
        if p[y] >= 'a' && p[y] <= 'z'{
            if s[x] != p[y]{
                return false
            } else {
                if y < len(s) && s[x] == save {
                    for x < len(s) && s[x] == save {
                        x++
                    }
                    y++
                }
                y++
            }
        } else if p[y] == '.' {
            save := s[x]
            if y < len(p) - 1 && p[y+1] == '*'{
                for x < len(s) && s[x] == save{
                    x++
                }
                y++
            }
            y++
        }
    }
    if x == len(s) {
        return true
    } else {
        return false
    }
}

package main

import (
    "fmt"
    //"hash/fnv"
    "crypto/md5"
)

func a(x string) []byte{
    /*
    h := fnv.New64()
    h.Reset()
    h.Write([]byte(x))
    r := h.Sum64()
    */
    h := md5.New()
    h.Write([]byte(x))
    r := h.Sum(nil)
    fmt.Println(r)
    fmt.Printf("%x\n", r)
    return r
}


func main(){
    s := "aabb"
    a(s)
    mk := fmt.Sprintf()
    //var res = a(s)
    //fmt.Println(res)

    /*
    n := 0
    defer func(){
        n++;
        fmt.Printf("fi %d\n", n)
    }()
    fmt.Printf("main %d\n", n)
    */
    //test := "he1h3eh1he"
    /*
    reg, err := regexp.Compile(`\d+`)
    if err != nil{
        fmt.Println("aaaa")
    } else {
        fmt.Printf("%s,%v\n", reg.Find([]byte("1233213123")), err)
    }
    */
// "Hello",
}


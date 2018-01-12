package main

import (
	"fmt"
	"strings"
)

//Given an absolute path for a file (Unix-style), simplify it.
//
//For example,
//path = "/home/", => "/home"
//path = "/a/./b/../../c/", => "/c"

func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	fmt.Println(ss)
	curPath := make([]string, 0)
	for _, s := range ss {
		if len(s) == 0 {
			continue
		} else if s == "." {
			continue
		} else if s == ".." {
			if len(curPath) > 0 {
				curPath = curPath[0 : len(curPath)-1]
			}
		} else {
			curPath = append(curPath, s)
		}

		//fmt.Printf("-%s-\n", s)
	}
	res := "/"
	for i, s := range curPath {
		res += s
		if i < len(curPath)-1 {
			res += "/"
		}
	}
	return res
}

func main() {
	a := "/a/./b/../../c/cd//dc/"
	fmt.Println(simplifyPath(a))
}

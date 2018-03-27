package main

import "fmt"

func count(rt *TreeNode) int {
	if rt == nil {
		return 0
	}
	ls := count(rt.Left)
	rs := count(rt.Right)
	fmt.Println(rt.Val, ls, rs)
	if ls == -1 || rs == -1 {
		return -1
	} else {
		if ls < rs {
			ls, rs = rs, ls
		}
		if ls-rs > 1 {
			return -1
		} else {
			return ls + 1
		}
	}

}

func isBalanced(root *TreeNode) bool {
	r := count(root)
	if r != -1 {
		return true
	} else {
		return false
	}
}

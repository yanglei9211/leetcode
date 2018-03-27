package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	fmt.Println(nums)
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		cur := TreeNode{nums[0], nil, nil}
		return &cur
	} else {
		mid := 0 + len(nums)>>1
		rt := TreeNode{nums[mid], nil, nil}
		if mid > 0 {
			l := sortedArrayToBST(nums[:mid])
			rt.Left = l
		}
		if mid+1 < len(nums) {
			r := sortedArrayToBST(nums[mid+1:])
			rt.Right = r
		}
		return &rt
	}
}

func inorderTraversal(r *TreeNode) {
	if r == nil {
		return
	}
	inorderTraversal(r.Left)
	fmt.Printf("%d ", r.Val)
	inorderTraversal(r.Right)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	res := sortedArrayToBST(a)
	inorderTraversal(res)
}

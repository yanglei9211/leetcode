package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(rt *TreeNode, sum int) bool {
	if rt == nil {
		return false
	}
	if rt.Val == sum && rt.Left == nil && rt.Right == nil {
		return true
	}

	ls := hasPathSum(rt.Left, sum-rt.Val)
	if ls {
		return true
	}
	rs := hasPathSum(rt.Right, sum-rt.Val)
	if rs {
		return true
	}
	return false
}

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	maxHeight := 0
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	if leftHeight < rightHeight {
		maxHeight = rightHeight + 1
	} else {
		maxHeight = leftHeight + 1
	}
	return maxHeight
}


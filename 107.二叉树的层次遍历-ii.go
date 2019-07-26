/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	res = getTreeValue(root, 0, res)
	return reverse(res)
}

func getTreeValue(root *TreeNode, depth int, res [][]int) [][]int{
	if root == nil {
		return res
	}
	if depth + 1 > len(res) {
		tmp := []int{}
		res = append(res, tmp)
	}
	res[depth] = append(res[depth], root.Val)
	if root.Left != nil {
		res = getTreeValue(root.Left, depth + 1, res)
	}
	if root.Right != nil {
		res = getTreeValue(root.Right, depth + 1, res)
	}
	return res
}

func reverse(res [][]int) [][]int {
	for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

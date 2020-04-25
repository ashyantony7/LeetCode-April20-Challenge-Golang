/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type Pair struct {
	diameter int
	depth    int
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs(root *TreeNode) Pair {

	if root == nil {
		return Pair{0, 0}
	} else {
		left := dfs(root.Left)
		right := dfs(root.Right)
		diameter := max(left.depth+right.depth, max(left.diameter, right.diameter))
		return Pair{diameter, max(left.depth, right.depth) + 1}
	}

}

func diameterOfBinaryTree(root *TreeNode) int {
	res := dfs(root)
	return res.diameter
}

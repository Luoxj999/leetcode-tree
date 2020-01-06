/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

func solve(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if (a != nil && b != nil && a.Val == b.Val) {
		return solve(a.Left, b.Right) && solve(a.Right, b.Left)
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
    return solve(root, root)
}
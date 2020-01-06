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

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if (p == nil && q == nil) {
		return true
	} else if ((p == nil && q != nil) || (p != nil && q == nil)) {
		return false;
	}

	if (p.Val == q.Val && isSameTree(p.Left, q.Left) == true && isSameTree(p.Right, q.Right) == true) {
		return true
	} else {
		return false
	}
}
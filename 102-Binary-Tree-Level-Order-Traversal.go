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

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if (root == nil) {
		return result
	}

	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		vals := make([]int, 0)
		for i := 0; i < len(curLevel); i++ {
			tn := curLevel[i]
			vals = append(vals, tn.Val)
			if (tn.Left != nil) {
				nextLevel = append(nextLevel, tn.Left)
			}
			if (tn.Right != nil) {
				nextLevel = append(nextLevel, tn.Right)
			}
		}
		result = append(result, vals)
		curLevel = nextLevel
	}

	return result
}
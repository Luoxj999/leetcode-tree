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

func rightSideView(root *TreeNode) []int { 
	result := make([]int, 0)
	if (root == nil) {
		return result
	}

	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for len(curLevel) > 0 {
		result = append(result, curLevel[len(curLevel) - 1].Val)
		nextLevel := make([]*TreeNode, 0)
		for i := 0; i < len(curLevel); i++ {
			tn := curLevel[i]
			if tn.Left != nil {
				nextLevel = append(nextLevel, tn.Left)
			}
			if tn.Right != nil {
				nextLevel = append(nextLevel, tn.Right)
			}
		}
		curLevel = nextLevel
	}

	return result
}


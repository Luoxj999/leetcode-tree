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

func inorderTraversal(root *TreeNode) []int {
	var result []int
	
	if (root == nil) {
		return result
	}

	var stk []*TreeNode
	tn := root;
	for {
		if (tn != nil) {
			stk = append(stk, tn)
			tn = tn.Left
		} else {
			if (len(stk) == 0) {
				return result
			}

			tn = stk[len(stk) - 1]
			stk = stk[:len(stk) - 1]
			result = append(result, tn.Val)
			tn = tn.Right
		}
	}
}
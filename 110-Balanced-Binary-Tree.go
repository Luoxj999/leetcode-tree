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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func betweenOne(a int, b int) bool {
	if a == b {
		return true
	}
	temp := a - b
	if -1 <= temp && temp <= 1 {
		return true
	} 
	return false
}

func solve(root *TreeNode) (int, bool) {
	if (root == nil) {
		return 0, true
	}

	leftHeight, leftBalanced := solve(root.Left)
	rightHeight, rightBalanced := solve(root.Right)

	if leftBalanced == false || rightBalanced == false {
		return 0, false
	}

	if betweenOne(leftHeight, rightHeight) == false {
		return 0, false
	}
	
	h := max(leftHeight, rightHeight)
	return h + 1, true
}

func isBalanced(root *TreeNode) bool {
    _, result := solve(root)
	return result
}
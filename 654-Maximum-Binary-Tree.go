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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	maxNum := 0
	maxIdx := -1
	for i := 0; i < len(nums); i++ {
		if maxIdx == -1 || nums[i] > maxNum {
			maxNum = nums[i]
			maxIdx = i
		}
	}

	temp := &TreeNode{Val: maxNum}
	temp.Left = constructMaximumBinaryTree(nums[:maxIdx])
	temp.Right = constructMaximumBinaryTree(nums[maxIdx + 1:])
	return temp
}

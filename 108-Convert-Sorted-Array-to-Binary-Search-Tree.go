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

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	// 构造一颗满二叉树
// 	total := len(nums)
// 	root := &TreeNode{0, nil, nil}
// 	total--
// 	level := make([]*TreeNode, 0)
// 	level = append(level, root)
// 	for total > 0 {
// 		nextLevel := make([]*TreeNode, 0)
// 		for i := 0; i < len(level); i++ {
// 			curNode := level[i]
// 			if total > 0 {
// 				curNode.Left = &TreeNode{0, nil, nil}
// 				total--
// 				nextLevel = append(nextLevel, curNode.Left)
// 			} else {
// 				break
// 			}

// 			if total > 0 {
// 				curNode.Right = &TreeNode{0, nil, nil}
// 				total--
// 				nextLevel = append(nextLevel, curNode.Right)
// 			} else {
// 				break
// 			}
// 		}
// 		level = nextLevel
// 	}

// 	// 中序遍历二叉树，赋值
// 	curNumIdx = 0
// 	var stk []*TreeNode
// 	tn := root
// 	for {
// 		if (tn != nil) {
// 			stk = append(stk, tn)
// 			tn = tn.Left
// 		} else {
// 			if (len(stk) == 0) {
// 				return root
// 			}

// 			tn = stk[len(stk) - 1]
// 			stk = stk[:len(stk) - 1]
// 			tn.Val = nums[curNumIdx]
// 			curNumIdx++
// 			tn = tn.Right
// 		}
// 	}
// }

// 递归地一半左一半右，就平衡了
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return buildTree(&nums, 0, len(nums) - 1)
}

func buildTree(nums *[]int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	midNode := &TreeNode{nums[mid], nil, nil}
	left := buildTree(nums, l, mid - 1)
	right := buildTree(nums, mid + 1, r)
	midNode.Left = left
	midNode.Right = right
	return midNode
}
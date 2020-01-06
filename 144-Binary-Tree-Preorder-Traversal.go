/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//  func visit(tn *TreeNode, result *[]int) {
// 	if (tn == nil) {
// 		return
// 	}
// 	*result = append(*result, tn.Val)
// 	visit(tn.Left, result)
// 	visit(tn.Right, result)
// }

// func preorderTraversal(root *TreeNode) []int {
//     var result []int
// 	visit(root, &result)
// 	return result
// }

func preorderTraversal(root *TreeNode) []int {
	var result []int
	
	if root == nil {
		return result
	}

	var stk []*TreeNode
	stk = append(stk, root)
	for len(stk) > 0 {
		tn := stk[len(stk) - 1]
		stk = stk[:len(stk) - 1]
		result = append(result, tn.Val)
		if (tn.Right != nil) {
			stk = append(stk, tn.Right)
		}
		if (tn.Left != nil) {
			stk = append(stk, tn.Left)
		}
	}

	return result
}
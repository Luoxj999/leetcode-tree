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


func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if (t1 == nil && t2 == nil) {
        return nil
    }

    temp := new(TreeNode)
    temp.Val = 0

    if (t1 != nil) {
        temp.Val += t1.Val 
    }
    if (t2 != nil) {
        temp.Val += t2.Val 
    }
    if (t1 == nil) {
        temp.Left = mergeTrees(nil, t2.Left)
        temp.Right = mergeTrees(nil, t2.Right)
    } else if (t2 == nil) {
        temp.Left = mergeTrees(t1.Left, nil)
        temp.Right = mergeTrees(t1.Right, nil)
    } else {
        temp.Left = mergeTrees(t1.Left, t2.Left)
        temp.Right = mergeTrees(t1.Right, t2.Right)
    }
    return temp
}
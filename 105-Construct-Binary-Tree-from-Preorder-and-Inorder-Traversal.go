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

func find(preorder []int, inorder []int, pl int, pr int, il int, ir int, *TreeNode tn) *TreeNode {
    if (pl > pr) {
        return
    }

    rootp := pl
    rooti := -1
    leftLen := 0
    for i := il; i < ir; i++ {
        if (inorder[i] == root) {
            rooti = i
        } else {
            leftLen++
        }
    }
    
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    var tn *TreeNode
    tn = find(0, len(preorder) - 1, 0, len(inorder) - 1, tn)
}
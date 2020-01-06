/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import (
	"fmt"
)
 
type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode) *TreeNode {
	if (root == nil) {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	left := solve(root.Left)
	right := solve(root.Right)

	if right == nil {
		right = left
	}

	if left != nil {
		left.Right = root.Right
		root.Right = root.Left
	}
	root.Left = nil
	
	return right
}

// 把左子树放右边，右子树放左子树的最末端的右边
func flatten(root *TreeNode) {
    solve(root)
}

func visit(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	visit()
}

func main() {
	a := &TreeNode{3, nil, nil}
	b := &TreeNode{4, nil, nil}
	c := &TreeNode{2, a, b}
	d := &TreeNode{6, nil, nil}
	e := &TreeNode{5, nil, d}
	f := &TreeNode{1, c, e}
	flatten(f)
}

    1
   / \
  2   5
 / \   \
3   4   6

2   
/ \   
3   4   



1
 \
  2
   \
    3
     \
      4
	   \
	    5
	     \
	      6
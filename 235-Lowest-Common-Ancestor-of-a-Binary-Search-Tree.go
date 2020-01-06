/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

package main

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

func solve(root, p, q *TreeNode) (*TreeNode, bool) {
	if (root == nil) {
		return nil, false
	}

	tn1, b1 := solve(root.Left, p, q)
	tn2, b2 := solve(root.Right, p, q)

	if (b1 == true) {
		return tn1, b1
	} else if (b2 == true) {
		return tn2, b2
	}

	var pp *TreeNode
	var qq *TreeNode

	if tn1 != nil && tn1.Val == p.Val {
		pp = tn1
	} 
	if tn1 != nil && tn1.Val == q.Val {
		qq = tn1
	} 
	if tn2 != nil && tn2.Val == p.Val {
		pp = tn2
	} 
	if tn2 != nil && tn2.Val == q.Val {
		qq = tn2
	} 
	
	if root.Val == p.Val && qq != nil {
		return root, true
	}

	if root.Val == q.Val && pp != nil {
		return root, true
	}

	if pp != nil && qq != nil {
		return root, true
	}

	if root.Val == p.Val {
		return root, false
	} else if root.Val == q.Val {
		return root, false
	}

	if (pp != nil) {
		return pp, false
	} else if (qq != nil) {
		return qq, false
	} else {
		return nil, false
	}
}

//返回非空表示找到或者p或者q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 不用往下找了，结果肯定在上层或者它自己
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 左右都有返回，最近公共祖先一定是它
	if left != nil && right != nil {
		return root
	}

	// 这里返回的可能是结果或者p或者q
	if left != nil {
		return left
	}
	// 这里返回的可能是结果或者p或者q
	if right != nil {
		return right
	}
	// 没找到
	return nil
}

func main() {
	a := &TreeNode{3, nil, nil}
	b := &TreeNode{5, nil, nil}
	c := &TreeNode{4, a, b}
	d := &TreeNode{0, nil, nil}
	e := &TreeNode{2, d, c}
	f := &TreeNode{7, nil, nil}
	g := &TreeNode{9, nil, nil}
	h := &TreeNode{8, f, g}
	i := &TreeNode{6, e, h}
	r := lowestCommonAncestor(i, e, h)
	fmt.Println(r.Val)
}
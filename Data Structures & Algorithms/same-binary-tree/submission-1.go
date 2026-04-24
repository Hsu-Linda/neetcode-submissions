/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1 2
// 

func isSameTree(p *TreeNode, q *TreeNode) bool {
    // exclude the p or q are nil
	if p == nil && q == nil {
		return true
	} else if p == nil || q== nil {
		return false
	}

	// compare the value
	if p.Val != q.Val {
		return false
	}
	

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// p's value < root's value, then left side, otherwise right side
// q's value < root's value, then left side, otherwise right side
// so if p is smaller and q is bigger than rootm then output is root
// opposite  the output also is root
// if they are both right side or both left side 
// then again put the left/right node as root 
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val == root.Val || q.Val == root.Val {
		return root
	} else if ( p.Val < root.Val && root.Val < q.Val ) ||
		( q.Val < root.Val && root.Val < p.Val ) {
			return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}

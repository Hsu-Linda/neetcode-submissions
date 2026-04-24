/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// []
// []_[]
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max (first, second int ) int {
	if first > second {
		return first
	}
	return second
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// root currentNode's value != subRoot root's value, then currentNode in Root is not the subRoot's root
// we can consider the current's left or right as subRoot root's root

// if currentNode's value == subRoot root's value
// then we can consider is currentNode's left tree is same to subroot's left tree 
// and also currentNode's right is same to subRoot's right 

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	} else if root == nil {
		// root is nil but subroot not nil
		return false
	}

	// root and subRoot both not nil
	if root.Val != subRoot.Val {
		// we can consider left or right node as root
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
	
	if isSame(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}

func isSame (first *TreeNode, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	} else if first == nil || second == nil {
		return false
	}

	if first.Val != second.Val {
		return false
	}

	return isSame(first.Left, second.Left) && isSame(first.Right, second.Right)
}

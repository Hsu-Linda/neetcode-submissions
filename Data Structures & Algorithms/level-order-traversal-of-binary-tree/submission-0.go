/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// root
// root.left root.right
// root.left.left root.left.right  root.right.left root.right.right

// [root]
// 1, [root.left root.right]
// 2, 3 [root.left.left root.left.right  root.right.left root.right.right]
// 4, 5, 6, 7 nil nil nil

// when we meet nil we can just skip

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
	
	output := make([][]int, 0)
	parrent := make([]*TreeNode, 0) // we don't want to include nil
	parrent = append(parrent, root)

	for len(parrent) != 0 {
		curV := make([]int, 0)
		nextParrent:= make([]*TreeNode, 0)
		
		for _, node := range parrent {
			if node == nil {
				continue
			}
			curV = append(curV, node.Val)
			if node.Left != nil {
				nextParrent = append(nextParrent, node.Left)
			} 
			if node.Right != nil {
				nextParrent = append(nextParrent, node.Right)
			}
		}
		
		output = append(output, curV)
		parrent = nextParrent
	}

	return output
}

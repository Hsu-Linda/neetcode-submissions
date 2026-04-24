/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// [1 2 3 4]

// 1 2 3 4 2 3 4
//         f

// if it's cycle it doesn't end -> but we can't know it's non end or not
func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]int) 
	// []
	for head != nil {
		_, ok := nodeMap[head]
		if !ok {
			nodeMap[head] = 0
		} else {
			return true
		}
		head = head.Next
	}

	return false	
}
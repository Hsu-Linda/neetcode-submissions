/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1,2,4 []
// 1,3,5 []
// [1][1][2][3][4][5]

// head -> [1][1][2][3][4][5]
// head next is [1,nil]
// shift head to [1,nil]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{}
	var curM *ListNode = head
	cur1 := list1
	cur2 := list2
	
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			curM.Next = cur2
			cur2 = cur2.Next
		} else if cur2 == nil {
			curM.Next = cur1
			cur1 = cur1.Next
		} else if cur1.Val >= cur2.Val {
			curM.Next = cur2
			cur2 = cur2.Next
		} else {
			curM.Next = cur1
			cur1 = cur1.Next
		}
		curM = curM.Next
	}

	return head.Next
}

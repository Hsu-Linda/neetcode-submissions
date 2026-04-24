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
    dump := &ListNode{}
	var curM *ListNode = dump
	
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curM.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			curM.Next = list1
			list1 = list1.Next
		} else if list1.Val >= list2.Val {
			curM.Next = list2
			list2 = list2.Next
		} else {
			curM.Next = list1
			list1 = list1.Next
		}
		curM = curM.Next
	}

	return dump.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// dummy   4, 3, 2
// 

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next : head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
		
	}
	return dummy.Next
}

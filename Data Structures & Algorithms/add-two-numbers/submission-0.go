/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{}
	cur := dump
	nextDigit := 0
	for {
		var vl1, vl2 int
		if l1 == nil && l2 == nil && nextDigit == 0 {
			break
		} 
		if l1 != nil {
			vl1 = l1.Val
		}
		if l2 != nil {
			vl2 = l2.Val
		}
		
		sum := vl1 + vl2 + nextDigit
		if sum > 9 {
			nextDigit = sum/10
			sum = sum%10
		} else {
			nextDigit = 0
		}
		cur.Next = &ListNode{Val : sum}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dump.Next
}

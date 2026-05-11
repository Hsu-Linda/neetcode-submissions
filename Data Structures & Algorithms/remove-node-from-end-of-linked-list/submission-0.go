/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dump := &ListNode {
		Val : 0,
		Next : head,
	}
	cur := dump
	for cur != nil {
		remove := cur
		nthNext := nthNext(remove, n)
		if nthNext != nil && nthNext.Next == nil {
			remove.Next = remove.Next.Next 
			break
		} else if nthNext == nil{
			return nil
		}
		
		cur = cur.Next
	}
	return dump.Next
}

// we nth still's not nil
// n+1th Next is nil 
// that's previous of remove item 
func nthNext (head *ListNode, n int) *ListNode  {
	cur := head
	for i:= 0; i < n; i++ {
		if cur == nil {
			//over the listNode
			return nil
		}
		cur = cur.Next
	}
	return cur
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 

// [1, 2, 3, 4, 5] 1, 3
// dump 1 2 3 4 5 
// dump 
// 

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dump := &ListNode{Next: head}
	cur := dump  // dump
	curIndex := 0 
	
	
	var reverseBefore *ListNode
	for cur != nil && curIndex < left {
		reverseBefore = cur // dump
		
		cur = cur.Next // 1
		curIndex ++  // 1
	}
	
	
	reverseEnd := cur // 1
	reverseHead := reverseEnd
	cur = cur.Next // 2
	curIndex ++ // 2

	for cur != nil && curIndex <= right {
		next := cur.Next
		cur.Next = reverseHead  // 3- 2 - 1
		reverseHead = cur
		
		cur = next //  3 4
		curIndex ++ // 3 4
	}

	reverseBefore.Next = reverseHead // dump - 3
	reverseEnd.Next = cur    // 1 -  4

	return dump.Next
}


func addHead(source *ListNode, newHead *ListNode) *ListNode {
	newHead.Next = source
	return newHead
}
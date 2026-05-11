/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// we need to know left and right each of them are longer than k or not 
// reverse or not change

// turtle and rabbit

// 0 1 2
// 0 2 

func reverseKGroup(head *ListNode, k int) *ListNode {
    output := &ListNode{}
	outputEnd := output
	curHead := head
	for {
		isFull, nextHead := cut(curHead,k)
		if isFull {
			reverseHead := reverse(curHead)
			outputEnd.Next = reverseHead
			outputEnd = curHead
		} else {
			outputEnd.Next = curHead
		}
		
		
		if nextHead == nil {
			return output.Next
		}
		curHead = nextHead
		outputEnd = initOutputEnd(outputEnd)
	}
	


}

// boundary_head, next_Boundary_Head
func cut(curBoundaryHead *ListNode, k int) (isFull bool, nextBoundaryHead *ListNode,){
	cur := curBoundaryHead
	if cur == nil {
		return false, nil
	}
	
	for i:=1; i<k; i++ {
		if cur.Next == nil {
			return false, nil
		} 
		cur = cur.Next
	}
	next := cur.Next
	cur.Next = nil
	return true, next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}


func initOutputEnd(cur *ListNode) *ListNode{
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}



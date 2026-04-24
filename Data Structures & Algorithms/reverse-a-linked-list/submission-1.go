/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// [0,1,2,3]
// [0]         -> [0,nil]
// [1]         -> [1, addr1]  addr1 point to the [0,nil]   [1]->[0]
// [2]         -> [2, addr2]  addr2 point to the [1, addr1] [2]->[1]->[0]
// [3]         -> [3, addr3]  addr3 point to the [2, addr2] [3]->[2]->[1]->[0]
// nil 

func reverseList(head *ListNode) *ListNode {

	cur := head
	var rCur *ListNode = nil
	
	for cur != nil {
		temp := &ListNode {Val: cur.Val, Next: rCur}
		rCur = temp
		cur = cur.Next
	}
	
	return rCur
}

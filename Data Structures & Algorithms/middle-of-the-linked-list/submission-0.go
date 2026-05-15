/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1    2,   3     
// 1, 2  3  fast nil 

// 1    2,   3    4
// 1,2  3,4  4,5  6  fast.Next = nil 

// 1 2 3 4 5 nil
// 1 2 3 4 5 6 nil

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

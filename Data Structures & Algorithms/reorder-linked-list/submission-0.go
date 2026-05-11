// [0, 1, 2, 3, 4, 5, 6]
// head -> 1 2 3 4 5 ... 6 
// firstP mid.Next   secondP 1  secondNext 2   secondP.Next -> firstP  1->0
// firstP 1   secondP 2  secondNext 3   2->1->0
// firstP 2   secondP 3  secondNext 4   
// firstP 3   secondP 4  secondNext 5
// firstP 4   secondP 5  secondNext 6
// firstP 5   secondP 6  secondNext nil
// secondP.Next -> firstP
// firstP = secondP
// secondP = secondNext
// secondNext = secondP.Next
// when secondNext is nil that's the end, (include this round)


func reorderList(head *ListNode) {
	mid := getMid(head)
	end := reverse(mid.Next)
	mid.Next = nil

	for head != nil && end != nil{
		tempFront := head.Next
		head.Next = end
		end = end.Next
		head = head.Next
		head.Next = tempFront
		head = head.Next
	}
	

}

// 0 1 2 3
// 0 1 
// 0 2  

// 0, 1, 2, 3, 4, 5, 6
// 0, 1, 2, 3
// 0, 2, 4, 6, next

// 0, 1, 2, 3, 4, 5
// 1, 2
// 2, 4, 5

// mid is include in right 
func getMid(head *ListNode) *ListNode {
	quick := head
	slow := head
	
	for quick.Next != nil && quick.Next.Next != nil{
		slow = slow.Next
		quick = quick.Next 
		quick = quick.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var cur *ListNode = nil
	next := head
	
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}
	return cur
}

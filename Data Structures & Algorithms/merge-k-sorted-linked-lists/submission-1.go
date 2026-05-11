/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// lists = [[1,2,4],[1,3,5],[3,6]]
// head1 1 2 4 nil
// head2 1 3 5 nil
// head3 3 6 nil
// merg: 1, 1, 2, 3, 3 ,4 ,5, 6

// lists = []
// head
// merg : nil

// lists = [[]]
// head1 nil
// merg : nil


func mergeKLists(lists []*ListNode) *ListNode {
    // initDump(lists)
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	dump := &ListNode{}
	cur := dump
	for {
		min := getMin(lists)
		if lists[min] == nil {
			break
		}
		cur.Next = lists[min]
		lists[min] = (*lists[min]).Next
		cur = cur.Next
	}

	return dump.Next
}

// already make sure length of lists longer than 1
// return the index of lists
func getMin(lists []*ListNode) int {
	min := 0
	for i :=1 ;i<len(lists); i++ {
		min = getMinBy2(lists, min, i)
	}

	return min
}

// if both are nil return nil
// if one of them not nil return smaller
// if both of them not nil and val are same, return firstone
func getMinBy2(lists[]*ListNode, first, second int) int {
	if lists[first] == nil && lists[second] == nil {
		return first
	} else if lists[first] == nil {
		return second
	} else if lists[second] == nil {
		return first
	} else if (*lists[first]).Val > (*lists[second]).Val {
		return second
	}
	return first
}

type ListNode struct{
	Value int
	Next *ListNode
}
type MyCircularQueue struct {
	DummyHead *ListNode
	End *ListNode
	
	Capacity int
	Count int 
}
// DummyHead -> 1 -> 2(End)
// DummyHead(End) if End pointer is Dump also which means Empty

func Constructor(k int) MyCircularQueue {
	dummy := &ListNode{}
	return MyCircularQueue {
		DummyHead: dummy,
		End: dummy,
		Capacity: k,
		Count: 0,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Count == this.Capacity {
		return false
	}
	this.End.Next = &ListNode{Value : value}
	this.End = this.End.Next
	this.Count ++
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.DummyHead.Next == nil {
		return false
	}
	if this.End == this.DummyHead.Next {
		this.End = this.DummyHead
	}
	this.DummyHead.Next = this.DummyHead.Next.Next

	this.Count --
	return true
}


func (this *MyCircularQueue) Front() int {
    if this.DummyHead.Next == nil {
		return -1
	}
	return this.DummyHead.Next.Value
}


func (this *MyCircularQueue) Rear() int {
    if this.End == this.DummyHead {
		return -1
	}
	return this.End.Value
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.End == this.DummyHead 
}


func (this *MyCircularQueue) IsFull() bool {
	return this.Count == this.Capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 
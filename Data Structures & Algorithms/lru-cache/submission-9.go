// 1, 2
// if we can successfly get, which means we put that before
//
type LRUCache struct {
    capacity int
	store map[int]*ListNode
	head *ListNode
	end *ListNode
}

type ListNode struct {
	Key int
	Val int
	Previous *ListNode
	Next *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		store: make(map[int]*ListNode),
		head: &ListNode{},
		end: nil,
	}
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.store[key]
	if !ok {
		return -1
	}
	
	if this.end != node {
        this.moveToEnd(node)
    }
	
	return node.Val
	
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.store[key]; ok {
		node.Val = value
		this.moveToEnd(node)
		return 
	}
	
	
	newNode := &ListNode{Key: key, Val: value, Next: nil}
	this.append(newNode)
	

	if len(this.store) > this.capacity {
		removeNode := this.head.Next
		this.remove(removeNode)
	}
}

func (this *LRUCache) remove (node *ListNode){
	delete(this.store, node.Key)
	node.Previous.Next = node.Next
	if node.Next != nil {
		node.Next.Previous = node.Previous
	}
	
	if node == this.end {
		if node.Previous == this.head {
			this.end = nil
		}
		this.end = node.Previous
	}
}

func (this *LRUCache) append(node *ListNode) {
	this.store[node.Key] = node
	node.Next = nil
	if this.end == nil {
		this.end = node
		this.head.Next= node
		node.Previous = this.head
		return 
	}
	
	this.end.Next = node
	node.Previous = this.end
	this.end = this.end.Next
}

func (this *LRUCache) moveToEnd(node *ListNode) {
	this.remove(node)
	this.append(node)
}



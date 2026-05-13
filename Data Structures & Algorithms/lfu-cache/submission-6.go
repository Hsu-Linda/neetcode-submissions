type Node struct {
	Key      int
	Value    int
	Frequent int
	
	Prev *Node
	Next *Node
}

type DoubleLinkedList struct{
	Head *Node
	Tail *Node
}

func NewList() *DoubleLinkedList {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head
	
	return &DoubleLinkedList {
		Head: head,
		Tail: tail,
	}
}

func (list *DoubleLinkedList) Add(node *Node) {
	prev := list.Tail.Prev
	prev.Next = node
	node.Prev = prev

	node.Next = list.Tail
	list.Tail.Prev = node
}

func(list *DoubleLinkedList) Remove(node *Node) {
	prev := node.Prev
	next := node.Next
	
	prev.Next = next
	next.Prev = prev

	node.Prev = nil
	node.Next = nil
}

type LFUCache struct {
	Capacity int
	MinFreq int
	
	KeyToNode map[int]*Node
	FrequentToList map[int]*DoubleLinkedList
}


func Constructor(capacity int) LFUCache {
   return LFUCache{
		Capacity:capacity,
		MinFreq:1,
		KeyToNode: make(map[int]*Node),
		FrequentToList: make(map[int]*DoubleLinkedList),
		
   }
}




func (this *LFUCache) Get(key int) int {
	node, ok := this.KeyToNode[key]
	if !ok {
		return -1
	}
	
	this.FrequentToList[node.Frequent].Remove(node)
	node.Frequent++
	list := this.GetList(node.Frequent)
	list.Add(node)
	return node.Value
}

func (this *LFUCache) GetList(freq int) *DoubleLinkedList{
	if list, ok := this.FrequentToList[freq]; ok {
		return list
	}

	list := NewList()
	this.FrequentToList[freq] = list
	return list

}


func (this *LFUCache) Put(key int, value int)  {
    node, ok := this.KeyToNode[key]
	if ok {
		node.Value = value
		this.FrequentToList[node.Frequent].Remove(node)
		node.Frequent++
		list := this.GetList(node.Frequent)
		list.Add(node)
		return
	}
	
	if len(this.KeyToNode) == this.Capacity {
		this.RemoveMin()
	}

    newNode := &Node{Key: key, Value: value, Frequent: 1}
	this.KeyToNode[key] = newNode

	list := this.GetList(1)
	list.Add(newNode)
	this.MinFreq = 1
}

func  (this *LFUCache) RemoveMin(){
	for {
		if list, ok := this.FrequentToList[this.MinFreq]; ok{
			if list.Head.Next != list.Tail {
				remove := list.Head.Next
				delete(this.KeyToNode, remove.Key)
				
				prev := list.Head
				next := list.Head.Next.Next
				prev.Next = next
				next.Prev = prev
				break
			}
		}
		this.MinFreq ++
	}
}
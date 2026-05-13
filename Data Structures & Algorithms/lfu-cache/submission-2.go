type ListNode struct {
	Value int
	Key int
	UseCount int
	Previous *ListNode
	Next *ListNode
}

// dummy - 1 - 2 compare to next useCount if same swich 
// dummy 2 - 1  
// store 1:1,  3:3
// useCount 1: addr3  2: ddr1

type FrequentInfo struct{
	DummyHead *ListNode
	Recent *ListNode
}

type LFUCache struct {
	Store map[int]*ListNode
	Frequent map[int]*FrequentInfo  // frequent : last node
	Capacity int
	MinFreq int
}


func Constructor(capacity int) LFUCache {
   return LFUCache{
		Store: make(map[int]*ListNode),
		Frequent: make(map[int]*FrequentInfo),
		Capacity:capacity,
		MinFreq:1,
   }
}


func (this *LFUCache) Get(key int) int {
	node, ok := this.Store[key]
	if !ok {
		return -1
	}
	
	this.FreqAdd(node)
	return node.Value
}


func (this *LFUCache) Put(key int, value int)  {
    node, ok := this.Store[key]
	if ok {
		node.Value = value
		this.FreqAdd(node)
		return 
	}
	
	if len(this.Store) == this.Capacity {
		this.Remove()
	}


    fmt.Printf("HERE")
    newNode := &ListNode{Key: key, Value: value, UseCount: 1}
	this.Store[key] = newNode
    freq, ok := this.Frequent[1]
    if !ok {
        freq = this.InitFreq(1)
    }
    freq.Recent.Next = newNode
    newNode.Previous = this.Frequent[1].Recent
    freq.Recent = freq.Recent.Next
	fmt.Printf("DEBUG_PUT: %v \n", freq.Recent.Key)
	
	this.MinFreq = 1

}

func (this *LFUCache) FreqAdd(node *ListNode){
	
	
	//  from current frequent linkedlist
	node.Previous.Next = node.Next
	if this.Frequent[node.UseCount].DummyHead.Next == nil {
	    this.Frequent[node.UseCount].Recent = this.Frequent[node.UseCount].DummyHead
	}
	
	if node.Next!= nil {
		node.Next.Previous = node.Previous
	}

    node.UseCount ++
	// append to frequent ++ linkedlist
	frequent, ok := this.Frequent[node.UseCount]
	if !ok {
        frequent = this.InitFreq(node.UseCount)
	} 
    frequent.Recent.Next = node
    node.Previous = frequent.Recent
    frequent.Recent = frequent.Recent.Next
	
}
 
func (this *LFUCache) Remove() {
	this.RenewMinFreq()
	freq, _:= this.Frequent[this.MinFreq]
	delete(this.Store, freq.DummyHead.Next.Key)
	freq.DummyHead.Next = freq.DummyHead.Next.Next
	if freq.DummyHead.Next != nil {
	    freq.DummyHead.Next.Previous = freq.DummyHead
	}else {
	    freq.Recent = freq.DummyHead
	}
}

func (this *LFUCache) RenewMinFreq() {
	for {
	    fmt.Printf("DEBUG RENEWMIN %v \n", this.MinFreq)
		if freq, ok := this.Frequent[this.MinFreq]; ok {
		    fmt.Printf("DEBUG RENEWMIN DummyHead.Next %v \n", freq.DummyHead.Next)
		    fmt.Printf("DEBUG RENEWMIN Recent %v \n", freq.Recent.Key)
			if freq.Recent != freq.DummyHead {
				fmt.Println("BREAK")
				break
			}
			delete(this.Frequent, this.MinFreq)
		}
		this.MinFreq ++
	}
}

func (this *LFUCache) InitFreq(freq int) *FrequentInfo{
    if info, ok := this.Frequent[freq]; ok{
        return info
    }
    
    dummy := &ListNode{}
    initInfo := &FrequentInfo{
        DummyHead: dummy,
        Recent: dummy,
    }
    this.Frequent[freq] = initInfo
    
    return initInfo
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */

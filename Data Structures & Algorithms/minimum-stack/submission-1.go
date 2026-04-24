// O(1)
// Get min -> min heap 
// min heap -> push O(logn) pop O(1)
// we need to keep the order  slice [1,2,3,423424, 2343] O(1)
// push we can keep that in slice and async to construct the min heap in the background 

type MinStack struct {
	stack []int
	minStack []int // big -> small
}

func Constructor() MinStack {
	return MinStack{
		stack : make([]int, 0, 10),
		minStack : make([]int, 0, 10),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || 
        val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	poped := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.minStack[len(this.minStack)-1] == poped {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -99999
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return -99999
	}
	return this.minStack[len(this.minStack)-1]
}

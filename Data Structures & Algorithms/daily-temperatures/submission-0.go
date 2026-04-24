// when we meet some biger the the biggest then pop out all 
// [30,0][] curIndex= 1 curVal= 38
// [38][1,] //since 38>30, pop out 30 and push 38
// [38] curIndex=2 curVal=30
// [(38,1),(30,2)] curIndex=3 curVal=36
// [(38,1)(36,3)][1,_,1] // since 36 > 30, pop out 30 and let the output index 2 is curIndex(3)-2=1
// [(38,1),(36,3)][1,_,1] curInex=4 curVal=35
// [(38,1)(35,4)][1,_,1] curIndex=5 curVal=40
// [][1,4,1,_,1]
// we maintain a decrease list big->small, and always popout from the end, which means index=len(listA)-1
// if curVal is bigger than the end of list val,
// 1. pop out that
// 2. make the output list, it's index of A is curIndex-popOut's index, A is popOut's index
// 3. for loop till the end of decrease list is bigger than curVal
// if curVal is smaller than the end of the list val
// 1. push it into the decrease list

// and finally we can't foun morer cur val, but the maintain list still have elements, then just let output as 0

type tempDay struct {
	temp int
	day int // index
}

func dailyTemperatures(temperatures []int) []int {
	output := make([]int, len(temperatures))
	waiting := make([]tempDay, 0, len(temperatures))
	waiting = append(waiting, tempDay {temp: temperatures[0], day: 0})
	
	// get current index and value
	for curIndex:= 1; curIndex<len(temperatures); curIndex++ {
		// get the end of the waiting list's index and value
		// if the waiting list empty, which means we don't have any compare current base line
		// -> just push current into waiting list
		for len(waiting) > 0 &&
			temperatures[curIndex] > waiting[len(waiting)-1].temp {
			found := waiting[len(waiting)-1]
			waiting = waiting[:len(waiting)-1]
			output[found.day] = curIndex-found.day
		}
		waiting = append(waiting, tempDay{temp: temperatures[curIndex], day:curIndex})
	}
	return output
}


// find max two * dis
// if it's far and also it's bigger 
// if it's far but smaller 
// if it's close but bigger
// if it's cloase and it's smaller -> we don't want 

func maxArea(heights []int) int {
	if len(heights) < 2 {
		return 0
	}
	
	left := 0
	right := len(heights)-1
	max := area(heights, left, right)
	for left < right {
		if heights[left] < heights[right] {
			left ++
		} else {
			right --
		}
		area := area(heights, left, right)
		if area > max {
			max = area
		}
	}
	return max
}



func area (heights []int, indexL, indexR int) int {
	return (indexR-indexL) * min(heights[indexL], heights[indexR])
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func maxArea(heights []int) int {
	max := 0
	for iL, vL := range heights {
		for iR := iL+1; iR<len(heights); iR++ {
			area := (iR-iL) * min(vL, heights[iR])
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

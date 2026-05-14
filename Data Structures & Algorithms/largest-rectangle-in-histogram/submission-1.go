type Rectangle struct {
	height int
	start  int
}
// prevHeigh = 0 [ 7       , max = 0
// prevHeigh = 7 [ 1      , max = 7
// prevHeigh = 7 [ 1 7     , max = 7

func largestRectangleArea(heights []int) int {
	activeRectangle := make([]Rectangle, 0, len(heights))
	maxArea := 0

//  1 ,2 , 3 ,4 ,5, ,6 ,7, 3
	for i:=0; i<len(heights); i++ {
		start := i
		for len(activeRectangle) > 0 && 
			heights[i] <= activeRectangle[len(activeRectangle)-1].height {
			
			top := activeRectangle[len(activeRectangle)-1]
			activeRectangle = activeRectangle[:len(activeRectangle)-1]
			
			area := top.height * (i - top.start)
			if area > maxArea  {
				maxArea = area
			}

			start = top.start
		}

		activeRectangle = append(activeRectangle, Rectangle{height: heights[i], start: start })
	}

	for _, rec := range activeRectangle {
		area := rec.height * (len(heights)- rec.start)
		if area > maxArea {
			maxArea = area
		}
	}		

	return maxArea
}

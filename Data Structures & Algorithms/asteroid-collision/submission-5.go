// [11, 7, -10, 9]
// [11, 9]
// index is relative position in space
// absolute value -> size
// same speed 
// positive is right negative left

// [2,4,-4,-1]
// 4-> -4 <- both explode
// 2 -> <- -1 2

// if positive put it int
// [2, 4 ] 
// if negative 
// [2,4 ] get the last compare which one is bigger 
// [2,4, ] 4,-4 
// if same pop out the last 
// if positive one is bigger, then the negative one is explode
// , just continue don't need to do anything to the slice
// if negative is bigger 
// then meet the next positvie (get the last of slice again)
// if don't have any more positive in that 
// [put it int]

// [7, 9, -8]
// [7]
// [7] 9
// [7, 9] -8
// 


func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	
	for _, e := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, e)
			continue
		}
		
		if e > 0 {
			stack = append(stack, e)
			continue
		}
		
		// e is negative // len(stack) > 0
		// [-2,-2,-2]
		// [-2, ] -2
		last := stack[len(stack)-1]
		
		// last positive e negative len>0
		for {	
			if last < 0 {
				stack = append(stack, e)
				break
			} else if last > -1*e {
				// e byebye
				break
			} else if last == -1*e {
				// e byebye
				stack = stack[:len(stack)-1]
				break
			}
			// last < |e|
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				last = stack[len(stack)-1]
				continue	
			}
			stack = append(stack, e)
			break
		}

	}
	return stack
}

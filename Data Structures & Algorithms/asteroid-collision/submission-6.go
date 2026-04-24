// stack:[]   [10,2, -5] 
// [10]  [2,-5]
// [10, 2]  [-5]
// [10, -5]
// [10]


func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	
	for len(asteroids) != 0 {
		if len(stack) == 0 {
			stack = append(stack, asteroids[0])
			asteroids = asteroids[1:]
			continue
		}
		
		if asteroids[0] > 0 {
			stack = append(stack, asteroids[0])
			asteroids = asteroids[1:]
			continue
		}
		
		last := stack[len(stack)-1]
		if last < 0 {
			stack = append(stack, asteroids[0])
			asteroids = asteroids[1:]
			continue
		} else if last > -1*asteroids[0] {
			asteroids = asteroids[1:]
			continue
		} else if last == -1*asteroids[0] {
			asteroids = asteroids[1:]
			stack = stack[:len(stack)-1]
			continue
		}  else {
			stack = stack[:len(stack)-1]
			continue
		}

	}
	return stack
}

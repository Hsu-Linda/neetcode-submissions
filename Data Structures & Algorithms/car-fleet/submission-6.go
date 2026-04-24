// 1 4 7 10  1+3*x  
// 4 6 8 10  4+2*x
// when 1+3*x >= 4+2*x && 4+2*x <= target as 1 fleet



// position in increse order 
// when get the closet one(7) and count how much time it needs to spend to arrive = 3 hours
// for loop from closer to the destination car to the farrest car
// each of the use 3 hours to count if they can arrive or not 
// if we meet any of them, can't arrive destination then stop, since just only one-lane
// and then count that one how much time it needs to spend to arrive -> same strategy above
// 7+1*x = 10 ; x=3
// 4+2*3= 10
// 1+2*3 =7   7 < 10, which means 3 hours can't arrive
// 1+2*x=10;   x=4.5
// 0+1*4.5 =4.5  4.5<10
type car struct {
	position int 
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	
	// stack monotonic
	// big to small
	cars := make([]car, 0, len(position))
	for i, e := range position {
		newone := car{position: e, speed: speed[i]}
		if len(cars) == 0 {
			cars = append(cars, newone)
			continue
		}
		left, right := 0, len(cars)-1
		var mid int
		for left <= right {
			mid = (left + right)/2
			if e < cars[mid].position {
				right = mid-1
				continue
			}
			left = mid+1
		}
		cars = append(cars, car{})
		if right < mid {
            // newone needs just one the mid left side
            copy(cars[mid+1:], cars[mid:])
            cars[mid] = newone
        } else {
            // newone needs just on the mid right side
            copy(cars[mid+2:], cars[mid+1:])
            cars[mid+1] = newone
        }
	}	
	
	// get the end of the cars 
	fleets := 1
	fleetTop := cars[len(cars)-1] // at least one, won't panic
	cars = cars[:len(cars)-1]
	for len(cars) > 0 {
		fleetTopHours := float64(target-fleetTop.position)/float64(fleetTop.speed)
		cur := cars[len(cars)-1]
		cars = cars[:len(cars)-1]
		if float64(cur.position)+float64(cur.speed) * fleetTopHours >= float64(target) {
			continue
		}
		fleets ++
		fleetTop = cur
	}
	return fleets
}

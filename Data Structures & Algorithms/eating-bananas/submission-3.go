func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	if len(piles) == 1 {
		k := piles[0]/h
		if piles[0] % h > 0 {
			return k+1
		}
		return k
	}

	if h < len(piles) {
		// we can't find the answer
		return -1
	}

	// we need to find the biggest val in piles, as the k(output)
	big := piles[0]
	for cur :=1; cur<len(piles); cur ++ {
		if big < piles[cur] {
			big = piles[cur]
		}
	}

	if h == len(piles) {
		return big	
	}

	// h > len(piles)
	// i.e. piles = [1,4,3,2], h = 9
	// the biggest k = the biggest of piles
	// the smalles k = 0
	// sorted so we can use binary search
	// k in [1,2,3,4]   big is 4
	lowestSpeed :=1
	highestSpeed := big // 4 -> 2
	for lowestSpeed < highestSpeed{
		midSpeed := (lowestSpeed+highestSpeed)/2 // 2 -> 1
		sumH :=0
		for _, e := range piles {
			if e%midSpeed > 0 {
				sumH += e/midSpeed +1
			} else {
				sumH += e/midSpeed
			}
		}
		
		// if mid is 1, sumH is 10 -> 10 > h(9)
		// if mid is 3 sumH is 5 < 9(h)
		// if mid is 2 sumH is 6 < 9(h)
		if sumH == h {
            if highestSpeed == lowestSpeed {
                return midSpeed    
            } else {
                highestSpeed = midSpeed
            }
		} else if sumH > h && midSpeed+1 <= highestSpeed { // sumH > h is invalid
			lowestSpeed = midSpeed+1
		} else if sumH < h {
			highestSpeed = midSpeed
		}
	}
	return highestSpeed
}

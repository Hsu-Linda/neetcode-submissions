func maxProfit(prices []int) int {
	minPrice:= prices[0]
	bestProfit := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] > minPrice {
			curProfit := prices[i]-minPrice
			if curProfit > bestProfit {
				bestProfit = curProfit
			}
		} else {
			minPrice = prices[i]
		}
	}

	return bestProfit
}

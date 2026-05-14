// (100,1) (85, 6)  

type StockSpanner struct {
	historyPrice []HistoryPriceInfo
}

type HistoryPriceInfo struct {
	Price int
	Days  int // include itself
}
func Constructor() StockSpanner {
	return StockSpanner{
		historyPrice: make([]HistoryPriceInfo, 0, 10),
	}
}

func (this *StockSpanner) Next(price int) int {
	daysTotal := 0
	
	for len(this.historyPrice) > 0 &&
		this.historyPrice[len(this.historyPrice)-1].Price <= price {
		
		daysTotal += this.historyPrice[len(this.historyPrice)-1].Days
		this.historyPrice = this.historyPrice[:len(this.historyPrice)-1]

	}

	this.historyPrice = append(this.historyPrice, HistoryPriceInfo{Price: price, Days: daysTotal+1})
	return daysTotal+1
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 
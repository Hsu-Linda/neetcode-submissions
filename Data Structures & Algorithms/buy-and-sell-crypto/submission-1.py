# 10 1
# 10 1 5 6 
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        maxProfit = 0
        buyDay, sellDay = 0, 0
        
        for i in range(len(prices)):
            if prices[i] < prices[buyDay]:
                if sellDay > buyDay:
                    profit = prices[sellDay]-prices[buyDay]
                    if maxProfit < profit:
                        maxProfit = profit
                buyDay = i
                sellDay = i
            else:
                if prices[i] > prices[sellDay]:
                    sellDay = i

        return max(maxProfit,(prices[sellDay]-prices[buyDay]))

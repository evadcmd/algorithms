package main

func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		// We want to sell at high and buy at low
		if c := hold + prices[i] - fee; c > cash {
			cash = c
			// if we decided to buy it, then cash := hold + prices[i] - fee,
			// so h := cash - prices[i] = hold - fee,
			// h is < hold, vice verse.
		} else if h := cash - prices[i]; h > hold {
			hold = h
		}
	}
	return cash
}

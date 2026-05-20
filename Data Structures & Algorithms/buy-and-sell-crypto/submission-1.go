func maxProfit(prices []int) int {
	minprice := prices[0]
	maxprice := 0

	for _, prices := range prices {
		//ngecek price lebih kecil dari minprice
		if prices < minprice {
			minprice = prices
		}
		//ngubah maxprofit based on currprice - minprice
		if prices - minprice > maxprice {
			maxprice = prices - minprice
		}
	}
	return maxprice
}

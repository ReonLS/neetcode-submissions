func maxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0

	//iterate array
	for _, num := range prices{

		//check min as buy price
		if num < minPrice{
			minPrice = num
		}

		//check profit and max, if bigger then change max
		if num - minPrice > max{
			max = num - minPrice
		}
	}
	return max
}

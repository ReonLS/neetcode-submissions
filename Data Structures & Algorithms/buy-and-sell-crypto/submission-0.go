func maxProfit(prices []int) int {
	max := 0

	//pointer 1
	for i, first := range prices {
		//pointer 2, access val after i
		for _, j := range prices[i+1:] {
			if (j-first) > max {
				max = j-first
			}
		}
	}
	return max
}

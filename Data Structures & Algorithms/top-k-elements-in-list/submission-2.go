func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	inverse := make(map[int][]int)
	var final []int

	for _, num := range nums{
		freq[num]++
	}

	for elm, occur := range freq{
		inverse[occur] = append(inverse[occur], elm)
	}

	for i := len(nums); i > 0; i--{
		for _, elm := range inverse[i]{
			final = append(final, elm)

			if len(final) == k{
				return final
			}
		}
	}
	return final
}

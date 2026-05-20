func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	inverse := make([][]int, len(nums)+1)
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

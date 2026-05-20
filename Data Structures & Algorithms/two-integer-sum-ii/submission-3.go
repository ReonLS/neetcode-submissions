func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)
	final := []int{}

	for idx, num := range numbers{
		complement := target - num

		if _, ok := hash[complement]; ok{
			final = append(final, hash[complement], idx+1)
			return final
		}
		hash[num] = idx+1
	}
	return final
}

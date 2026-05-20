func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	final := []int{}

	for left < right{
		complement := target - numbers[left]

		if numbers[left] + numbers[right] == target{
			final = append(final, left+1, right+1)
			return final
		} else if numbers[right] > complement{
			right--
		} else {
			left++
		}
	}
	return final
}

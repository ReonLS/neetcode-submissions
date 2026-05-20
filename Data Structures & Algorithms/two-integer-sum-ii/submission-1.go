func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var data []int

	for left < right{
		if numbers[left] + numbers[right] > target{
			right--
		} else if numbers[left] + numbers[right] < target{
			left++
		} else {
			data = append(data, left+1, right+1)
			return data
		}
	}
	return data
}

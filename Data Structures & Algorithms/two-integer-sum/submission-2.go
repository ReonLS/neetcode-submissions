func twoSum(nums []int, target int) []int {
    arr := make(map[int]int)
	result := []int{}

	for idx, num := range nums{
		complement := target - num

		if value, ok := arr[complement]; ok{
			result = append(result, value, idx)
			return result
		} else {
			arr[num] = idx
		}
	}
	return result
}

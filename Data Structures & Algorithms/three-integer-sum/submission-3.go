func threeSum(nums []int) [][]int {
	var data [][]int

	sort.Ints(nums)

	for idx, num := range nums{
		need := 0 - num
		left, right := idx+1, len(nums)-1

		if idx > 0 && nums[idx] == nums[idx-1]{
			continue
		}

		for left < right{
			sum := nums[left] + nums[right]

			if sum < need{
				left++
			} else if sum > need{
				right--
			} else {
				var temp = []int{num, nums[left], nums[right]}
				data = append(data, temp)
				
				// skip duplicate lefts
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// skip duplicate rights
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// then move past the current pair
				left++
				right--
			}
		} 
	}
	return data
}

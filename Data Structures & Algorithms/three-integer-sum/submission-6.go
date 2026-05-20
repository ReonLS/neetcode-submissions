func threeSum(nums []int) [][]int {
	var data [][]int

	sort.Ints(nums)

	for idx, num := range nums{
		//all positive int
		if num > 0 {
			break
		}

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
				data = append(data, []int{num, nums[left], nums[right]})

				// skip duplicate lefts
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				
				left++
				right--
			}
		} 
	}
	return data
}

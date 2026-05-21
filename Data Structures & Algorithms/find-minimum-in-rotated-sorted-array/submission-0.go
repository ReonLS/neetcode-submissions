func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	min := nums[0]

	for left <= right{	
		mid := (left + right) / 2

		//update min
		if nums[mid] < min{
			min = nums[mid]
		}

		//cek segmentasi sorted ada di kiri or kanan, berarti opposite of normal binary search
		if nums[mid] > nums[right]{
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return min
}

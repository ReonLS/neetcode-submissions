func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	//meaning alr found peak, if condition invalid
	for left < right{
		mid := (left + right) / 2

		//if mid > right of mid, then peak could be mid or left of mid
		if mid == len(nums)-1 || nums[mid] > nums[mid + 1]{
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

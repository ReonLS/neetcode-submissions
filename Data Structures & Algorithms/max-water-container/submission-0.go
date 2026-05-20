func maxArea(heights []int) int {
	max := 0
	left, right := 0, len(heights)-1

	for left < right{
		compare := min(heights[left], heights[right]) * (right-left)

		if max < compare{
			max = compare
		}

		if heights[right] > heights[left] {
			left++
		} else {
			right--
		}
	}
	return max
}

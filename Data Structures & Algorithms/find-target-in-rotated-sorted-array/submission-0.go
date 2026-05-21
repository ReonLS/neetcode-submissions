func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right{
		middle := (left + right) / 2

		//cek middle jawaban or no
		if nums[middle] == target{
			return middle
		}

		//sortednya left
		if nums[middle] >= nums[left]{

			//cek target diantara sorted left gk
			if target >= nums[left] && target <= nums[middle]{
				right = middle - 1
			// target di sorted right	
			} else {
				left = middle + 1
			}

		//sortednya right
		} else if nums[middle] <= nums[right]{
			if target >= nums[middle] && target <= nums[right]{
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return - 1
}

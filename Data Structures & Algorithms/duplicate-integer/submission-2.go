func hasDuplicate(nums []int) bool {
	//hashmap dengan value sebagai key, ngecek key udh exist gk
	//exist lgsg return true
    hash := make(map[int]bool)

	for _ ,each := range nums{
		//exist ato engga
		_, isExist := hash[each]

		if isExist {
			return true
		}

		//berarti unique value
		hash[each] = true
	}
	return false
}

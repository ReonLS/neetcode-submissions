func getConcatenation(nums []int) []int {
	//buat array baru untuk tampung concat
	//masukin array1 ke newArr
	//append array2 ke newArr

	//(typedata, size awal, capacity)
	newA := make([]int, len(nums))

	//return number of value copied, it this case we dont need
	_ = copy(newA, nums)
	newA = append(newA, nums...)

	return newA
}

func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int)
	left,count,max := 0,0,0

	//iterate array
	for _, char := range s{
		hash[char]++

		//loop while count freq > 1
		for hash[char] > 1{
			hash[rune(s[left])]--
			left++
			count--
		}
		count++

		if count > max{
			max = count
		}
	}
	return max
}

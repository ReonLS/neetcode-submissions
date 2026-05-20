func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int)
	left, max := 0,0

	//iterate array
	for idx, char := range s{
		hash[char]++

		//loop while count freq > 1
		for hash[char] > 1{
			hash[rune(s[left])]--
			left++
		}

		//check current window size, set as max if bigger
		if idx-left+1 > max{
			max = idx-left+1
		}
	}
	return max
}

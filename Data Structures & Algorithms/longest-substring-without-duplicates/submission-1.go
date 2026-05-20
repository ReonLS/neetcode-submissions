func lengthOfLongestSubstring(s string) int {
	unique := make(map[rune]int)
	count, left, max := 0,0,0

	for _, char := range s{
		unique[char]++
		count++

		for unique[char] > 1{
			leftChar := s[left]
			unique[rune(leftChar)]--
			count--
			left++
		}

		if count > max{
			max = count
		}
	}
	return max
}

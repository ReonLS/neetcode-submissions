func lengthOfLongestSubstring(s string) int {
	unique := make(map[rune]int)
	count, left, max := 0,0,0

	for _, char := range s{
		unique[char]++

		if unique[char] == 1{
			count++
		}

		for unique[char] > 1{
			leftChar := s[left]
			unique[leftchar]--
			left++
		}

		if count > max{
			max = count
		}
	}
	return max
}

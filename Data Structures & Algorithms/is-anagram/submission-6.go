func isAnagram(s string, t string) bool {
	freq := make(map[rune]int)
	unique := []rune{}

	if len(s) != len(t){
		return false
	}

	for _, char := range s{
		freq[char]++

		if freq[char] == 1{
			unique = append(unique, char)
		}
	}

	for _, char := range t{
		freq[char]--
	}

	for _, char := range unique{
		if freq[char] != 0{
			return false
		}
	}

	return true
}

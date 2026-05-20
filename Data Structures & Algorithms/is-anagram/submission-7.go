func isAnagram(s string, t string) bool {
	freq := make(map[rune]int)

	if len(s) != len(t){
		return false
	}

	for idx , char := range s{
		freq[char]++
		freq[rune(t[idx])]--
	}

	for _, value := range freq{
		if value != 0{
			return false
		}
	}

	return true
}

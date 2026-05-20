func isPalindrome(s string) bool {
	var cleaned []rune

	for _, alphanum := range s{
		if unicode.IsLetter(alphanum) || unicode.IsDigit(alphanum){
			cleaned = append(cleaned, unicode.ToLower(alphanum))
		}
	}

	left := 0
	right := len(cleaned)-1

	for left < right {
		if cleaned[left] != cleaned[right]{
			return false
		}
		left++
		right--
	}
	return true
}

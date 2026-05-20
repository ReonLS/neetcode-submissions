func isPalindrome(s string) bool {
	var cleaned []rune
	//remove all non alphanumeric char
	for _, char := range s{
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	//two pointer
	left := 0
	right := len(cleaned)-1

	//selama index kanan lebih gede, jadi gk nabrak
	for left < right {
		if cleaned[left] == cleaned[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

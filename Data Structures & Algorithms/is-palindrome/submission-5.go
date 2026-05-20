func isPalindrome(s string) bool {
    var cleaned []rune

	for _, char := range s{
		if unicode.IsLetter(char) || unicode.IsDigit(char){
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}

	if len(cleaned) == 0{
		return true
	}

	left, right := 0, len(cleaned)-1

	for left < right{
		if cleaned[left] != cleaned[right]{
			return false
		}
		left++
		right--
	}
	return true
}

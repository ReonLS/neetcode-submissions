func isPalindrome(s string) bool {
    reg := regexp.MustCompile("[^A-Za-z0-9]")
    cleaned := strings.ToLower(reg.ReplaceAllString(s, ""))

    if len(cleaned) == 0 {
        return true
    }

    left, right := 0, len(cleaned)-1
    for left < right {
        if cleaned[left] != cleaned[right] {
            return false
        }
        left++
        right--
    }
    return true
}
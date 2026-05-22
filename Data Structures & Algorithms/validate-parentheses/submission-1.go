func isValid(s string) bool {
    //waktu masukin closed bracket, di stack harus ada open bracket otherwise false
	//waktu masukin open bracket, lgsg append aja

	stack := []rune{}
	hash := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	for _, char := range s{

		//char == closed bracket
		if _, ok := hash[char]; ok{

			//but stack empty
			if len(stack) == 0{
				return false

			//stack not empty
			} else {

				//top == open bracket equiv
				if stack[len(stack)-1] == hash[char]{
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}

		//else append
		} else {
			stack = append(stack, char)
		}
	}

	//stack should be empty to return true
	if len(stack) == 0{
		return true
	} else {
		return false
	}
}

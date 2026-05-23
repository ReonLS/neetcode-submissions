func isValidSudoku(board [][]byte) bool {
	//track seen digit with true false, store them in hashmap that track row,col, row and col

	rowHash := make(map[int]map[byte]bool)
	colHash := make(map[int]map[byte]bool)
	rowAndColHash := make(map[string]map[byte]bool)

	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			char := board[i][j]

			//check digit or .
			if char == '.'{
				continue
			}

			row := i/3
			col := j/3
			temp := []string{strconv.Itoa(row), strconv.Itoa(col)}
			key := strings.Join(temp, "")

			if rowHash[i] == nil{
				rowHash[i] = make(map[byte]bool)
			}
			if colHash[j] == nil{
				colHash[j] = make(map[byte]bool)
			}
			if rowAndColHash[key] == nil{
				rowAndColHash[key] = make(map[byte]bool)
			}

			if rowHash[i][char] || colHash[j][char] || rowAndColHash[key][char]{
				return false
			} else {
				rowHash[i][char] = true
				colHash[j][char] = true
				rowAndColHash[key][char] = true
			}
		}
	}
	return true
}

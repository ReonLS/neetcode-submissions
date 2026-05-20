func groupAnagrams(strs []string) [][]string {
    var data [][]string
    occur := make(map[int]bool)
 
    for i, row := range strs{
        //mastiin kereset setiap poin elemen
        var indexA [26]int
		var tampungan []string
 
        //uniqueness validation, means key was not found in map
        if !occur[i] {
            occur[i] = true
            tampungan = append(tampungan, row)
        } else {
            //ketemu value udh dimap, lgsg skip
            continue
        }
 
        //count character index
        for j := 0; j < len(row); j++{
            //inc count character element "row" (row elemen strs)
            indexA[row[j] - 'a']++
        }
 
        //pointer ke kata setelah i, untuk bandingin yg anagram
        for k_idx, point := range strs[i+1:]{
            //simple validation
            if len(point) != len(row) { continue }
			indexB := indexA
           
            //minus ind count of char compared to strs[i]
            for k := 0; k < len(point); k++{
                indexB[point[k] - 'a']--
            }
           
            //bandingin count, kalo sama berarti anagram
            isAnagram := true
            for ind := 0; ind < 26; ind++{
				if indexB[ind] != 0{
					isAnagram = false
					break
				}
            }
			if isAnagram {
				occur[i+1+k_idx] = true
				tampungan =  append(tampungan, point)
			}
        }
		data = append (data,tampungan)
    }
    return data
}
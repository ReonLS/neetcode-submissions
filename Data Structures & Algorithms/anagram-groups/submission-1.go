func groupAnagrams(strs []string) [][]string {
	occur := make(map[string][]string)
	var data [][]string

	for _, unsorted := range strs {
		chars := strings.Split(unsorted, "")
		sort.Strings(chars)
		sorted := strings.Join(chars, "")

		occur[sorted] = append(occur[sorted], unsorted)
	}
	
	for _, value := range occur{
		data = append(data, value)
	}
	return data
}

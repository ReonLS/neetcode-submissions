func groupAnagrams(strs []string) [][]string {
	freq := make(map[string][]string)
	var data [][]string

	for _, unsorted := range strs{
		chars := strings.Split(unsorted, "")
		sort.Strings(chars)
		sorted := strings.Join(chars, "")

		freq[sorted] = append(freq[sorted], unsorted)
	}

	for _,value := range freq{
		data = append(data, value)
	}
	return data
}

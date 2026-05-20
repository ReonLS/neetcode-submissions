func groupAnagrams(strs []string) [][]string {
    freq := make(map[string][]string)
    final := [][]string{}

    for _, word := range strs{
        unsorted := strings.Split(word, "")
        sort.Strings(unsorted)
        sorted := strings.Join(unsorted, "")

        freq[sorted] = append(freq[sorted], word)
    }

    for _, sublist := range freq{
        final = append(final, sublist)
    }

    return final
}

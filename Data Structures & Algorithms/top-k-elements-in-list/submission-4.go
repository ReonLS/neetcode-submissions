func topKFrequent(nums []int, k int) []int {
    //bucketsort
    freq := make(map[int]int)
    inv := make(map[int][]int)
    final := []int{}

    for _, num := range nums{
        freq[num]++
    }

    for elm, occ := range freq{
        inv[occ] = append(inv[occ], elm)
    }

    for i := len(nums); i > 0; i--{
        for _, elm := range inv[i]{
            final = append(final, elm)

            if len(final) == k {
                return final
            }
        }
    }
    return final
}

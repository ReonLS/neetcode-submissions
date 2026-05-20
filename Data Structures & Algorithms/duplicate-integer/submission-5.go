func hasDuplicate(nums []int) bool {
  freq := make(map[int]int)

  for _, num := range nums{
    freq[num]++

    if freq[num] > 1{
        return true
    }
  }
  return false
}

func twoSum(nums []int, target int) []int {
   freq := make(map[int]int)
   final := []int{}

   for idx, num := range nums{
    complement := target - num

    if _, ok := freq[complement]; ok{
        final = append(final, freq[complement], idx)
        return final
    } else{
        freq[num] = idx
    }
   }
   return final
}

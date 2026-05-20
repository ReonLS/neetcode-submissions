func twoSum(nums []int, target int) []int {
    //value sebagai key, index sebagai value
    check := make(map[int]int)
    var final []int

    for i, point := range nums{
        needed := target - point

        //check value sebagai key exist or no
        value, isExist := check[needed]
        
        if isExist{
            final = []int{value, i}
            return final
        }
        //key masih unique, input ke map
        check[point] = i
    }
    return final
}
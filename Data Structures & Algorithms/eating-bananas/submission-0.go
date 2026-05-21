func minEatingSpeed(piles []int, h int) int {
	//going through input to get max array, doing binary on min - max speed
    max := 0

    //getting max speed
    for _, num := range piles{
        if num > max{
            max = num
        }
    }
    final := max
    left, right := 1, max

    //divide each pile with potential speed, middle, and comparing to h
    for left <= right{
        mid := (left + right) / 2
        var tempMin int

        for _, nums := range piles{
             tempMin += int(math.Ceil(float64(nums)/float64(mid)))
        }

        if tempMin > h {
            left = mid + 1
        } else {
            right = mid - 1
        }

        //update min
        if mid < final && tempMin <= h{
            final = mid
        }
    }
    return final
}

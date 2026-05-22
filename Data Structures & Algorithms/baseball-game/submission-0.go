func calPoints(operations []string) int {
	list := []int{}
	var sum int

	for _, char := range operations{
		if char == "+"{
			list = append(list, (list[len(list)-2] + list[len(list)-1]))
		} else if char == "D"{
			list = append(list, (list[len(list)-1] * 2))
		} else if char == "C"{
			list = list[:len(list)-1]
		} else {
			val, _ := strconv.Atoi(char)
			list = append(list, val)
		}
	}

	for _, num := range list{
		sum += num
	}
	return sum
}

func singleNumber(nums []int) int {
	mapInt := make(map[int]int, 0)
	for _, v := range nums {
		if _, valid := mapInt[v]; !valid {
			mapInt[v] = v
		} else {
			delete(mapInt, v)
		}
	}
	first := 0
	for k, _ := range mapInt {
		first = k
	}

	return first
}
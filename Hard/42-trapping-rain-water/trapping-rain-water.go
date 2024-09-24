func trap(height []int) int {
	sum := 0
	maxValue := 0
	auxSum := 0
	auxI := 0
	posMax := 0
	for pos, h := range height {
		if h >= maxValue {
			if auxSum > 0 {
				sum = sum + auxSum
				auxI = 0
				auxSum = 0
			}
			maxValue = h
			posMax = pos
		} else {
			auxI++
			auxSum = auxSum + (maxValue - h)
		}
	}
	if height[len(height)-1] <= maxValue {
		maxValue = 0
		auxI = 0
		auxSum = 0
		for i := len(height) - 1; i >= posMax; i-- {
			if height[i] > maxValue {
				if auxSum >= 0 {
					sum = sum + auxSum
					auxI = 0
					auxSum = 0
				}
				maxValue = height[i]
			} else {
				auxI++
				auxSum = auxSum + (maxValue - height[i])
			}
		}
	}

	return sum
}

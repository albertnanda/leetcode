func maxArea(height []int) int {
	var maxArea, currentArea int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] > height[j] {
			currentArea = height[j] * (j - i)
			j--
		} else {
			currentArea = height[i] * (j - i)
			i++
		}
		if currentArea > maxArea {
			maxArea = currentArea
		}

	}

	return maxArea
}

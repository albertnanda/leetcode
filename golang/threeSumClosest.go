func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	totalNumbers := len(nums)
	minDifference := abs(nums[0] + nums[1] + nums[2] - target)
	for currentIndex := 0; currentIndex < totalNumbers-2; currentIndex++ {
		for left, right := currentIndex+1, totalNumbers-1; left < right; {
			currentSum := nums[left] + nums[right] + nums[currentIndex]
			currentDiff := (currentSum - target)
			if currentDiff == 0 {
				return target
			}
			if minDifference > abs(currentSum-target) {
				minDifference = abs(currentSum - target)
				result = currentSum
			}
			if currentDiff < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

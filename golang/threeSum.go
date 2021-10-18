func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	totalNumbers := len(nums)
	for currentIndex := 0; currentIndex < totalNumbers-2; currentIndex++ {
		left := currentIndex + 1
		right := totalNumbers - 1
		if (currentIndex > 0) && nums[currentIndex] == nums[currentIndex-1] {
			continue
		}
		for left < right {
			if current_sum := nums[left] + nums[right] + nums[currentIndex]; current_sum == 0 {
				res = append(res, []int{nums[left], nums[currentIndex], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if current_sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

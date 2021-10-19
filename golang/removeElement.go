func removeElement(nums []int, val int) int {
	modified_length := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[modified_length] = nums[i]
		modified_length++
	}
	return modified_length

}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	modifiedLength := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[modifiedLength] = nums[i]
			modifiedLength++
		}
	}
	return modifiedLength

}

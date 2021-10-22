func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	backtrack(nums, []int{}, &result)
	return result

}
func backtrack(nums []int, currentStack []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, append([]int{}, currentStack...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		backtrack(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(currentStack, nums[i]), result)
	}

}

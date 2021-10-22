func firstMissingPositive(nums []int) int {
	seen := make(map[int]bool)
	for _, value := range nums {
		if value < 1 || value > len(nums) {
			continue
		}
		seen[value] = true
	}
	var i int
	for i = 0; i < len(nums); i++ {
		if _, ok := seen[i+1]; !ok {
			return i + 1
		}
	}
	return i + 1
}

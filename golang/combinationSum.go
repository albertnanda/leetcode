func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, []int{}, &result, 0)
	return result
}
func backtrack(candidates []int, target int, currentStack []int, result *[][]int, idx int) {
	if target == 0 {
		*result = append(*result, append([]int{}, currentStack...))
		return
	}
	if idx >= len(candidates) || target < candidates[idx] {
		return
	}
	backtrack(candidates, target, currentStack, result, idx+1)
	backtrack(candidates, target-candidates[idx], append(currentStack, candidates[idx]), result, idx)

}

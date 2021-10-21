func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}
func backtrack(candidates []int, target int, idx int, currentStack []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, currentStack...))
		return
	}
	if idx == len(candidates) || target < candidates[idx] {
		return
	}
	backtrack(candidates, target-candidates[idx], idx+1, append(currentStack, candidates[idx]), result)
	idx++
	for x, k := candidates[idx-1], idx; k < len(candidates); k++ {
		if candidates[k] != x {
			idx = k
			break
		} else {
			idx++
		}
	}
	backtrack(candidates, target, idx, currentStack, result)
}

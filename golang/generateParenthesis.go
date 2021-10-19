func generateParenthesis(n int) []string {
	result := make([]string, 0)
	return backtrack(result, "", 0, 0, n)
}

func backtrack(result []string, current_string string, left int, right int, max int) []string {
	if len(current_string) == 2*max {
		result = append(result, current_string)
		return result
	}
	if left < max {
		result = backtrack(result, current_string+"(", left+1, right, max)
	}
	if right < left {
		result = backtrack(result, current_string+")", left, right+1, max)
	}
	return result
}

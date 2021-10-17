func longestPalindrome(s string) string {
	var start, end, maxLen int
	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		if len1 > len2 {
			maxLen = len1
		} else {
			maxLen = len2
		}
		if maxLen > (end - start) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}

	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		right++
		left--
	}
	return right - left - 1

}

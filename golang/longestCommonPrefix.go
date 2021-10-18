func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for strIndex := 1; strIndex < len(strs); strIndex++ {
		if prefix == "" {
			break
		}
		for charIndex, value := range prefix {
            if charIndex == len(strs[strIndex]) {
                prefix = prefix[:charIndex]
                break
            }
			if value == rune(strs[strIndex][charIndex]) {
				continue
			}
			prefix = prefix[:charIndex]
			break
		}
	}
	return prefix

}

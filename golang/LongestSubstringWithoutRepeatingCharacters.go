func lengthOfLongestSubstring(s string) int {
	var currentLength, maxLength int
	currentStartIndex := -1
	var seenChar [128]int
	for i := range seenChar {
		seenChar[i] = -1
	}
	for index, value := range s {
		if charSeenIndex := seenChar[value]; charSeenIndex > currentStartIndex {
			currentStartIndex = charSeenIndex
		}
		seenChar[value] = index
		currentLength = index - currentStartIndex
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

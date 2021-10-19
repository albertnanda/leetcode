func strStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < lenHaystack; i++ {
		if haystack[i] == needle[0] {
			if i+lenNeedle > lenHaystack {
				return -1
			} else if haystack[i:i+lenNeedle] == needle {
				return i
			}
		}
	}
	return -1

}

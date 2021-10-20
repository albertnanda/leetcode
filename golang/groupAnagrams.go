func makeKey(s string) string {
	result := make([]byte, 26)
	for _, char := range s {
		result[char-97] += byte(char)
	}
	return string(result)
}
func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		key := makeKey(str)
		dict[key] = append(dict[key], str)
	}
	result := make([][]string, len(dict))
	i := 0
	for _, v := range dict {
		result[i] = v
		i++

	}
	return result
}

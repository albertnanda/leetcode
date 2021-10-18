var table = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
    
    if digits == "" {
		return []string{}
	}
	result := []string{""}

	for _, digit := range digits {
		temp := []string{}
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(table[int(digit)-48]); k++ {
				temp = append(temp, result[j]+string(table[int(digit)-48][k]))

			}
		}
		result = temp
	}

	return result
}

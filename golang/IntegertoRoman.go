func intToRoman(num int) string {
	var result string
	roman := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	for flag := 0; num > 0; flag, num = flag+1, num/10 {
		result = roman[flag][num%10] + result
	}
	return result
}

func myAtoi(s string) int {
	var seen bool
	var result int
	sign := 1
	for _, char := range s {
		if char == ' ' {
			if seen {
				break
			} else {
				continue
			}
		}
		if char == '+' {
			if seen {
				break
			} else {
				sign = 1
				seen = true
				continue
			}
		}
		if char == '-' {
			if seen {
				break
			} else {
				sign = -1
				seen = true
				continue
			}
		}
		seen = true
		if char >= '0' && char <= '9' {
			if (sign == 1) && ((result > math.MaxInt32/10) || (result == math.MaxInt32/10 && (int(char)-48) > 7)) {
				return math.MaxInt32
			}

			if (sign == -1) && ((result*sign < math.MinInt32/10) || (result*sign == math.MinInt32/10 && (int(char)-48) >= 8)) {
				return math.MinInt32
			}
			result = result*10 + (int(char) - 48)
		} else {
			break
		}
	}
	result *= sign
	return result

}

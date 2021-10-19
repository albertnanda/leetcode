func abs(a int) (int, int) {
	if a >= 0 {
		return a, 1
	}
	return -a, -1
}
func divide(dividend int, divisor int) int {
	var quotient int
	dividend, nSign := abs(dividend)
	divisor, dSign := abs(divisor)
	for i := 31; i >= 0; i-- {
		if divisor<<i <= dividend {
			dividend -= divisor << i
			quotient += 1 << i
		}
	}
	quotient = quotient * nSign * dSign
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	return quotient
}

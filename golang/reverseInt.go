func reverse(x int) int {
	var temp, rev int

	for x != 0 {
		temp = x % 10
		if (rev > math.MaxInt32/10) || (rev == math.MaxInt32/10 && temp > 7) {
			return 0
		}

		if (rev < math.MinInt32/10) || (rev == math.MinInt32/10 && temp < -8) {
			return 0
		}
		rev = rev*10 + temp
		x /= 10
	}
	return rev
}

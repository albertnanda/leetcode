func isPalindrome(x int) bool {
	var rev int
	if (x < 0) || (x%10 == 0 && x != 0) {
		return false
	}
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10

}

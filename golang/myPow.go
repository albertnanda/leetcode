func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1.0
	case n == 1:
		return x
	case x == 1.0:
		return 1.0
	case n < 0:
		return myPow(1/x, -n)
	case n%2 == 0:
		return myPow(x*x, n/2)
	default:
		return x * myPow(x*x, (n-1)/2)
	}
}

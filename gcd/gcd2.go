package gcd

func GCD2(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

package gcd

func GCD1(a, b int64) int64 {
	if b == 0 {
		return a
	} else {
		return GCD1(b, a%b)
	}
}

package main

import (
	"fmt"
	"greatest-common-divisor/gcd"
)

const useRecursive = true

func main() {
	var a, b int64 = 30, 12
	fmt.Printf("a = %d, b = %d\n", a, b)

	var fn func(int64, int64) int64
	if useRecursive {
		fn = gcd.GCD1
	} else {
		fn = gcd.GCD2
	}
	g := fn(a, b)
	fmt.Printf("gcd = %d\n", g)
}

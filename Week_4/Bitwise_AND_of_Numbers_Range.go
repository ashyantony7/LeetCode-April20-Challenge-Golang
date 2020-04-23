package main

func rangeBitwiseAnd(m int, n int) int {

	shifts := 0
	for m != n {
		m >>= 1
		n >>= 1
		shifts++
	}

	return m << shifts
}

package main

import "fmt"

func Sqrt(x float64) float64 {
	const maxIterations = 100
	var lo, hi, m float64 = 0, x, x

	for i := 0; i < maxIterations; i++ {
		m = (lo + hi) / 2
		x_ := m * m

		switch {
		case x_ == x:
			return m
		case x_ < x:
			lo = m
		default:
			hi = m
		}
	}

	return m
}

func main() {
	for x := 0; x < 144; x++ {
		fmt.Printf("sqrt(%v) -> %v\n", x, Sqrt(float64(x)))
	}
}

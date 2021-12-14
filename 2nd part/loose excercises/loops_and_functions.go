package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x
	var prev float64

	for prev != z {
		fmt.Println(z)
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(1))
}

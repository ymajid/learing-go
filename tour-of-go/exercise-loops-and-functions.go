package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	lastZ := 0.0

	for math.Abs(z - lastZ) > 1e-8 {
		lastZ = z
		z -= (z*z - x) / (2 * z)
		}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

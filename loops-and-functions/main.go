/*
Computers typically compute the square root of x using a loop.
Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)

Repeating this adjustment makes the guess better and better
until we reach an answer that is as close to the actual square root as can be.
*/
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for oldZ := 0.0; math.Abs(z-oldZ) > 1e-15; {
		oldZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}

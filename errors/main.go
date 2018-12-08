/*
Sqrt should return a non-nil error value when given a negative number,
as it doesn't support complex numbers.
*/
package main

import (
	"fmt"
	"math"
)

type MyError string

// Alternative way that can return any error message
func (e MyError) Error() string {
	return string(e) // can not write: return e
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
		//return x, MyError("cannot Sqrt negative number")
	}
	z := 1.0
	for oldZ := 0.0; math.Abs(z-oldZ) > 1e-15; {
		oldZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

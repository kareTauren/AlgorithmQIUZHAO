package practice

import (
	"fmt"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	n := 2.0
	k := 2.0
	pow := math.Pow(n, k)
	fmt.Println(pow)
	fmt.Println()

	rs := myPow(n, 2.0)
	fmt.Println("myPow ", rs)

	divide := myPowDivide(n, 3.0)
	fmt.Println("myPowDivide ", divide)
}

func myPow(x float64, n int) float64 {
	rsFunc := func(x float64, n int) float64 {
		ans := 1.0
		for i := 0; i < n; i++ {
			ans *= x
		}
		return ans
	}

	if n < 0 {
		return 1 / rsFunc(x, -n)
	}
	return rsFunc(x, n)
}

// Divide
func myPowDivide(x float64, n int) float64 {
	if n < 0 {
		return 1 / divideFunc(x, -n)
	}
	return divideFunc(x, n)
}

func divideFunc(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ans := divideFunc(x, n/2)
	if n&1 == 0 {
		return ans * ans
	}
	return ans * ans * x
}

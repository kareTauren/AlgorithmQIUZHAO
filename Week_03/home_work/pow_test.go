package home_work

import (
	"fmt"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(.00001, 2147483647))
	// fmt.Println(myPow(2.0, -2))
	fmt.Println(math.Pow(0.00001, 2147483647))

	fmt.Println(myPow2(0.00001, 2147483647))
}

func myPow(x float64, n int) float64 {
	call := func(x float64, n int) float64 {
		ans := 1.0
		for i := 0; i < n; i++ {
			ans *= x
		}
		return ans
	}

	if n < 0 {
		return 1 / call(x, -n)
	}
	return call(x, n)
}

func myPow2(x float64, n int) float64 {
	if n < 0 {
		return 1 / myDivide(x, -n)
	}
	return myDivide(x, n)
}

func myDivide(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ans := myDivide(x, n/2)
	if n&1 == 0 {
		return ans * ans
	}
	return ans * ans * x
}

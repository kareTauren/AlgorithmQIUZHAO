package climbStairs

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 2; i < n; i++ {
		tmp := a
		a = b
		b = a + tmp
	}
	return b
}

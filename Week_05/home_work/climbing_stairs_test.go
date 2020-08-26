package home_work

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairsRecursion(5))
	fmt.Println(climbStairsMemo(5))
	fmt.Println(climbStairsTop(5))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairsRecursion(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairsRecursion(n-1) + climbStairsRecursion(n-2)
}

func climbStairsMemo(n int) int {
	memo := make([]int, n+1)
	return helper(n, memo)
}

func helper(n int, memo []int) int {
	if n <= 2 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = helper(n-1, memo) + helper(n-2, memo)
	}
	return memo[n]
}

func climbStairsTop(n int) int {
	if n <= 2 {
		return n
	}
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

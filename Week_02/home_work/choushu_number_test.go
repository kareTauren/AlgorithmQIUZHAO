package home_work

import (
	"fmt"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(10))
}

// https://leetcode-cn.com/problems/chou-shu-lcof/solution/mian-shi-ti-49-chou-shu-dong-tai-gui-hua-qing-xi-t/
/**
试用 DP 方程搞定
*/
func nthUglyNumber(n int) int {
	minFunc := func(n, m int) int {
		if n > m {
			return m
		}
		return n
	}

	var (
		a = 0
		b = 0
		c = 0
	)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1 := dp[a] * 2
		n2 := dp[b] * 3
		n3 := dp[c] * 5
		dp[i] = minFunc(minFunc(n1, n2), n3)
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

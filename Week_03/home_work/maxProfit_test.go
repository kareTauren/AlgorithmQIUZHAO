package home_work

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {

	max := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		// diff := prices[i+1] - prices[i]
		// if diff > 0 {
		// 	ans += diff
		// }
		ans += max(prices[i+1]-prices[i], 0)
	}
	return ans
}

package home_work

import (
	"fmt"
	"strconv"
	"testing"
)

// 先判断一位的情况
// 若当前是 i，则 1<= s[i-1] <=9,那么则有 dp[i] = dp[i] + dp[i-1]，因为不大于 10 就只有一个中解码方式。
//
// 再判断二位的情况
// 若当前是 i, 则 10<= s[i-2] <=26,那么则有 dp[i] = dp[i] + dp[i-2]，因为如果大于 10 小于 26，则不止一种解法。

func TestNumDecodings(t *testing.T) {
	fmt.Println(numDecodings("9"))
}

func numDecodings(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		frist, _ := strconv.Atoi(s[i-1 : i])
		second, _ := strconv.Atoi(s[i-2 : i])
		if frist >= 1 && frist <= 9 {
			dp[i] += dp[i-1]
		}
		if second >= 10 && second <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

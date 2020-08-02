package practice

import (
	"fmt"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	dfs("", n, n, &ans)
	return ans
}

func dfs(tempStr string, l int, r int, ans *[]string) {
	if l == 0 && r == 0 {
		fmt.Println(tempStr)
		*ans = append(*ans, tempStr)
		return
	}

	// 左边括号数不能大于右边括号数量
	if l > r {
		return
	}

	if l > 0 {
		dfs(tempStr+"(", l-1, r, ans)
	}
	if r > 0 {
		dfs(tempStr+")", l, r-1, ans)
	}
}

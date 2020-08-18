package home_work

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
}

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	max := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	stack := make([]int, 0)
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

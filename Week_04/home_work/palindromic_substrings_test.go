package home_work

import (
	"fmt"
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	fmt.Println(countSubstrings("abc"))
}

// 对于每个可能的回文串中心位置，尽可能扩大它的回文区间 [left, right]。
// 当 left >= 0 and right < N and S[left] == S[right] 时，
// 扩大区间。此时回文区间表示的回文串为 S[left], S[left+1], ..., S[right]。
func countSubstrings(s string) int {
	if len(s) == 0 || s == "" {
		return 0
	}

	extendPalindrome := func(s string, left int, right int, count *int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			*count++
		}
	}

	count := 0
	for i := 0; i < len(s); i++ {
		extendPalindrome(s, i, i, &count)
		extendPalindrome(s, i, i+1, &count)
	}
	return count
}

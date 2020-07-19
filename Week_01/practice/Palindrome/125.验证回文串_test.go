package Palindrome

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	// 1. 取出多余的空格，字符 2. 反转字符，忽略大小比较
	clearS := clearStr(s)

	reverseS := reverse(clearS)

	return strings.ToLower(clearS) == strings.ToLower(reverseS)
}

func reverse(s string) string {
	rs := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		rs[i], rs[len(s)-i-1] = rs[len(s)-i-1], rs[i]
	}
	return string(rs)
}

func clearStr(s string) string {
	var rn []rune
	for _, v := range s {
		if isAlnum(v) {
			rn = append(rn, v)
		}
	}
	return string(rn)
}

func isAlnum(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

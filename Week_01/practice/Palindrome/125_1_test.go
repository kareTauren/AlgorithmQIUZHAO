package Palindrome

import (
	"fmt"
	"strings"
	"testing"
)

func TestPalindrome(t *testing.T) {
	fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome2("race a car"))
	fmt.Println(isPalindrome2("我妈和妈，我"))
}

func isPalindrome2(s string) bool {
	var sgood []rune
	for _, v := range s {
		if isAlnum(v) {
			sgood = append(sgood, v)
		}
	}

	for i := 0; i < len(sgood)/2; i++ {
		if strings.ToLower(string(sgood[i])) != strings.ToLower(string(sgood[len(sgood)-1-i])) {
			return false
		}
	}
	return true
}

// func isAlnum(bt byte) bool {
// 	return bt >= 'A' && bt <= 'Z' || bt >= 'a' && bt <= 'z' || bt >= 0 && bt <= 9
// }

package homework

import (
	"fmt"
	"testing"
)

// 两个单词如果包含相同的字母，次序不同，则称为字母易位词(anagram)
func TestIsAnagram(t *testing.T) {
	s := "a"
	s2 := "b"
	fmt.Println(isAnagram(s, s2))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 一个数组里面负责加一，一个负责减一，如果相同的就是++，另一个就-- // 如果小于 0 就证明不是易位词
	cache := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		cache[t[i]-'a']--
		if cache[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

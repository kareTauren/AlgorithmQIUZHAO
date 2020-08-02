package practice

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 || digits == "" {
		return []string{}
	}

	letterMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz"}

	ans := make([]string, 0)
	backTrackStr(digits, &ans, "", letterMap, 0)
	return ans
}

func backTrackStr(str string, ans *[]string, tempStr string, letterMap map[byte]string, index int) {
	if len(tempStr) == len(str) {
		*ans = append(*ans, tempStr)
		return
	}
	mapKey := str[index]
	s := letterMap[mapKey]
	for i := 0; i < len(s); i++ {
		backTrackStr(str, ans, tempStr+string(s[i]), letterMap, index+1)
	}
}

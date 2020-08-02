package home_work

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("29"))
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
	backTrackLetter(digits, &ans, letterMap, 0, "")
	return ans
}

func backTrackLetter(str string, ans *[]string, letterMap map[byte]string, index int, tempAns string) {
	if index == len(str) {
		*ans = append(*ans, tempAns)
		return
	}

	key := str[index]
	s := letterMap[key]
	for i := 0; i < len(s); i++ {
		backTrackLetter(str, ans, letterMap, index+1, tempAns+string(s[i]))
	}
}

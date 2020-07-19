package homework

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 1, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] = digits[i] % 10

		if digits[i] != 0 {
			return digits
		}
	}
	// 不用新开切片，直接把第一位变成 1，后面追加0 即可。
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}

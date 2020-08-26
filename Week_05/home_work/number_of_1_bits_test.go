package home_work

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	// fmt.Println(2 & 1)
	// fmt.Println(1 & 1)
	fmt.Println(hammingWeight(12))
	fmt.Println(hammingWeight2(12))
}

func hammingWeight(num uint32) int {
	var mask uint32 = 1 // 二进制数为 0000001
	count := 0
	for i := 0; i < 32; i++ {
		// 两个数进行位运算，结果只能是 1 或者 0， 如果是1，如果是 1 就证明该位是 1
		if (num & mask) != 0 {
			count++
		}
		mask <<= 1 // 001 => 010 左移一位
	}
	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1)
	}
	return count
}

package home_work

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	// fmt.Printf("%b \n", 10)
	// fmt.Println(1 << 10)
	// fmt.Printf("%b \n", 1<<10)

	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	fmt.Printf("%b \n", num)
	ret := uint32(0)
	power := uint32(31)
	fmt.Printf("power %b \n", power)
	for num != 0 {
		fmt.Println((num & 1))
		fmt.Printf("(num & 1) << power %032b \n", (num&1)<<power)
		ret += (num & 1) << power
		fmt.Println("ret ", ret)
		num = num >> 1
		power -= 1
	}
	fmt.Printf("%b \n", ret)

	return ret
}

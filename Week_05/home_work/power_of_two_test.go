package home_work

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	fmt.Printf("%b \n", 218)
	fmt.Printf("%b \n", 217)
	fmt.Println(isPowerOfTwo(218))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

package home_work

import (
	"fmt"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	fmt.Println(lemonadeChange([]int{10, 10}))
}

// if (customer pays with $5) five++;
// if (customer pays with $10) ten++, five--;
// if (customer pays with $20) ten--, five-- or five -= 3;
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

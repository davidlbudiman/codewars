package katas

import (
	"fmt"
)

func SumDigPow(a, b uint64) (l []uint64) {
	for i := a; i <= b; i++ {
		i_str := fmt.Sprintf("%d", i)
		var sum uint64
		for idx, digit := range i_str {
			var pow_sum = uint64(digit - '0')
			for j := 1; j <= idx; j++ {
				pow_sum = pow_sum * uint64(digit-'0')
			}
			sum = sum + pow_sum
		}
		if uint64(sum) == i {
			l = append(l, i)
		}
	}
	return l
}

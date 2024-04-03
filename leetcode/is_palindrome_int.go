package leetcode

import "fmt"

func IsPalindromeDigit(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	xStr := fmt.Sprintf("%d", x)
	for i, j := 0, len(xStr)-1; i < len(xStr)/2 && j >= len(xStr)/2; i, j = i+1, j-1 {
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}

func IsPalindromeDigit_mondayguy(num int) bool {
	if num < 0 {
		return false
	}
	x := num
	reversed := 0
	for x != 0 {
		reversed = 10*reversed + x%10
		x /= 10
	}
	return (reversed == num)
}
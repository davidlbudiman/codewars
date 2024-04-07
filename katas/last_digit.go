package katas

import (
	"math/big"
	"strconv"
)

func LastDigit_hustlahusky(n1, n2 string) int {
	a, b := big.NewInt(0), big.NewInt(0)
	a.SetString(n1, 10)
	b.SetString(n2, 10)

	a.Exp(a, b, big.NewInt(10))

	return int(a.Int64())
}

func LastDigit_akar_0(n1, n2 string) int {
	max := func(a, b int) int {
		if b > a {
			return b
		} else {
			return a
		}
	}
	pow := func(n, e int) int {
		res := 1
		for i := 0; i < e; i++ {
			res *= n
		}
		return res
	}
	if n2 == "0" {
		return 1
	}
	a, _ := strconv.Atoi(n1[len(n1)-1:])
	b, _ := strconv.Atoi(n2[max(0, len(n2)-2):])
	return pow(a, (b%4)+4) % 10
}

func LastDigit_s1nn0n(n1, n2 string) int {

	var l = map[int][]int{
		0: {0},
		1: {1},
		2: {6, 2, 4, 8},
		3: {1, 3, 9, 7},
		4: {6, 4},
		5: {5},
		6: {6},
		7: {1, 7, 9, 3},
		8: {6, 8, 4, 2},
		9: {1, 9},
	}

	//last digit of n1
	x, _ := strconv.Atoi(n1[len(n1)-1:])

	if len(n2) > 2 {
		n2 = n2[len(n2)-2:]
	}

	y, _ := strconv.Atoi(n2)

	return l[x][int(y)%len(l[x])]

}

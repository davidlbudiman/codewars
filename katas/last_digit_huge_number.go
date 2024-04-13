// For a given list [x1, x2, x3, ..., xn] compute the last (decimal) digit of x1 ^ (x2 ^ (x3 ^ (... ^ xn))).

// E. g., with the input [3, 4, 2], your code should return 1 because 3 ^ (4 ^ 2) = 3 ^ 16 = 43046721.

// Beware: powers grow incredibly fast. For example, 9 ^ (9 ^ 9) has more than 369 millions of digits. lastDigit has to deal with such numbers efficiently.

// Corner cases: we assume that 0 ^ 0 = 1 and that lastDigit of an empty list equals to 1.

// This kata generalizes Last digit of a large number; you may find useful to solve it beforehand.
package katas

func LastDigitHugeNumber(a []int) int {
	if len(a) == 0 {
		return 1
	}
	var l = map[int][]int{
		0: {0},
		1: {1},
		2: {6, 2, 4, 8},
		3: {1, 3, 9, 7},
		4: {6, 4},
		5: {5},
		6: {6},
		7: {9, 3, 1, 7},
		8: {4, 2, 6, 8},
		9: {1, 9},
	}

	for i := len(a) - 2; i > -1; i-- {
		x := a[i] % 10
		if a[i] / 10 > 0 && x == 0 {
			continue
		}
		if a[i+1] == 0 {
			a[i] = 1
		} else {
			y := a[i+1] % len(l[x])
			a[i] = l[x][y]
		}
	}

	return a[0] % 10
}

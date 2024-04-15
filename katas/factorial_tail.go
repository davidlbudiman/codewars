// The problem
// How many zeroes are at the end of the factorial (https://en.wikipedia.org/wiki/Factorial) of 10? 10! = 3628800, i.e. there are 2 zeroes. 16! (or 0x10!) in hexadecimal would be 0x130777758000, which has 3 zeroes.

// Scalability
// Unfortunately, machine integer numbers has not enough precision for larger values. Floating point numbers drop the tail we need. We can fall back to arbitrary-precision ones - built-ins or from a library, but calculating the full product isn't an efficient way to find just the tail of a factorial. Calculating 100'000! in compiled language takes around 10 seconds. 1'000'000! would be around 10 minutes, even using efficient Karatsuba algorithm (https://en.wikipedia.org/wiki/Karatsuba_algorithm)

// Your task
// is to write a function, which will find the number of zeroes at the end of (number) factorial in arbitrary radix = base for larger numbers.

// base is an integer from 2 to 256
// number is an integer from 1 to 1'000'000
// Note Second argument: number is always declared, passed and displayed as a regular decimal number. If you see a test described as 42! in base 20 it's 42(10) not 42(20) = 82(10).
package katas

import (
	"math/big"
)

func Zeroes(base, number int) int {
	res := calculateFactorial(number)
	return getTrailingZeroes(base, res)
}

func calculateFactorial(number int) *big.Int {
	x := big.NewInt(1)
	x.MulRange(1, int64(number))
	return x
}

func getTrailingZeroes(base int, x *big.Int) (numberOfZeroes int) {
	for x.Cmp(&big.Int{}) != 0 {
		y := big.NewInt(int64(base))
		m := &big.Int{}
		x.DivMod(x, y, m)
		if m.Cmp(&big.Int{}) == 0 {
			numberOfZeroes++
		} else {
			break
		}
	}
	return
}

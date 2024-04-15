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

import "strconv"

func Zeroes (base, number int) int {
	res := calculateFactorial(number)
	s := translateToXDecimal(base, res)
	return getTrailingZeroes(s)
}

func calculateFactorial(number int) string {
	return ""
}

func translateToXDecimal(base int, res string) (s string) {
	// max uint64: 18446744073709551615 -> length: 20
	for i := 0; i < len(res); i += 19 {
		x, _ := strconv.ParseUint(res[i:i+20], 10, 64)
	}
	return ""
}

func getTrailingZeroes(s string) (numberOfZeroes int) {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != '0' {
			break
		}
		numberOfZeroes++
	}
	return 0
}

package katas

import (
	"strings"
)

var romans = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
}

var keys = []string{
	"IV",
	"IX",
	"XL",
	"XC",
	"CD",
	"CM",
	"I",
	"V",
	"X",
	"L",
	"C",
	"D",
	"M",
}

func Decode(roman string) (sum int) {
	for _, k := range keys {
		if c := strings.Count(roman, k); c > 0 {
			sum = sum + c*romans[k]
			roman = strings.Replace(roman, k, "", -1)
		}
	}
	return sum
}

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Decode_daddyobrown(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}

func Decode_Mtscream(roman string) int {
	var sum int
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= Roman[byte(v)]
		} else {
			sum += Roman[byte(v)]
		}
	}
	return sum
}

var mapping map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func Decode_devhak2(roman string) int {
	var intern []int
	var result int
	for _, r := range roman {
		intern = append(intern, mapping[r])
	}

	for i := 1; i < len(intern); i++ {
		if intern[i-1] >= intern[i] {
			result += intern[i-1]
		} else {
			result -= intern[i-1]
		}
	}

	result += intern[len(intern)-1]

	return result
}

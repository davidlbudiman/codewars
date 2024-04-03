package leetcode

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	if len(s) > 0 {
		runes := []rune{}
		for _, v := range s {
			if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
				if v >= 'A' && v <= 'Z' {
					v = v + 32
				}
				runes = append(runes, v)
			}
		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			if runes[i] != runes[j] {
				return false
			}
		}
	}
	return true
}

func IsPalindrome_gsalomao(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i = i + 1
		j = j - 1
	}

	return true
}

package katas

import (
	"regexp"
	"strings"
)

func ValidBraces(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	var openers = map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	var lifo []rune
	for _, c := range str {
		if len(lifo) == 0 || c == '{' || c == '[' || c == '(' {
			lifo = append(lifo, c)
		} else if (c == '}' || c == ']' || c == ')') &&
			lifo[len(lifo)-1] == openers[c] {
			lifo = lifo[:len(lifo)-1]
		} else {
			return false
		}
	}
	return len(lifo) == 0
}

func ValidBraces_v1(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	var closer_re = regexp.MustCompile(`[\]})]`)
	var opener_re = regexp.MustCompile(`[\[{(]`)

	var runes = []rune(str)
	var lifo []rune
	for i := 0; i < len(runes); i++ {
		if len(lifo) == 0 || opener_re.MatchString(string(runes[i])) {
			lifo = append(lifo, runes[i])
		} else if closer_re.MatchString(string(runes[i])) {
			var opener rune
			switch runes[i] {
			case '}':
				opener = '{'
			case ')':
				opener = '('
			case ']':
				opener = '['
			default:
				return false
			}
			if lifo[len(lifo)-1] == opener {
				lifo = lifo[:len(lifo)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return len(lifo) == 0
}

func ValidBraces_iGiant(str string) bool {
	stack := make([]rune, 0)
	for _, c := range str {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func ValidBraces_yuhongyuleon(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}

func ValidBraces_Zaphod137(str string) bool {
	lastLength := -1
	length := len(str)

	for lastLength != length {
		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)

		lastLength = length
		length = len(str)
	}
	return length == 0
}

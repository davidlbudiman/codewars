package katas

import (
	"fmt"
	"strings"
)

func Accum(s string) (r string) {
	s = strings.ToUpper(s)
	runes := []rune(s)

	for idx, ru := range runes {
		i := 1
		r = r + fmt.Sprintf("%c", ru)
		lower := fmt.Sprintf("%c", ru+32)
		for ; i <= idx; i++ {
			r = r + lower
		}
		if idx+1 != len(runes) {
			r = r + "-"
		}
	}

	return r
}

func Accum_niosop(s string) string {
    parts := make([]string, len(s))
    
    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }
    
    return strings.Join(parts, "-")
}

package katas

import "fmt"

var mPermutations = map[string]int{}

func Permutations(s string) []string {
	// your code here
	mPermutations = map[string]int{}
	if len(s) == 1 {
		return []string{s}
	}
	permute("", s)
	a := []string{}
	for k := range mPermutations {
		a = append(a, k)
	}
	return a
}

func permute(prefix string, s string) {
	for i := 0; i < len(s); i++ {
		t := s[0:i] + s[i+1:]
		permute(prefix+fmt.Sprintf("%c", s[i]), t)
	}
	if len(s) == 0 {
		mPermutations[prefix] = 1
	}
}

func Permutations_wasmup(s string) (a []string) {
	if len(s) < 2 {
		return []string{s}
	}
	m := map[string]bool{}
	for _, sub := range Permutations_wasmup(s[1:]) {
		for i := 0; i <= len(sub); i++ {
			st := sub[0:i] + s[0:1] + sub[i:]
			if m[st] {
				continue
			}
			m[st] = true
			a = append(a, st)
		}
	}
	return
}

func p(s []rune, l, n int, r *[]string, sn map[string]bool) {
	if !sn[string(s)] {
		*r = append(*r, string(s))
		sn[string(s)] = true
	}
	for i := l; i < n; i++ {
		s[l], s[i] = s[i], s[l]
		p(s, l+1, n, r, sn)
		s[l], s[i] = s[i], s[l]
	}
}

func Permutations_Matbabs(s string) []string {
	var r []string
	sn := make(map[string]bool)
	p([]rune(s), 0, len(s), &r, sn)
	return r
}

package katas_test

import (
	"math/rand"
	"sort"
	"time"

	. "davidbudiman.xyz/codewars/katas"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func referencePermutationsSolution(s string) []string {
	l := len(s)
	switch l {
	case 0:
		return []string{}
	case 1:
		return []string{s}
	}
	m := make(map[string]bool)
	arr, c := referencePermutationsSolution(s[1:]), string(s[0])
	for _, p := range arr {
		for k := 0; k < l; k++ {
			m[p[:k]+c+p[k:]] = true
		}
	}
	o := make([]string, len(m))
	var i int
	for x := range m {
		o[i] = x
		i++
	}
	sort.Strings(o)
	return o
}

func doTest(s string, expected []string) {
	actual := Permutations(s)
	sort.Strings(actual)
	Expect(actual).To(Equal(expected))
}

var _ = Describe("Tests", func() {
	It("Sample tests - Unique letters", func() {
		doTest("a", []string{"a"})
		doTest("ab", []string{"ab", "ba"})
		doTest("abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"})
		abcd := []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
			"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
			"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
			"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}
		doTest("abcd", abcd)
		doTest("bcda", abcd)
		doTest("dcba", abcd)
	})
	It("Sample tests - Duplicate letters", func() {
		doTest("aa", []string{"aa"})
		doTest("aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"})
		doTest("aaaab", []string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"})
		doTest("aaaaab", []string{"aaaaab", "aaaaba", "aaabaa", "aabaaa", "abaaaa", "baaaaa"})
	})
	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 100; i++ {
			l := 1 + rand.Intn(8)
			arr := make([]rune, l)
			for j := 0; j < l; j++ {
				arr[j] = rune(97 + rand.Intn(26))
			}
			s := string(arr)
			expected := referencePermutationsSolution(s)
			doTest(s, expected)
		}
	})
})

package leetcode_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/codewars/leetcode"
)

func canConstructTest(ransomNote, magazine string, expected bool) {
	It(fmt.Sprintf("should return %v for \"%v\"", expected, ransomNote), func() {
		Expect(leetcode.CanConstruct(ransomNote, magazine)).To(Equal(expected))
	})
}

var _ = Describe("Can Construct", func() {
	canConstructTest("aa", "ab", false)
	canConstructTest("ab", "aa", false)
	canConstructTest("ab", "aab", true)
	canConstructTest("aa", "aab", true)
})

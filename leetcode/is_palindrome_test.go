package leetcode_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/leetcode"
)

func isPalindromeTest(str string, expected bool) {
	It(fmt.Sprintf("should return %v for \"%v\"", expected, str), func() {
		Expect(leetcode.IsPalindrome(str)).To(Equal(expected))
	})
}

var _ = Describe("Is Palindrome", func() {
	isPalindromeTest("aaaaa", true)
	isPalindromeTest("race e car", true)
	isPalindromeTest("race a car", false)
	isPalindromeTest("abc,def,fed,CBA", true)
})

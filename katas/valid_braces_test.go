package katas_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/katas"
)

func singleTest(str string, res bool) {
	It(fmt.Sprintf("should return %v for \"%v\"", res, str), func() {
		Expect(katas.ValidBraces(str)).To(Equal(res))
	})
}

var _ = Describe("Valid Braces", func() {
	singleTest("(){}[]", true)
	singleTest("([{}])", true)
	singleTest("(}", false)
	singleTest("[(])", false)
	singleTest("[({)](]", false)
})

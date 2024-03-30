package katas_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/codewars/katas"
)

var _ = Describe("Split Strings", func() {
	It("Basic tests", func() {
		Expect(katas.SplitStrings("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(katas.SplitStrings("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(katas.SplitStrings("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})

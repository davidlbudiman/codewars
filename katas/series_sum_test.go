package katas_test

import (
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"

	"davidbudiman.xyz/katas"
)
var _ = Describe("SeriesSum Tests", func() {
  It("should pass provided tests", func() {
     Expect(katas.SeriesSum(1)).To(Equal("1.00"), "n = 1")
     Expect(katas.SeriesSum(2)).To(Equal("1.25"), "n = 2")
     Expect(katas.SeriesSum(3)).To(Equal("1.39"), "n = 3")
     Expect(katas.SeriesSum(4)).To(Equal("1.49"), "n = 4")
  })
})

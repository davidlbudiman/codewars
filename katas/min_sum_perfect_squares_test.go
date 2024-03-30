package katas_test

import (
  "math/rand"
  "time"
  "math"

  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  . "davidbudiman.xyz/katas"
)


func referenceSolutionSumOfSquares(n uint64) uint64 {
  a := n
  for a % 4 == 0 { a >>= 2}
  if a  % 8 == 7 { return 4 }
  s := uint64(math.Sqrt(float64(n)))
  if n == s * s { return 1 }
  for k := uint64(1) ; k <= s ; k++ {
    a = n - k * k
    s :=  uint64(math.Floor(math.Sqrt(float64(a))))
    if a == s * s { return 2 }
  }
  return 3
}


var _ = Describe("Tests", func() {
     It("Sample tests", func() {
       Expect(SumOfSquares(15)).To(Equal(uint64(4)))
       Expect(SumOfSquares(16)).To(Equal(uint64(1)))
       Expect(SumOfSquares(17)).To(Equal(uint64(2)))
       Expect(SumOfSquares(18)).To(Equal(uint64(2)))
       Expect(SumOfSquares(19)).To(Equal(uint64(3)))
     })
     It("Hard fixed tests", func() {
       Expect(SumOfSquares(2017)).To(Equal(uint64(2)))
       Expect(SumOfSquares(1008)).To(Equal(uint64(4)))
       Expect(SumOfSquares(3456)).To(Equal(uint64(3)))
       Expect(SumOfSquares(4000)).To(Equal(uint64(2)))
       Expect(SumOfSquares(12321)).To(Equal(uint64(1)))
     })
     It("Very hard fixed tests", func() {
       Expect(SumOfSquares(661915703)).To(Equal(uint64(4)))
       Expect(SumOfSquares(999887641)).To(Equal(uint64(1)))
       Expect(SumOfSquares(999950886)).To(Equal(uint64(3)))
       Expect(SumOfSquares(999951173)).To(Equal(uint64(2)))
       Expect(SumOfSquares(999998999)).To(Equal(uint64(4)))
     })
     It("Random tests", func() {
       rand.Seed(time.Now().UTC().UnixNano())
       for i := 0 ; i < 100 ; i++ {
         n := 100000000 + rand.Uint64() % 899999999
         Expect(SumOfSquares(n)).To(Equal(referenceSolutionSumOfSquares(n)))
       }
     })
})
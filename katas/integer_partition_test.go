package katas_test

import (
	"math/rand"

	. "davidbudiman.xyz/codewars/katas"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Cases", func() {
	It("the number of integer partitions of 1 should be correct", func() {
		Expect(Partitions(1)).To(Equal(1))
	})

	It("the number of integer partitions of 2 should be correct", func() {
		Expect(Partitions(2)).To(Equal(2))
	})

	It("the number of integer partitions of 3 should be correct", func() {
		Expect(Partitions(3)).To(Equal(3))
	})

	It("the number of integer partitions of 4 should be correct", func() {
		Expect(Partitions(4)).To(Equal(5))
	})

	It("the number of integer partitions of 5 should be correct", func() {
		Expect(Partitions(5)).To(Equal(7))
	})

	It("the number of integer partitions of 6 should be correct", func() {
		Expect(Partitions(6)).To(Equal(11))
	})

	It("the number of integer partitions of 7 should be correct", func() {
		Expect(Partitions(7)).To(Equal(15))
	})

	It("the number of integer partitions of 10 should be correct", func() {
		Expect(Partitions(10)).To(Equal(42))
	})

	It("the number of integer partitions of 15 should be correct", func() {
		Expect(Partitions(15)).To(Equal(176))
	})

	It("the number of integer partitions of 20 should be correct", func() {
		Expect(Partitions(20)).To(Equal(627))
	})

	It("the number of integer partitions of 30 should be correct", func() {
		Expect(Partitions(30)).To(Equal(5604))
	})

	It("the number of integer partitions of 50 should be correct", func() {
		Expect(Partitions(50)).To(Equal(204226))
	})

	It("the number of integer partitions of 100 should be correct", func() {
		Expect(Partitions(100)).To(Equal(190569292))
	})
})

func partitions(n int, cache map[int]int) int {
	if n <= 1 {
		return 1
	}
	if _, ok := cache[n]; !ok {
		m := n - 1
		size := 0
		step := 2
		for m >= 0 {
			prevSize := partitions(m, cache)
			if step/2%2 == 1 {
				size += prevSize
			} else {
				size -= prevSize
			}
			if step%2 == 1 {
				m -= step
			} else {
				m -= step / 2
			}
			step++
		}
		cache[n] = size
	}
	return cache[n]
}

var _ = Describe("Random Tests", func() {
	cache := make(map[int]int, 99)
	rand.Seed(GinkgoRandomSeed())
	for i := 0; i < 15; i++ {
		n := rand.Intn(80) + 20
		size := partitions(n, cache)
		It("Should work for a random n", func() {
			Expect(Partitions(n)).To(Equal(size))
		})
	}
})

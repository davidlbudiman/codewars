package katas_test

import (
	"math/rand"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/codewars/katas"
)

func pow(n rune, p int) uint64 {
	m := uint64(n - 48)
	res := uint64(1)
	for i := 0; i < p; i++ {
		res *= m
	}
	return res
}

func referenceSolution(a, b uint64) (res []uint64) {
	for n := a; n <= b; n++ {
		var s uint64
		for i, d := range strconv.FormatUint(n, 10) {
			if d != '0' {
				s += pow(d, i+1)
			}
		}
		if s == n {
			res = append(res, n)
		}
	}
	return res
}

func dorandtests(nTests int, minA, maxA, minB, maxB uint64) {
	for i := 0; i < nTests; i++ {
		a := minA + rand.Uint64()%(maxA-minA)
		b := minB + rand.Uint64()%(maxB-minB)
		dotest(a, b, referenceSolution(a, b))
	}
}

func dotest(a, b uint64, expected []uint64) {
	actual := katas.SumDigPow(a, b)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "With a = %d, b = %d", a, b)
	} else {
		Expect(actual).To(Equal(expected), "With a = %d, b = %d", a, b)
	}
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest(1, 10, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		dotest(1, 100, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89})
		dotest(10, 89, []uint64{89})
		dotest(10, 100, []uint64{89})
		dotest(90, 100, nil)
		dotest(90, 150, []uint64{135})
		dotest(50, 150, []uint64{89, 135})
		dotest(10, 150, []uint64{89, 135})
		dotest(89, 135, []uint64{89, 135})
	})
	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		dorandtests(40, 100, 500, 510, 1500)
		dorandtests(40, 10, 1999, 2000, 3000)
		dorandtests(10, 2620000, 2640000, 2647000, 2648000)
		dorandtests(10, 12157692622039623000, 12157692622039625500, 12157692622039625600, 12157692622039625700)
	})
})

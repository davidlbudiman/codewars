package katas_test

import (
	"math/rand"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "davidbudiman.xyz/codewars/katas"
)

func lastDigitPow(n, e int) int {
	res := 1
	for i := 0; i < e; i++ {
		res *= n
	}
	return res
}

func lastDigitReferenceSolution(n1, n2 string) int {
	max := func(a, b int) int {
		if b > a {
			return b
		} else {
			return a
		}
	}
	if n2 == "0" {
		return 1
	}
	a, _ := strconv.Atoi(n1[len(n1)-1:])
	b, _ := strconv.Atoi(n2[max(0, len(n2)-2):])
	return lastDigitPow(a, (b%4)+4) % 10
}

func randLong() string {
	arr := make([]rune, 1+5*rand.Intn(lastDigitPow(10, rand.Intn(6))))
	arr[0] = rune(49 + rand.Intn(9))
	for i := range arr[1:] {
		arr[i+1] = rune(48 + rand.Intn(10))
	}
	return string(arr)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(LastDigit("4", "1")).To(Equal(4))
		Expect(LastDigit("4", "2")).To(Equal(6))
		Expect(LastDigit("9", "7")).To(Equal(9))
		Expect(LastDigit("10", "10000000000")).To(Equal(0))
		Expect(LastDigit("1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376")).To(Equal(6))
		Expect(LastDigit("3715290469715693021198967285016729344580685479654510946723", "68819615221552997273737174557165657483427362207517952651")).To(Equal(7))
		Expect(LastDigit("3715290469715693021198967285016729344580685479654510946723", "0")).To(Equal(1))
	})
	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 100; i++ {
			a, b := randLong(), randLong()
			Expect(LastDigit(a, b)).To(Equal(lastDigitReferenceSolution(a, b)))
		}
	})
})

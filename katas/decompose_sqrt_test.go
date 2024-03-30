package katas_test

import (
	"fmt"

	. "davidbudiman.xyz/katas"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func isInc(arr []int64) int {
	lg := len(arr)
	if lg < 2 {
		return 0
	}
	for i := 0; i < lg-1; i++ {
		if arr[i] >= arr[i+1] {
			return 0
		}
	}
	return 1
}
func total(arr []int64, m int64) int {
	var sum int64 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] * arr[i]
	}
	if sum == m {
		return 1
	}
	return 0
}
func doDecomposeTest(n int64, exp []int64) {
	fmt.Printf("n %v\n", n)
	var ans = Decompose(n)
	var success = false
	fmt.Printf("Actual %v\n", ans)
	fmt.Printf("Possible Expect %v\n", exp)
	if len(ans) == 0 && len(exp) == 0 {
		success = true
	} else {
		st := isInc(ans)
		t := total(ans, n*n)
		if st == 0 || t == 0 {
			println("Not increasinly sorted or bad sum of squares")
			success = false
		} else {
			success = true
		}
	}
	Expect(success).To(Equal(true))
	println("*")
}

var _ = Describe("Tests Decompose", func() {

	It("should handle basic cases", func() {
		doDecomposeTest(50, []int64{1, 3, 5, 8, 49})
		doDecomposeTest(5, []int64{3, 4})
		doDecomposeTest(2, []int64{})
		doDecomposeTest(7, []int64{2, 3, 6})

	})
})

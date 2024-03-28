package leetcode_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"davidbudiman.xyz/codewars/leetcode"
)

func mergeTest(nums1, nums2, expected []int, m, n int) {
	It(fmt.Sprintf("should return %v for \"%v\"", expected, nums1), func() {
		leetcode.Merge(nums1, m, nums2, n)
		Expect(nums1).To(Equal(expected))
	})
}

var _ = Describe("Merge", func() {
	mergeTest([]int{1, 2, 3, 0, 0, 0}, []int{2, 3, 4}, []int{1, 2, 2, 3, 3, 4}, 3, 3)
	mergeTest([]int{1, 2, 3, 0, 0, 0}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, 3, 3)
	mergeTest([]int{0}, []int{1}, []int{1}, 0, 1)
	mergeTest([]int{1}, []int{}, []int{1}, 1, 0)
	mergeTest([]int{0}, []int{}, []int{0}, 1, 0)
})

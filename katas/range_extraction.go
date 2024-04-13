// A format for expressing an ordered list of integers is to use a comma separated list of either

// individual integers
// or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'. The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"
// Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.

// Example:

// solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"
package katas

import (
	"fmt"
	"strings"
)

func RangeExtraction(nums []int) string {
	var res []string
	if len(nums) == 1 {
		res = append(res, createSummary(nums[0], nums[0]))
	} else if len(nums) > 1 {
		var min, max = nums[0], nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] != max+1 {
				res = append(res, createSummary(min, max))
				min, max = nums[i], nums[i]
			} else {
				max = nums[i]
			}
		}
		res = append(res, createSummary(min, max))
	}
	return strings.Join(res, ",")
}

func createSummary(min, max int) (r string) {
	if min == max {
		r = fmt.Sprintf("%d", min)
	} else {
		if max-min == 1 {
			r = fmt.Sprintf("%d,%d", min, max)
		} else {
			r = fmt.Sprintf("%d-%d", min, max)
		}
	}
	return r
}

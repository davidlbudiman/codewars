package leetcode

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) (res []string) {
;

func SummaryRanges_babureddyh(nums []int) []string {
	result := []string{}
	r := ""

	for i := 0; i < len(nums); {
		r = r + strconv.Itoa(nums[i])
		j := i + 1
		for ; j < len(nums) && nums[j]-nums[j-1] == 1; j++ {
		}
		if j-1 > i {
			r = r + "->" + strconv.Itoa(nums[j-1])
		}
		result = append(result, r)
		r = ""
		i = j
	}
	return result
}

func SummaryRanges_escribapetrus(nums []int) []string {
	var answer []string
	var store []int

	if len(nums) == 0 {
		return answer
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1]+1 == nums[i] {
			store = append(store, nums[i])
			continue
		}
		answer = append(answer, make_string(store))
		store = nil
		store = append(store, nums[i])
	}
	return append(answer, make_string(store))
}

func make_string(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	if len(arr) == 1 {
		return fmt.Sprintf("%d", arr[0])
	}
	return fmt.Sprintf("%d->%d", arr[0], arr[len(arr)-1])
}

package leetcode

import "sort"

func SearchInsert(nums []int, target int) int {
	search := func() int {
		mid := len(nums) / 2
		if len(nums) == 1 {
			if target < nums[mid] {
				return -1
			} else if target > nums[mid] {
				return 1
			} else {
				return 0
			}
		}
		var idx int
		if target < nums[mid] {
			idx = SearchInsert(nums[:mid], target)
		} else {
			idx = mid + SearchInsert(nums[mid:], target)
		}
		return idx
	}()
	if search < 0 {
		return 0
	}
	return search
}

func SearchInsert_troll(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

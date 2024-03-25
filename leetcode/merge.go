package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	var nums1_temp = make([]int, m)
	for i, _ := range nums1_temp {
		nums1_temp[i] = nums1[i]
	}
	i, j, k := 0, 0, 0
	for {
		if i == m || j == n {
			break
		}
		if nums1_temp[i] < nums2[j] {
			nums1[k] = nums1_temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	if i < m {
		for {
			nums1[k] = nums1_temp[i]
			i++
			k++
			if i == m {
				break
			}
		}
	}
	if j < n {
		for {
			nums1[k] = nums2[j]
			j++
			k++
			if j == n {
				break
			}
		}
	}
}

// The above algorithm will have O(nlogn) complexity due to the loops to place the numbers in a sorted order takes
// place at least once and at most three times.
// The below solution shows you O(n) complexity, in reality it loops twice. But since we don't put constant in 
// complexity, hence O(n)
func merge_counting_sort(nums1 []int, m int, nums2 []int, n int) {
	var min, max int
	var counting_maps = map[int]int{}
	for i := 0; i < m; i++ {
		if nums1[i] < min {
			min = nums1[i]
		} else if nums1[i] > max {
			max = nums1[i]
		}
		counting_maps[nums1[i]] = counting_maps[nums1[i]] + 1
	}
	for j := 0; j < n; j++ {
		if nums2[j] < min {
			min = nums2[j]
		} else if nums2[j] > max {
			max = nums2[j]
		}
		counting_maps[nums2[j]] = counting_maps[nums2[j]] + 1
	}
	for i, j := 0, min; j <= max; j++ {
		if i == len(nums1) {
			break
		}
		if v, ok := counting_maps[j]; ok {
			for k := v; k > 0; k-- {
				nums1[i] = j
				i++
			}
		}
	}
}

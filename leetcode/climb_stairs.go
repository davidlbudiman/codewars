package leetcode

func ClimbStairs(n int) int {
	var last, next = 0, 1
	for i := 0; i < n; i++ {
		last, next = next, last+next
	}
	return next
}

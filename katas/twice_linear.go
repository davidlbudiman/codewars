// Consider a sequence u where u is defined as follows:

// The number u(0) = 1 is the first one in u.
// For each x in u, then y = 2 * x + 1 and z = 3 * x + 1 must be in u too.
// There are no other numbers in u.
// Ex: u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]

// 1 gives 3 and 4, then 3 gives 7 and 10, 4 gives 9 and 13, then 7 gives 15 and 22 and so on...

// Task:
// Given parameter n the function dbl_linear (or dblLinear...) returns the element u(n) of the ordered (with <) sequence u (so, there are no duplicates).

// Example:
// dbl_linear(10) should return 22

// Note:
// Focus attention on efficiency
package katas

var dblLinears = []int{1}
var linearExist = map[int]bool{1: true}

func DblLinear(n int) int {
	if n >= len(dblLinears) {
		fillDblLinears(n * 3)
	}
	return dblLinears[n]
}

func fillDblLinears(n int) {
	lastElement := (dblLinears[len(dblLinears)-1] - 1) / 3
	index := getElementIndex(lastElement)
	for i := len(dblLinears) - 1; i < n; i += 2 {
		index++
		lastElement = dblLinears[index]
		x := lastElement*2 + 1
		if !linearExist[x] {
			insertElement(x)
			linearExist[x] = true
		}
		y := lastElement*3 + 1
		if !linearExist[y] {
			insertElement(y)
			linearExist[y] = true
		}
	}
}

func getElementIndex(e int) int {
	for i, x := range dblLinears {
		if x == e {
			return i
		}
	}
	return -1
}

func insertElement(e int) {
	if e > dblLinears[len(dblLinears)-1] {
		dblLinears = append(dblLinears, e)
	} else {
		for i, x := range dblLinears[:len(dblLinears)-1] {
			if x < e && e < dblLinears[i+1] {
				dblLinears = append(dblLinears[:i+2], dblLinears[i+1:]...)
				dblLinears[i+1] = e
				break
			}
		}
	}
}

func DblLinear_imsingh(n int) int {

	// Code
	u := []int{1}
	i := 0
	j := 0

	var y int
	var z int

	for len(u) <= n {
		y = 2*u[i] + 1
		z = 3*u[j] + 1

		if y < z {
			u = append(u, y)
			i++
		} else if y > z {
			u = append(u, z)
			j++
		} else {
			u = append(u, y)
			i++
			j++
		}
	}
	return u[len(u)-1]

}

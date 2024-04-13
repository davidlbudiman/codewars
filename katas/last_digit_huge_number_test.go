package katas

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LastDigitHugeNumber(t *testing.T) {
	var r1 int = rand.Intn(100)
	var r2 int = rand.Intn(10)
	var pow int = int(math.Pow(float64(r1%10), float64(r2)))

	tests := []struct {
		name   string
		a      []int
		expect int
	}{
		{ name: "test 1", a: []int{}, expect: 1 },
		{ name: "test 2", a: []int{0, 0}, expect: 1},
		{ name: "test 3", a: []int{0, 0, 0}, expect: 0 },
		{ name: "test 4", a: []int{1, 2}, expect: 1 },
		{ name: "test 5", a: []int{3, 4, 5}, expect: 1 },
		{ name: "test 6", a: []int{4, 3, 6}, expect: 4 },
		{ name: "test 7", a: []int{7, 6, 21}, expect: 1 },
		{ name: "test 8", a: []int{12, 30, 21}, expect: 6 },
		{ name: "test 9", a: []int{2, 0, 1}, expect: 1 },
		{ name: "test 10", a: []int{2, 2, 2, 0}, expect: 4 },
		{ name: "test 11", a: []int{937640, 767456, 981242}, expect: 0 },
		{ name: "test 12", a: []int{123232, 694022, 140249}, expect: 6 },
		{ name: "test 13", a: []int{499942, 898102, 846073}, expect: 6 },
		{ name: "test 1 random", a: []int{r1}, expect: r1 % 10 },
		{ name: "test 2 random", a: []int{r1, r2}, expect: pow % 10 },
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, LastDigitHugeNumber(tt.a))
		})
	}
}

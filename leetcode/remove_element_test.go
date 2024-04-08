package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		val       int
		expect    int
		expectArr []int
	}{
		{
			name:      "test 1",
			nums:      []int{3, 2, 2, 3},
			val:       3,
			expect:    2,
			expectArr: []int{2, 2},
		},
		{
			name:      "test 2",
			nums:      []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:       2,
			expect:    5,
			expectArr: []int{0, 1, 4, 0, 3},
		},
		{
			name:      "test 3",
			nums:      []int{},
			val:       0,
			expect:    0,
			expectArr: []int{},
		},
		{
			name:      "test 4",
			nums:      []int{2},
			val:       3,
			expect:    1,
			expectArr: []int{2},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			l := removeElement(tt.nums, tt.val)
			assert.Equal(t, tt.expect, l)
			assert.Equal(t, tt.expectArr, tt.nums[:l])
		})
	}
}

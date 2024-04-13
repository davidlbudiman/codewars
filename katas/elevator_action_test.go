package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Order(t *testing.T) {
	tests := []struct {
		name   string
		level  int
		queue  []Person
		expect []int
	}{
		{name: "First day", level: 1, queue: []Person{{From: 3, To: 2}, {From: 5, To: 2}, {From: 2, To: 1},
			{From: 2, To: 5}, {From: 4, To: 3}}, expect: []int{2, 5, 4, 3, 2, 1}},
		{name: "Second day", level: 1, queue: []Person{{From: 5, To: 3}, {From: 4, To: 2}}, expect: []int{5, 4, 3, 2}},
		{name: "Third day", level: 4, queue: []Person{{From: 5, To: 3}, {From: 4, To: 2}}, expect: []int{5, 4, 3, 2}},
		{name: "Fourth day", level: 4, queue: []Person{{From: 4, To: 2}, {From: 5, To: 3}}, expect: []int{4, 2, 5, 3}},
		{name: "Fifth day", level: 1, queue: []Person{{From: 5, To: 1}, {From: 4, To: 3}, {From: 2, To: 1}},
			expect: []int{5, 4, 3, 2, 1}},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, Order(tt.level, tt.queue))
		})
	}
}

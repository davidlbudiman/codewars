package katas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroes(t *testing.T) {
	tests := []struct {
		base   int
		number int
		exp    int
	}{
		{base: 10, number: 10, exp: 2},
		{base: 16, number: 16, exp: 3},
		{base: 40, number: 10, exp: 2},
		{base: 12, number: 26, exp: 10},
		{base: 17, number: 100, exp: 5},
		{base: 2, number: 524288, exp: 524287},
	}
	for _, tc := range tests {
		tt := tc
		t.Run("Zeroes", func(t *testing.T) {
			assert.Equal(t, tt.exp, Zeroes(tt.base, tt.number))
		})
	}
}

package katas

import (
	"fmt"
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
	}
	for i, tc := range tests {
		tt := tc
		t.Run(fmt.Sprintf("Zeroes %d", i), func(t *testing.T) {
			assert.Equal(t, tt.exp, Zeroes(tt.base, tt.number))
		})
	}
}

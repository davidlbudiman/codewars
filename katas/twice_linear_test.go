package katas

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwiceLinear(t *testing.T) {
	tests := []struct {
		n   int
		exp int
	}{
		{n: 10, exp: 22},
		{n: 50, exp: 175},
		{n: 100, exp: 447},
		{n: 500, exp: 3355},
		{n: 1000, exp: 8488},
		{n: 60000, exp: 1511311},
	}
	for i, tc := range tests {
		tt := tc
		t.Run(fmt.Sprintf("Twice Linear: %d", i+1), func(t *testing.T) {
			assert.Equal(t, tt.exp, DblLinear(tt.n))
		})
	}
}

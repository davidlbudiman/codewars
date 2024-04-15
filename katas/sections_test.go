package katas

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSections(t *testing.T) {
	tests := []struct{
		k int
		expect int
	}{
		{k: 1, expect: 1},
		{k: 4, expect: 4},
		{k: 423128, expect: 0},
		{k: 1369, expect: 4},
		{k: 2999824, expect: 28},
		{k: 111710084, expect: 64},
		{k: 40965767, expect: 160},
		{k: 2019, expect: 0},
	}
	for i, tc := range tests {
		tt := tc
		t.Run(fmt.Sprintf("Section %d", i + 1), func(t *testing.T) {
			assert.Equal(t, tt.expect, Sections(tt.k))
		})
	}
}

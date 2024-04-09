package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsSubsequence(t *testing.T) {
	tests := []struct{
		name string
		s string
		t string
		expect bool
	}{
		{
			name: "test 1",
			s: "abc",
			t: "ahbgdc",
			expect: true,
		},
		{
			name: "test 2",
			s: "axc",
			t: "ahbgdc",
			expect: false,
		},
		{
			name: "test 3",
			s: "abc",
			t: "ahcgdb",
			expect: false,
		},
		{
			name: "test 4",
			s: "",
			t: "ahbgdc",
			expect: true,
		},
		{
			name: "test 5",
			s: "b",
			t: "abc",
			expect: true,
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, isSubsequence(tt.s, tt.t))
		})
	}
}

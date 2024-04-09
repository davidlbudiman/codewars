package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsIsomorphic(t *testing.T) {
	tests := []struct{
		name string
		s string
		t string
		expect bool
	}{
		{
			name: "test 1",
			s: "paper",
			t: "title",
			expect: true,
		},
		{
			name: "test 2",
			s: "egg",
			t: "add",
			expect: true,
		},
		{
			name: "test 3",
			s: "foo",
			t: "bar",
			expect: false,
		},
		{
			name: "test 4",
			s: "foo",
			t: "bars",
			expect: false,
		},
		{
			name: "test 5",
			s: "badc",
			t: "baba",
			expect: false,
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, isIsomorphic(tt.s, tt.t))
		})
	}
}

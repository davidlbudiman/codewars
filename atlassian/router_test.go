package atlassian

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_AddRoutes(t *testing.T) {
	router := NewRouter()
	tests := []struct{
		name string
		route string
		result string
		expect map[string]string
		expectError error
	}{
		{
			name: "happy path",
			route: "/bar*", // /*bar, b*ar
			result: "result",
			expect: map[string]string{
				"/bar*": "result", 
			},
		},
		{
			name: "sad path - invalid characters in the route",
			route: "/bar%",
			result: "result",
			expect: nil,
			expectError: errors.New("invalid characters in the route"),
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			e := router.addRoute(tt.route, tt.result)
			if tt.expectError != nil {
				assert.EqualError(t, e, tt.expectError.Error())
			} else if tt.expectError == nil && e != nil {
				assert.Fail(t, "Unexpected error")
			} else {
				assert.EqualValues(t, tt.expect, router.tableOfRoutes)
			}
		})
	}
}

// Controller -> Service! -> (Repository)
//               Try Catch -> No DB, DB Not Found, Schema issues
//               RepositoryError

// callRoute -> baroo, barppp, bariii -> result

//      *
// /bar   /foo

// /bar/baz/foo

// *bar/

// *bar/

// /bar/baz/foo
// foobar/
// ^//bar[\w//]*//foo$
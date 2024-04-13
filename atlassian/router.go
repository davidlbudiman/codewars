// interface Router {
//   fun addRoute(path: String, result: String) : Unit;
//   fun callRoute(path: String) : String;

// }

// Usage:

// *()-
// _
// ?=&
// Router.addRoute("/bar", "result")
// Router.callRoute("/bar") -> "result"
// Router.addRoute("/bar/foo?name=Test", "result Test")
// "result %s", name, ....

package atlassian

import "errors"

type Router struct {
	tableOfRoutes map[string]string
	specialRouteKeys []string // ^//bar[\w//]*//foo$

	// O(n) -> time & space 
	// key -> hashed
	// key -> value
}

func (r *Router) addRoute(route, result string) error {
	var runes = []rune(route)
	for _, r := range runes {
		switch r {
		case '(':
		case ')':
		case '-':
		case '?':
		case '=':
		case '&':
			return errors.New("invalid characters in the route")
		}
	}
	r.tableOfRoutes[route] = result
	return nil
}

func (r *Router) callRoute(route string) (result string, e error) {
	return result, nil
}

func NewRouter() *Router {
	return &Router{
		tableOfRoutes: map[string]string{},
	}
}

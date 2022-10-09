// Godis DRY package.
//
// Contains miscellaneous functions that are good enough for re-use but small
// enough to not warrant its own module.
//
// Usage:
//
//	import "github.com/ggustafsson/godis/pkg/godis"
//
//	godis.X()
//
// Author: Göran Gustafsson <gustafsson.g@gmail.com>
//
// License: BSD 3-Clause
package godis

// Source: https://pkg.go.dev/golang.org/x/exp/slices#Contains
func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

// Source: https://pkg.go.dev/golang.org/x/exp/slices#Index
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

package thanos

import (
	"math/rand"
)

// Snap returns a subset of the provided slice containing half of the
// elements, selected at random, rounded down to the nearest integer.
//
// For example, if you provide a slice of `len = 4` or `len = 5`, the
// returned slice will contain 2 elements. If you provide a slice of
// `len = 6` or `len = 7`, the returned slice will contain 3 elements.
func Snap(v []interface{}) []interface{} {
	res := make([]interface{}, len(v)/2)
	for idx, i := range rand.Perm(len(v)) {
		if idx == len(res) {
			break
		}
		res[idx] = v[i]
	}
	return res
}

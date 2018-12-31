package thanos

import "testing"

func TestSnap(t *testing.T) {
	tests := []struct {
		in     []interface{}
		expect int
	}{
		{
			in:     []interface{}{},
			expect: 0,
		},
		{
			in:     []interface{}{1},
			expect: 0,
		},
		{
			in:     []interface{}{1, 2},
			expect: 1,
		},
		{
			in:     []interface{}{1, 2, 3},
			expect: 1,
		},
		{
			in:     []interface{}{1, 2, 3, 4},
			expect: 2,
		},
		{
			in:     []interface{}{1, 2, 3, 4, 5},
			expect: 2,
		},
	}

	for idx, tc := range tests {
		out := Snap(tc.in)

		if len(out) != tc.expect {
			t.Errorf("[TC #%d] Unexpected output length, expected=[%d], got=[%d]", idx, tc.expect, len(out))
		}

		// Ensure all output values are contained within the input slice.
		for _, v := range out {
			if !contains(tc.in, v) {
				t.Errorf("[TC #%d] Unexpected output value, got=[%d]", idx, v)
			}
		}

		// Since the input contains unique values, make sure none of them
		// were selected multiple times.
		if !unique(out) {
			t.Errorf("[TC #%d] Output contains duplicate value, got=[%d]", idx, out)
		}
	}
}

func contains(arr []interface{}, v interface{}) bool {
	for _, val := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func unique(arr []interface{}) bool {
	vals := make(map[interface{}]bool)
	for _, v := range arr {
		if _, ok := vals[v]; ok {
			return false
		}
		vals[v] = true
	}
	return true
}

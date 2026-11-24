package main

import "testing"

func TestCeil(t *testing.T) {
	root := tn(8, tn(4, tn(2, nil, nil), tn(6, nil, nil)), tn(12, tn(10, nil, nil), tn(14, nil, nil)))
	cases := []struct {
		key, want int
		ok        bool
	}{
		{11, 12, true},
		{8, 8, true},
		{15, 0, false}, // above the maximum
		{1, 2, true},
		{5, 6, true},
		{13, 14, true},
	}
	for _, c := range cases {
		got, ok := Ceil(root, c.key)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("Ceil(%d) = (%d,%v), want (%d,%v)", c.key, got, ok, c.want, c.ok)
		}
	}
	if _, ok := Ceil(nil, 5); ok {
		t.Error("empty tree has no ceil")
	}
}

package set

import "testing"

func TestSymmetricDifference(t *testing.T) {
	td := []struct {
		name   string
		s1     Set
		s2     Set
		expSet Set
	}{
		{"symmetric difference of different sets", New(1, 2, 3), New(4, 5, 6), New(1, 2, 3, 4, 5, 6)},
		{"symmetric difference of sets with elements in common", New(1, 2, 3), New(1, 2, 4), New(3, 4)},
		{"symmetric difference of same sets", New(1, 2, 3), New(1, 2, 3), New()},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s1.SymmetricDifference(tc.s2)
			if s.Len() != tc.expSet.Len() {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", tc.expSet.Len(), s.Len(), s.Get())
			}
			for _, n := range tc.expSet.Get() {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.Get())
				}
			}
		})
	}
}

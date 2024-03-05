package set

func New(items ...any) Set {
	s := &set{elements: make(map[any]struct{}, len(items))}

	for _, v := range items {
		s.elements[v] = struct{}{}
	}

	return s
}

type Set interface {
	Add(item any)
	Del(item any)
	Get() []any
	Len() int
	In(item any) bool
	// IsSubsetOf: checks whether set is subset of set2 or not.
	IsSubsetOf(set2 Set) bool
	// IsProperSubsetOf: checks whether set is proper subset of set2 or not.
	// ex: [1,2,3] proper subset of [1,2,3,4] -> true
	IsProperSubsetOf(set2 Set) bool
	// IsSupersetOf: checks whether set is superset of set2 or not.
	IsSupersetOf(set2 Set) bool
	// IsProperSupersetOf: checks whether set is proper superset of set2 or not.
	// ex: [1,2,3,4] proper superset of [1,2,3] -> true
	IsProperSupersetOf(set2 Set) bool
	// Union: gives new union set of both sets.
	// ex: [1,2,3] union [3,4,5] -> [1,2,3,4,5]
	Union(set2 Set) Set
	// Intersection: gives new intersection set of both sets.
	// ex: [1,2,3] Intersection [3,4,5] -> [3]
	Intersection(set2 Set) Set
	// Difference: gives new difference set of both sets.
	// ex: [1,2,3] Difference [3,4,5] -> [1,2]
	Difference(set2 Set) Set
	// SymmetricDifference: gives new symmetric difference set of both sets. 对称查分集，并集减去交集
	// ex: [1,2,3] SymmetricDifference [3,4,5] -> [1,2,4,5]
	SymmetricDifference(set2 Set) Set
}

type set struct {
	elements map[any]struct{}
}

func (s *set) Add(item any) {
	s.elements[item] = struct{}{}
}

func (s *set) Del(item any) {
	delete(s.elements, item)
}

func (s *set) Get() []any {
	items := make([]any, 0, s.Len())
	for k := range s.elements {
		items = append(items, k)
	}

	return items
}

func (s *set) Len() int {
	return len(s.elements)
}

func (s *set) In(item any) bool {
	_, ok := s.elements[item]
	return ok
}

func (s *set) IsSubsetOf(set2 Set) bool {
	if s.Len() > set2.Len() {
		return false
	}

	for _, v := range s.elements {
		if !set2.In(v) {
			return false
		}
	}

	return true
}

func (s *set) IsProperSubsetOf(set2 Set) bool {
	if s.Len() == set2.Len() {
		return false
	}

	return s.IsSubsetOf(set2)
}

func (s *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(s)
}

func (s *set) IsProperSupersetOf(subSet Set) bool {
	if s.Len() == subSet.Len() {
		return false
	}
	return s.IsSupersetOf(subSet)
}

func (s *set) Union(set2 Set) Set {
	unionSet := New()
	for _, item := range s.Get() {
		unionSet.Add(item)
	}
	for _, item := range set2.Get() {
		unionSet.Add(item)
	}
	return unionSet
}

func (s *set) Intersection(set2 Set) Set {
	intersectionSet := New()
	var minSet, maxSet Set = set2, s
	if s.Len() <= set2.Len() {
		minSet = s
		maxSet = set2
	}

	for _, item := range minSet.Get() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (s *set) Difference(set2 Set) Set {
	differenceSet := New()
	for _, item := range s.Get() {
		if !set2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (s *set) SymmetricDifference(set2 Set) Set {
	union := s.Union(set2)
	intersection := s.Intersection(set2)

	return union.Difference(intersection)
}

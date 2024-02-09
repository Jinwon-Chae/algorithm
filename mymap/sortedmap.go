package mymap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Element[Tkey constraints.Ordered, TValue any] struct {
	Key   Tkey
	Value TValue
}

type SortedMap[Tkey constraints.Ordered, TValue any] struct {
	Arr []Element[Tkey, TValue]
}

func (s *SortedMap[Tkey, TValue]) Add(key Tkey, value TValue) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr[idx].Value = value
		return
	}

	s.Arr = append(s.Arr[:idx],
		append([]Element[Tkey, TValue]{
			{Key: key, Value: value},
		}, s.Arr[idx:]...)...)
}

func (s *SortedMap[Tkey, TValue]) Get(key Tkey) (value TValue, ok bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		return s.Arr[idx].Value, true
	}

	var defaultV TValue
	return defaultV, false
}

func (s *SortedMap[Tkey, TValue]) Remove(key Tkey) (removed bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...)

		return true
	}

	return false
}

package set

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Set[T] is a generic Set type that uses a map as the underlying data structure
// The items of the set are stored as map keys. Since map keys have a unique contraint, it guarantees
// uniqueness of items in the set
type empty struct{}

type Set[T comparable] map[T]empty

var emptyval = empty{}

// New[T] creates a new empty set of type T
func New[T comparable]() Set[T] {
	return Set[T]{}
}

// From creates a new set of one or more items of type T
func From[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	for _, i := range items {
		s[i] = emptyval
	}
	return s
}

// Insert adds one or more items of type T to the set.
func (s Set[T]) Insert(items ...T) {
	for _, i := range items {
		s[i] = emptyval
	}
}

// Remove removes an item of type T from the set and returns true if the item existed.
// Remove returns false if the item did not exist in the set.
func (s Set[T]) Remove(item T) bool {
	if s.Contains(item) {
		delete(s, item)
		return true
	}
	return false
}

func (s Set[T]) Clear() {
	for i := range maps.Keys(s) {
		delete(s, i)
	}
}

// ToSlice returns an unordered slice of the sets items of type T
func (s Set[T]) ToSlice() []T {
	items := make([]T, 0, len(s))
	for i := range maps.Keys(s) {
		items = append(items, i)
	}
	return items
}

// Range loops over the set and calls the provided callback.
// If the callback returns true, iteration continues.
// If the callback returns false, iteration stops.
func (s Set[T]) Range(cb func(T) bool) {
	for i := range maps.Keys(s) {
		if !cb(i) {
			return
		}
	}
}

// Iter is an iterator over the set of items of type T
func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range maps.Keys(s) {
			if !yield(i) {
				return
			}
		}
	}
}

// Union combines the items of two different sets of type T into a single new set
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	new := From(s.ToSlice()...)
	new.Insert(s2.ToSlice()...)
	return new
}

// Intersect returns a new set of type T containing only the items common between the two sets.
func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	new := New[T]()

	var smaller, larger Set[T]
	if len(s) > len(s2) {
		smaller, larger = s2, s
	} else {
		smaller, larger = s, s2
	}

	for i := range maps.Keys(smaller) {
		if larger.Contains(i) {
			new.Insert(i)
		}
	}

	return new
}

// Difference returns a new set of type T containing only the items uncommon bewteen the two sets.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	new := New[T]()

	var smaller, larger Set[T]
	if len(s) > len(s2) {
		smaller, larger = s2, s
	} else {
		smaller, larger = s, s2
	}

	for i := range maps.Keys(larger) {
		if !smaller.Contains(i) {
			new.Insert(i)
		}
	}

	return new
}

func (s Set[T]) Size() int {
	return len(s)
}

// Contains returns true if item of type T is present in the set.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func SortedItems[T cmp.Ordered](s Set[T]) []T {
	items := s.ToSlice()
	slices.Sort(items)
	return items
}

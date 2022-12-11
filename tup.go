package ari

import "fmt"

type Tuple[K, V any] struct {
	key   K
	value V
}

// NewTuple(first J, second K) returns a new tuple
func NewTuple[K, V any](key K, value V) *Tuple[K, V] {
	return &Tuple[K, V]{key, value}
}

// ToString() returns a string representation of the tuple
func (t *Tuple[K, V]) ToString() string {
	return fmt.Sprintf("(%v, %v)", t.key, t.value)
}

// Key() returns the key of the tuple
func (t *Tuple[K, V]) Key() K {
	return t.key
}

// Value() returns the value of the tuple
func (t *Tuple[K, V]) Value() V {
	return t.value
}

// SetKey(key K) sets the key of the tuple
func (t *Tuple[K, V]) SetKey(key K) {
	t.key = key
}

// SetValue(value V) sets the value of the tuple
func (t *Tuple[K, V]) SetValue(value V) {
	t.value = value
}

// Equals(other *Tuple[K, V]) returns true if the tuples are equal
// and false otherwise, BEWARE: this method does a shallow comparison via string
func (t *Tuple[K, V]) Equals(other *Tuple[K, V]) bool {
	return t.ToString() == other.ToString()
}

// EqualsByComparator(other *Tuple[K, V], comparator func(K, K) bool) returns true if the tuples are equal
// and false otherwise according to the given comparator function
func (t *Tuple[K, V]) EqualsByComparator(other *Tuple[K, V], comparator func(K, K) bool) bool {
	return comparator(t.key, other.key)
}

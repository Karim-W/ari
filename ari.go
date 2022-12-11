package ari

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type ArrayList[e any] struct {
	elems []e
}

// Add(element any) adds an element to the list
func (l *ArrayList[e]) Add(elem e) {
	l.elems = append(l.elems, elem)
}

// Get(index int) returns the element at the given index
func (l *ArrayList[e]) Get(i int) e {
	return l.elems[i]
}

// Remove(index int) removes the element at the given index
func (l *ArrayList[e]) Remove(i int) e {
	elem := l.elems[i]
	l.elems = append(l.elems[:i], l.elems[i+1:]...)
	return elem
}

// Set(index int, element any) sets the element at the given index
func (l *ArrayList[e]) Set(i int, elem e) {
	l.elems[i] = elem
}

// Size() returns the size of the list
func (l *ArrayList[e]) Size() int {
	return len(l.elems)
}

/*
Sort(comparable func(any,any) bool) sorts the list
using the given comparison function
example usage:
list.Sort(func(a,b int) bool { return a < b })
*/
func (l *ArrayList[e]) Sort(
	comparable func(e, e) bool,
) {
	sort.Slice(l.elems, func(i, j int) bool {
		return comparable(l.elems[i], l.elems[j])
	})
}

// String() returns a string representation of the list
func (l *ArrayList[e]) String() string {
	return fmt.Sprint(l.elems)
}

// ToArray() returns an array representation of the list
func (l *ArrayList[e]) ToArray() []e {
	return l.elems
}

// ToSlice() returns a slice representation of the list
func (l *ArrayList[e]) ToSlice() []e {
	return l.elems
}

// ToSet() returns a set representation of the list
// the set will contain only unique elements
func (l *ArrayList[e]) ToSet() *ArrayList[e] {
	m := make(map[interface{}]bool)
	newList := &ArrayList[e]{}
	for i := range l.elems {
		if _, ok := m[l.elems[i]]; !ok {
			m[l.elems[i]] = true
			newList.Add(l.elems[i])
		}
	}
	return newList
}

// ForEach(function func(any)) iterates over the list
// and applies the given function to each element **SEQUENTIALLY**
func (l *ArrayList[e]) ForEach(callback func(elem e)) {
	for i := range l.elems {
		callback(l.elems[i])
	}
}

// ForEachAsync(function func(any)) iterates over the list
// and applies the given function to each element **CONCURRENTLY**
func (l *ArrayList[e]) ForEachAsync(callback func(element e)) {
	wg := sync.WaitGroup{}
	wg.Add(l.Size())
	for i := range l.elems {
		elem := l.elems[i]
		go func(elem e) {
			callback(elem)
			wg.Done()
		}(elem)
	}
	wg.Wait()
}

// Map(function func(any) any) returns a new list
// contianing the results of applying the given function
func (l *ArrayList[e]) Map(
	callback func(e) e,
) *ArrayList[e] {
	newList := &ArrayList[e]{}
	for i := range l.elems {
		newList.Add(callback(l.elems[i]))
	}
	return newList
}

// MapBy(function func() (string,any)) returns a hash map of entities using given property
func (l *ArrayList[e]) MapBy(
	mapper func(e) (string, e),
) map[string]e {
	m := make(map[string]e)
	for i := range l.elems {
		key, value := mapper(l.elems[i])
		m[key] = value
	}
	return m
}

// Filter(function func(any) bool) returns a new list
// containing only the elements that satisfy the given function
func (l *ArrayList[e]) Filter(
	callback func(element e) bool,
) *ArrayList[e] {
	newList := &ArrayList[e]{}
	for i := range l.elems {
		if callback(l.elems[i]) {
			newList.Add(l.elems[i])
		}
	}
	return newList
}

// ReduceRight(function func(any,any) any, initial any) returns the result
// method applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value
func (l *ArrayList[e]) ReduceRight(
	callback func(acc e, elem e) e,
) e {
	acc := l.elems[l.Size()-1]
	for i := l.Size() - 2; i >= 0; i-- {
		acc = callback(acc, l.elems[i])
	}
	return acc
}

// Find(function func(any) bool) returns the first element that satisfies the given function
func (l *ArrayList[e]) Find(
	callback func(element e) bool,
) *e {
	for i := range l.elems {
		if callback(l.elems[i]) {
			return &l.elems[i]
		}
	}
	return nil
}

// FindIndex(function func(any) bool) returns the index of the first element that satisfies the given function
func (l *ArrayList[e]) FindIndex(
	callback func(element e) bool,
) int {
	for i := range l.elems {
		if callback(l.elems[i]) {
			return i
		}
	}
	return -1
}

// Every(function func(any) bool) returns true if every element satisfies the given function
func (l *ArrayList[e]) Every(
	callback func(element e) bool,
) bool {
	for i := range l.elems {
		if !callback(l.elems[i]) {
			return false
		}
	}
	return true
}

// Some(function func(any) bool) returns true if at least one element satisfies the given function
func (l *ArrayList[e]) Some(
	callback func(element e) bool,
) bool {
	for i := range l.elems {
		if callback(l.elems[i]) {
			return true
		}
	}
	return false
}

// Includes(function func(any) bool) returns true if the list contains the given element
func (l *ArrayList[e]) Includes(
	arrowFunc func(element e) bool,
) bool {
	for i := range l.elems {
		if arrowFunc(l.elems[i]) {
			return true
		}
	}
	return false
}

// Reverse() returns a new list with the elements in reverse order
func (l *ArrayList[e]) Reverse() {
	for i := 0; i < l.Size()/2; i++ {
		j := l.Size() - i - 1
		l.elems[i], l.elems[j] = l.elems[j], l.elems[i]
	}
}

// Slice (start int, end int) returns a new list with the elements in the given range
func (l *ArrayList[e]) Slice(start int, end int) *ArrayList[e] {
	newList := &ArrayList[e]{}
	for i := start; i < end; i++ {
		newList.Add(l.elems[i])
	}
	return newList
}

// Copy() returns a new list with the same elements
func (l *ArrayList[e]) Copy() *ArrayList[e] {
	newList := &ArrayList[e]{}
	for i := range l.elems {
		newList.Add(l.elems[i])
	}
	return newList
}

// Clear() removes all elements from the list
func (l *ArrayList[e]) Clear() {
	l.elems = []e{}
}

// IsEmpty() returns true if the list is empty
func (l *ArrayList[e]) IsEmpty() bool {
	return len(l.elems) == 0
}

// IsNotEmpty() returns true if the list is not empty
func (l *ArrayList[e]) IsNotEmpty() bool {
	return len(l.elems) != 0
}

// Push(element any) adds the given element to the end of the list
func (l *ArrayList[e]) Push(elem e) {
	l.Add(elem)
}

// Pop() removes the last element from the list and returns it
func (l *ArrayList[e]) Pop() e {
	return l.Remove(l.Size() - 1)
}

// Shift() removes the first element from the list and returns it
func (l *ArrayList[e]) Shift() e {
	return l.Remove(0)
}

// Unshift(element any) adds the given element to the beginning of the list
func (l *ArrayList[e]) Unshift(elem e) int {
	l.Add(elem)
	return l.Size()
}

// Peek() returns the last element of the list
func (l *ArrayList[e]) Peek() e {
	return l.elems[l.Size()-1]
}

// PeekFirst() returns the first element of the list
func (l *ArrayList[e]) PeekFirst() e {
	return l.elems[0]
}

// PeekLast() returns the last element of the list
func (l *ArrayList[e]) PeekLast() e {
	return l.elems[l.Size()-1]
}

// PeekNth(n int) returns the nth element of the list
func (l *ArrayList[e]) PeekNth(n int) e {
	return l.elems[n]
}

// PeekNthFromLast(n int) returns the nth element of the list counting from the end
func (l *ArrayList[e]) PeekNthFromLast(n int) e {
	return l.elems[l.Size()-n-1]
}

// PeekRandom() returns a random element of the list
func (l *ArrayList[e]) PeekRandom() e {
	return l.elems[rand.Intn(l.Size())]
}

// Reduce(function func(any,any) any, initial any) returns the result
// the result of applying the given function to each element
func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// ToMapString returns a map[string]e from the list
// where the key is the result of the PrimaryKeyExtractor function given as parameter
func (l *ArrayList[e]) ToMapString(
	PrimaryKeyExtractor func(e) string,
) map[string]e {
	m := make(map[string]e)
	for i := range l.elems {
		m[PrimaryKeyExtractor(l.elems[i])] = l.elems[i]
	}
	return m
}

// ToChainedMapString returns a map[string][]e from the list
// where the key is the result of the PrimaryKeyExtractor function given as parameter
func (l *ArrayList[e]) ToChainedMapString(
	PrimaryKeyExtractor func(e) string,
) map[string][]e {
	m := make(map[string][]e)
	for i := range l.elems {
		key := PrimaryKeyExtractor(l.elems[i])
		if _, ok := m[key]; !ok {
			m[key] = []e{}
		}
		m[key] = append(m[key], l.elems[i])
	}
	return m
}

// Reduce(function func(any,any) any, initial any) returns the result
func (l *ArrayList[e]) Reduce(
	callback func(acc e, elem e) e,
	initial e,
) e {
	acc := initial
	for i := range l.elems {
		acc = callback(acc, l.elems[i])
	}
	return acc
}

// InitFromSlice(s []T) returns a new ArrayList[T] from the given slice
func InitFromSlice[T any](s []T) *ArrayList[T] {
	return &ArrayList[T]{
		elems: s,
	}
}

// InitFromHashMap[K, V any](m map[string]V) returns a new ArrayList[V] from the given map
func InitFromHashMap[K, V any](m map[string]V) *ArrayList[V] {
	list := &ArrayList[V]{}
	for _, v := range m {
		list.Add(v)
	}
	return list
}

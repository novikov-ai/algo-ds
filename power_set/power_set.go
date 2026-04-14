package main

import (
	"constraints"
	"os"
	"strconv"
)

type PowerSet[T constraints.Ordered] struct {
	values map[T]string
}

const size = 20000

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{
		values: make(map[T]string, size),
	}
}

func (ps *PowerSet[T]) Size() int {
	return len(ps.values)
}

func (ps *PowerSet[T]) Put(value T) {
	ps.values[value] = ""
}

func (ps *PowerSet[T]) Get(value T) bool {
	_, ok := ps.values[value]
	return ok
}

func (ps *PowerSet[T]) Remove(value T) bool {
	if !ps.Get(value) {
		return false
	}

	delete(ps.values, value)

	return true
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T] = Init[T]()

	for k := range ps.values {
		if !set2.Get(k) {
			continue
		}

		result.values[k] = ""
	}

	return result
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T] = *ps

	for k := range set2.values {
		result.values[k] = ""
	}

	return result
}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	var result PowerSet[T] = Init[T]()

	for k := range ps.values {
		if set2.Get(k) {
			continue
		}

		result.values[k] = ""
	}

	return result
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	if len(set2.values) > len(ps.values) {
		return false
	}

	for k := range set2.values {
		if !ps.Get(k) {
			return false
		}
	}

	return true
}

func (ps *PowerSet[T]) Equals(set2 PowerSet[T]) bool {
	if len(ps.values) != len(set2.values) {
		return false
	}

	return ps.IsSubset(set2)
}

package main

import "fmt"

type NativeCache[T any] struct {
	size   int
	slots  []string
	values []T
	hits   []int
}

func New[T any](size int) NativeCache[T] {
	return NativeCache[T]{
		size:   size,
		slots:  make([]string, size),
		values: make([]T, size),
		hits:   make([]int, size),
	}
}

func (nc *NativeCache[T]) HashFun(s string) int {
	var sum uint8

	for _, b := range s {
		sum += byte(b)
	}

	return int(sum) % nc.size
}

func (nc *NativeCache[T]) IsKey(key string) bool {
	index := nc.HashFun(key)

	if index >= len(nc.slots) {
		return false
	}

	v := nc.slots[index]

	return v == key
}

func (nc *NativeCache[T]) Get(key string) (T, error) {
	var result T

	ok := nc.IsKey(key)
	if !ok {
		return result, fmt.Errorf("not found")
	}

	index := nc.HashFun(key)

	nc.hits[index]++

	return nc.values[index], nil
}

func (nc *NativeCache[T]) Put(key string, value T) {
	index := nc.HashFun(key)
	if index >= len(nc.slots) {
		return
	}

	if nc.slots[index] == "" {
		nc.slots[index] = key
		nc.values[index] = value
		return
	}

	min := nc.hits[0]
	minHitIndex := 0

	for i, hit := range nc.hits {
		if hit < min {
			min = hit
			minHitIndex = i
			continue
		}
	}

	nc.slots[minHitIndex] = key
	nc.values[minHitIndex] = value
	nc.hits[minHitIndex] = 0
}
package main

import (
	"fmt"
	"os"
	"strconv"
)

type NativeDictionary[T any] struct {
	size   int
	slots  []string
	values []T
}

func Init[T any](sz int) NativeDictionary[T] {
	nd := NativeDictionary[T]{size: sz, slots: nil, values: nil}
	nd.slots = make([]string, sz)
	nd.values = make([]T, sz)
	return nd
}

func (nd *NativeDictionary[T]) HashFun(value string) int {
	var sum uint8

	for _, b := range value {
		sum += byte(b)
	}

	return int(sum) % nd.size
}

func (nd *NativeDictionary[T]) IsKey(key string) bool {
	index := nd.HashFun(key)

	if index >= len(nd.slots) {
		return false
	}

	v := nd.slots[index]

	return v == key
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	var result T

	ok := nd.IsKey(key)
	if !ok {
		return result, fmt.Errorf("not found")
	}

	index := nd.HashFun(key)

	return nd.values[index], nil
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
	index := nd.HashFun(key)
	nd.slots[index] = key
	nd.values[index] = value
}

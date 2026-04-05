package main

import (
	"os"
	"strconv"
)

const notFoundValue = -1

type HashTable struct {
	size  int
	step  int
	slots []string
}

func Init(sz int, stp int) HashTable {
	ht := HashTable{size: sz, step: stp, slots: nil}
	ht.slots = make([]string, sz)
	return ht
}

func (ht *HashTable) HashFun(value string) int {
	return len(value) % len(ht.slots)
}

func (ht *HashTable) SeekSlot(value string) int {
	index := ht.HashFun(value)

	attempts := ht.size // O(n)
	for attempts > 0 {
		if ht.slots[index] == "" {
			return index
		}

		index += ht.step

		if index >= ht.size {
			index -= ht.size
		}

		attempts--
	}

	return notFoundValue
}

func (ht *HashTable) Put(value string) int {
	index := ht.SeekSlot(value)
	if index == notFoundValue {
		return notFoundValue
	}

	ht.slots[index] = value

	return index
}

func (ht *HashTable) Find(value string) int {
	index := ht.HashFun(value)

	attempts := ht.size // O(n)
	for attempts > 0 {
		if ht.slots[index] == value {
			return index
		}

		if ht.slots[index] == "" {
			return notFoundValue
		}

		index += ht.step

		if index >= ht.size {
			index -= ht.size
		}

		attempts--
	}

	return notFoundValue
}

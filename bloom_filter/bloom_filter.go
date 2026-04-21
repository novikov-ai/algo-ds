package main

import (
	"os"
)

type BloomFilter struct {
	filter_len int
	filter     int
}

func New(filterSize int) *BloomFilter {
	bf := BloomFilter{
		filter_len: filterSize,
		filter:     0,
	}

	return &bf
}

func (bf *BloomFilter) Hash1(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum = abs(sum*17+code) % bf.filter_len
	}

	return 1 << sum
}

func (bf *BloomFilter) Hash2(s string) int {
	sum := 0
	for _, char := range s {
		code := int(char)
		sum = abs(sum*223+code) % bf.filter_len
	}
	return 1 << sum
}

// добавляем строку s в фильтр
func (bf *BloomFilter) Add(s string) {
	bf.filter |= bf.Hash1(s)
	bf.filter |= bf.Hash2(s)
}

// проверка, имеется ли строка s в фильтре
func (bf *BloomFilter) IsValue(s string) bool {
	mask := bf.Hash1(s) | bf.Hash2(s)

	return mask == (bf.filter & mask)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

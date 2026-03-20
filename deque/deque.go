package main

import (
	"fmt"
	"os"
)

type Deque[T any] struct {
	values []T
}

func (d *Deque[T]) Size() int {
	return len(d.values)
}

func (d *Deque[T]) AddFront(itm T) {
	d.values = append([]T{itm}, d.values...)
}

func (d *Deque[T]) AddTail(itm T) {
	d.values = append(d.values, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T

	if len(d.values) == 0 {
		return result, fmt.Errorf("not found")
	}

	result = d.values[0]

	dst := make([]T, len(d.values)-1)
	copy(dst, d.values[1:])

	d.values = dst

	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T

	if len(d.values) == 0 {
		return result, fmt.Errorf("not found")
	}

	last := len(d.values)-1

	result = d.values[last]

	var empty T
	d.values[last] = empty

	d.values = d.values[:last]

	return result, nil
}

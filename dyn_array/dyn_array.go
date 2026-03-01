package main

import (
	"os"
	"fmt"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
	arr := make([]T, sz)

	copy(arr, da.array[:da.count])

	da.capacity = sz
	da.array = arr
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	if index < 0 || index > da.count {
		return fmt.Errorf("bad index '%d'", index)
	}

	if da.count == da.capacity {
		da.MakeArray(da.capacity * 2)
	}

	right := da.array[index:da.count]
	shiftedRight := da.array[index+1 : da.count+1]

	copy(shiftedRight, right)

	da.array[index] = itm

	da.count++

	return nil
}

func (da *DynArray[T]) Remove(index int) error {
	if index < 0 || index >= da.count {
		return fmt.Errorf("bad index '%d'", index)
	}

	right := da.array[index+1 : da.count]
	shiftedLeft := da.array[index:da.count]

	var zero T

	da.array[index] = zero

	copy(shiftedLeft, right)

	da.count--

	if needShrink(da.count, da.capacity) {
		da.shrink()
	}

	return nil
}

func needShrink(count, cap int) bool {
	const fillRate = 0.5
	return float64(count/cap) < fillRate
}

func (da *DynArray[T]) shrink() {
	const minCap = 16

	updated := da.capacity * 2 / 3
	if int(updated) < minCap {
		da.capacity = minCap
	} else {
		da.capacity = updated
	}

	da.array = da.array[:da.capacity]
}

func (da *DynArray[T]) Append(itm T) {
	if da.count == da.capacity {
		da.MakeArray(2 * da.capacity)
	}

	da.array = append(da.array, itm)
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T

	if index < 0 || index >= da.count {
		return result, fmt.Errorf("bad index '%d'", index)
	}

	result = da.array[index]

	return result, nil
}
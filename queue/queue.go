package main

import (
	"os"
	"fmt"
)

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T

	if q.Size() == 0 {
		return result, fmt.Errorf("error")
	}

	result = q.elements[0]

	upd := make([]T, len(q.elements)-1)
	copy(upd, q.elements[1:])

	q.elements = upd

	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	q.elements = append(q.elements, itm)
}

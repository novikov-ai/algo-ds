package main

import (
	"fmt"
	"os"
)

type Stack[T any] struct {
	values []T
	count  int
}

func (st *Stack[T]) Size() int {
	return st.count
}

func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[st.count-1]

	return result, nil
}

// Мера сложности: O(1)
func (st *Stack[T]) Pop() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[st.count-1]

	st.values = st.values[:st.count-1]
	st.count -= 1

	return result, nil
}

// Мера сложности: O(1) амортизированно из-за возможных аллокаций на расширение массива
func (st *Stack[T]) Push(itm T) {
	st.values = append(st.values, itm)
	st.count += 1
}
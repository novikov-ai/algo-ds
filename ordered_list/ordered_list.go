package main

import (
	"constraints"
	"os"
)

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool
}

func New[T constraints.Ordered](asc bool) *OrderedList[T] {
	return &OrderedList[T]{
		_ascending: asc,
	}
}

func (l *OrderedList[T]) Count() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (l *OrderedList[T]) Add(item T) {
	node := &Node[T]{value: item}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	current := l.head
	for current != nil {
		compared := l.Compare(current.value, item)

		insertBefore := (l._ascending && compared >= 0) ||
			(!l._ascending && compared <= 0)

		if insertBefore {
			if current.prev == nil {
				node.next = l.head
				l.head.prev = node
				l.head = node
			} else {
				node.prev = current.prev
				node.next = current
				current.prev.next = node
				current.prev = node
			}
			return
		}

		current = current.next
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	if l.head == nil {
		return Node[T]{}, errors.New("not found")
	}

	current := l.head
	for current != nil {
		if current.value == n {
			return *current, nil
		}
		current = current.next
	}

	return Node[T]{}, errors.New("not found")
}

func (l *OrderedList[T]) Delete(n T) {
	current := l.head
	for current != nil {
		if current.value == n {
			if current.prev == nil {
				l.head = current.next
				if l.head != nil {
					l.head.prev = nil
				}
				return
			}
			if current.next == nil {
				l.tail = current.prev
				if l.tail != nil {
					l.tail.next = nil
				}
				return
			}
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
		current = current.next
	}
}

func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l._ascending = asc
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

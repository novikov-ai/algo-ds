package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	current := l.head
	count := 0
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *LinkedList2) Find(n int) (Node, error) {
	current := l.head
	for current != nil {
		if current.value == n {
			return *current, nil
		}
		current = current.next
	}

	return Node{value: -1}, errors.New("not found")
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	current := l.head
	for current != nil {
		if current.value == n {
			nodes = append(nodes, *current)
		}
		current = current.next
	}
	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	current := l.head
	for current != nil {
		if current.value == n {
			if current != l.head && current != l.tail {
				current.prev.next = current.next
				current.next.prev = current.prev
			}

			if current == l.head {
				l.head = current.next
				if current.next == nil {
					l.tail = nil
					return
				}
				current.next.prev = nil
			}

			if current == l.tail {
				l.tail = current.prev
				current.prev.next = nil
			}

			if !all {
				return
			}
		}

		current = current.next
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	if l.head == nil {
		l.head = &add
		l.tail = &add
		return
	}

	current := l.head
	for current != nil {
		if current != after {
			current = current.next
			continue
		}

		if current == l.tail {
			l.tail.next = &add
			add.prev = l.tail
			l.tail = &add
		} else {
			current.next.prev = &add
			add.next = current.next

			current.next = &add
			add.prev = current
		}
		return
	}

	if l.head == nil {
		l.head = &add
		l.tail = &add
		return
	}

	if l.tail == nil {
		l.tail = l.head
	}

	add.next = l.head
	l.head.prev = &add

	l.head = &add
}

func (l *LinkedList2) InsertFirst(first Node) {
	// clean node before inserting
	first.prev = nil
	first.next = nil

	if l.head == nil {
		l.head = &first
		l.tail = &first

		return
	}

	if l.tail == nil {
		l.tail = l.head
	}

	first.next = l.head
	l.head.prev = &first

	l.head = &first
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}

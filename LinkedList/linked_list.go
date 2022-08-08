package main

import (
	"os"
	"reflect"
	"errors"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	current := l.head
	count := 0

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (l *LinkedList) Find(n int) (Node, error) {
	current := l.head
	for current != nil {
		if current.value == n {
			return *current, nil
		}
		current = current.next
	}

	return Node{value: -1, next: nil}, errors.New("node was not found")
}

func (l *LinkedList) FindAll(n int) []Node {
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

func (l *LinkedList) Delete(n int, all bool) {
	var current, prev *Node
	current = l.head

	for current != nil {
		if current.value == n {
			if l.Count() == 1 {
				l.Clean()
				return
			}

			if current == l.head {
				l.head = current.next
			} else if current == l.tail {
				prev.next = nil
				l.tail = prev
			} else {
				prev.next = current.next
			}

			if !all {
				return
			}
		}

		prev = current
		current = current.next
	}
}

func (l *LinkedList) Insert(after *Node, add Node) {
	_, err := l.Find(after.value)
	if err != nil {
		l.InsertFirst(add)
		return
	}

	current := l.head
	for current != nil {
		if current == after {
			cache := current.next

			current.next = &add
			add.next = cache

			if l.tail == current {
				l.tail = &add
				return
			}
		}

		current = current.next
	}
}

func (l *LinkedList) InsertFirst(first Node) {
	if l.Count() == 0 {
		l.head = &first
		l.tail = &first
		return
	}

	first.next = l.head
	l.head = &first
}

func (l *LinkedList) Clean() {
	if l.Count() > 0 {
		l.head.next = nil
	}

	l.head = nil
	l.tail = nil
}

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
	count := 0
	current := l.head

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

	return Node{value: -1, next: nil}, errors.New("node was not found")
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
	_, err := l.Find(n)
	if err != nil {
		return
	}

	current := l.head
	for current != nil {
		if current.value != n {
			current = current.next
			continue
		}

		if current == l.head && current == l.tail {
			l.Clean()
			return
		}

		if current == l.head {
			next := current.next
			next.prev = nil
			l.head = next
		} else if current == l.tail {
			prev := current.prev
			prev.next = nil
			l.tail = prev
		} else {
			prev := current.prev
			next := current.next

			prev.next = next
			next.prev = prev
		}

		current = current.next

		if !all {
			return
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	current := l.head
	for current != nil {
		if current != after {
			current = current.next
			continue
		}

		if current == l.tail {
			current.next = &add
			add.prev = current

			l.tail = &add

		} else {
			current.next.prev = &add
			add.next = current.next

			current.next = &add
			add.prev = current
		}

		return
	}

	l.InsertFirst(add)
}

func (l *LinkedList2) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
		return
	}

	currentHead := l.head
	currentHead.prev = &first

	first.next = currentHead
	l.head = &first
}

func (l *LinkedList2) Clean() {
	if l.head != nil {
		l.head.next = nil
	}
	l.head = nil

	if l.tail != nil {
		l.tail.prev = nil
	}
	l.tail = nil
}

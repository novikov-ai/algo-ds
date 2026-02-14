package main

import (
	"errors"
	"os"
	"reflect"
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
	sum := 0
	node := l.head
	for node != nil {
		sum++
		node = node.next
	}

	return sum
}

func (l *LinkedList) Find(n int) (Node, error) {
	current := l.head
	for current != nil {
		if current.value == n {
			return *current, nil
		}
		current = current.next
	}

	return Node{value: -1, next: nil}, errors.New("not found")
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
	current := l.head
	var prev *Node = nil

	for current != nil {
		if current.value != n {
			prev = current
			current = current.next
			continue
		}

		switch current {
		case l.head:
			l.head = current.next
		case l.tail:
			prev.next = nil
			l.tail = prev
		default:
			prev.next = current.next
		}

		if !all {
			return
		}

		current = current.next
	}
}

// pre-conditions: after Node exists
func (l *LinkedList) Insert(after *Node, add Node) {
	if l.head == nil {
		l.head = &add
		l.tail = &add
		return
	}

	current := l.head
	for current != nil {
		if current == after {
			add.next = current.next
			after.next = &add

			if current == l.tail {
				l.tail = &add
			}
			return
		}

		current = current.next
	}
}

func (l *LinkedList) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
		return
	}

	first.next = l.head
	l.head = &first
}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
}

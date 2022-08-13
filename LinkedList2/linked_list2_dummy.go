package main

import "errors"

// [Dummy-HEAD] <---> [item] <---> ... <---> [item] <---> [Dummy-TAIL]

type NodeDummy struct {
	prev  *NodeDummy
	next  *NodeDummy
	value int
	dummy bool
}

type LinkedList2Dummy struct {
	head *NodeDummy
	tail *NodeDummy
}

func NewLinkedList2Dummy() LinkedList2Dummy {
	newList := LinkedList2Dummy{}
	newList.Clean()
	return newList
}

func (l *LinkedList2Dummy) AddInTailDummy(item NodeDummy) {
	tailPrev := l.tail.prev

	l.tail.prev = &item
	item.next = l.tail

	item.prev = tailPrev
	tailPrev.next = &item
}

func (l *LinkedList2Dummy) CountDummy() int {
	count := 0
	current := l.head

	for current != nil {
		if !current.dummy {
			count++
		}
		current = current.next
	}

	return count
}

func (l *LinkedList2Dummy) FindDummy(n int) (NodeDummy, error) {
	current := l.head
	for current != nil {
		if current.value == n && !current.dummy {
			return *current, nil
		}
		current = current.next
	}

	return NodeDummy{value: -1, next: nil}, errors.New("node was not found")
}

func (l *LinkedList2Dummy) FindAllDummy(n int) []NodeDummy {
	var nodes []NodeDummy

	current := l.head
	for current != nil {
		if current.value == n && !current.dummy {
			nodes = append(nodes, *current)
		}
		current = current.next
	}

	return nodes
}

func (l *LinkedList2Dummy) DeleteDummy(n int, all bool) {
	_, err := l.FindDummy(n)
	if err != nil {
		return
	}

	current := l.head
	for current != nil {
		if current.value != n || current.dummy {
			current = current.next
			continue
		}

		prev := current.prev
		next := current.next

		prev.next = next
		next.prev = prev

		current = current.next

		if !all {
			return
		}
	}
}

func (l *LinkedList2Dummy) InsertDummy(after *NodeDummy, add NodeDummy) {
	current := l.head
	for current != nil {
		if current != after || current.dummy {
			current = current.next
			continue
		}

		currentNext := current.next

		current.next = &add
		add.prev = current

		add.next = currentNext
		currentNext.prev = &add

		return
	}

	l.InsertFirstDummy(add)
}

func (l *LinkedList2Dummy) InsertFirstDummy(first NodeDummy) {
	headNext := l.head.next

	l.head.next = &first
	first.prev = l.head

	first.next = headNext
	headNext.prev = &first
}

func (l *LinkedList2Dummy) Clean() {
	l.head = &NodeDummy{dummy: true}
	l.tail = &NodeDummy{dummy: true}

	l.head.next = l.tail
	l.tail.prev = l.head
}

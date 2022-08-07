package main

import "errors"

func CombineTwoLists(l1, l2 LinkedList) (LinkedList, error) {
	combinedList := LinkedList{}

	if l1.Count() != l2.Count() {
		return combinedList, errors.New("len of input lists not equal")
	}

	currentL1 := l1.head
	currentL2 := l2.head

	for currentL1 != nil {
		combinedNode := Node{value: currentL1.value + currentL2.value}
		combinedList.AddInTail(combinedNode)

		currentL1 = currentL1.next
		currentL2 = currentL2.next
	}
	return combinedList, nil
}

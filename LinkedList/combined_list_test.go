package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombineTwoLists(t *testing.T) {
	testCases := []struct {
		title        string
		l1, l2       LinkedList
		combinedList LinkedList
	}{
		{title: "empty l1 and l2", l1: LinkedList{}, l2: LinkedList{}, combinedList: LinkedList{}},
		{title: "l1 and l2 has different len", l1: getLinkedList([]int{4, 1, 3}), l2: getLinkedList([]int{3, 1}), combinedList: LinkedList{}},
		{title: "l1 and l2 has only 1 element", l1: getLinkedList([]int{4}), l2: getLinkedList([]int{1}), combinedList: getLinkedList([]int{5})},
		{title: "l1 and l2 has many elements", l1: getLinkedList([]int{4, 1, 3, 5}), l2: getLinkedList([]int{6, -1, 7, 15}), combinedList: getLinkedList([]int{10, 0, 10, 20})},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.title, func(t *testing.T) {
			combined, err := CombineTwoLists(test.l1, test.l2)
			if test.l1.Count() != test.l2.Count() {
				require.Error(t, err)
			}

			require.Equal(t, test.combinedList, combined)
		})
	}
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestAddInTailDummy(t *testing.T) {
	t.Run("add in tail 3 times", func(t *testing.T) {
		list := NewLinkedList2Dummy()

		list.AddInTailDummy(NodeDummy{value: 1})
		require.Equal(t, 1, list.tail.prev.value)

		list.AddInTailDummy(NodeDummy{value: 3})
		require.Equal(t, 3, list.tail.prev.value)

		list.AddInTailDummy(NodeDummy{value: 5})
		require.Equal(t, 5, list.tail.prev.value)
	})
}

func TestCountDummy(t *testing.T) {
	t.Run("count len after insertion", func(t *testing.T) {
		list := NewLinkedList2Dummy()

		require.Equal(t, 0, list.CountDummy())

		list.AddInTailDummy(NodeDummy{value: 1})
		require.Equal(t, 1, list.CountDummy())

		list.AddInTailDummy(NodeDummy{value: 3})
		require.Equal(t, 2, list.CountDummy())

		list.AddInTailDummy(NodeDummy{value: 5})
		require.Equal(t, 3, list.CountDummy())
	})
}

func TestFindDummy(t *testing.T) {
	list := NewLinkedList2Dummy()
	nodeValues := []int{1, 3, 5, 7, 9, 11}
	for _, v := range nodeValues {
		list.AddInTailDummy(NodeDummy{value: v})
	}

	t.Run(fmt.Sprintf("finding 0 at %v", nodeValues), func(t *testing.T) {
		node, err := list.FindDummy(0)
		require.Error(t, err)
		require.Equal(t, NodeDummy{value: -1, next: nil}, node)
	})

	t.Run(fmt.Sprintf("finding 4 at %v", nodeValues), func(t *testing.T) {
		node, err := list.FindDummy(4)
		require.Error(t, err)
		require.Equal(t, NodeDummy{value: -1, next: nil}, node)
	})

	for _, v := range nodeValues {
		v := v

		t.Run(fmt.Sprintf("finding %v at %v", v, nodeValues), func(t *testing.T) {
			node, err := list.FindDummy(v)
			require.NoError(t, err)
			require.Equal(t, v, node.value)
		})
	}
}

func TestFindAllDummy(t *testing.T) {
	inputValues := []int{0, 1, 1, 3, 1, 5, 11}
	list := getLinkedListDummy(inputValues)

	testCases := []struct {
		value      int
		foundCount int
	}{
		{value: 1, foundCount: 3},
		{value: 5, foundCount: 1},
		{value: 19, foundCount: 0},
	}

	t.Run(fmt.Sprintf("finding 0 at %v", inputValues), func(t *testing.T) {
		nodes := list.FindAllDummy(0)
		for _, node := range nodes {
			require.Equal(t, 0, node.value)
		}
		require.Equal(t, 1, len(nodes))
	})

	for _, test := range testCases {
		test := test
		t.Run(fmt.Sprintf("finding %v at %v", test.value, inputValues), func(t *testing.T) {
			nodes := list.FindAllDummy(test.value)
			require.Equal(t, test.foundCount, len(nodes))
		})
	}
}

func TestDeleteDummy(t *testing.T) {
	testCases := []struct {
		title       string
		list        LinkedList2Dummy
		removeValue int
		removeAll   bool
		nodesLeft   int
		headValue   int
		tailValue   int
		leftValues  []int
	}{
		{
			title: "empty list", list: NewLinkedList2Dummy(),
			removeValue: 3, removeAll: true, nodesLeft: 0, leftValues: []int{},
		},
		{
			title: "delete the only node", list: getLinkedListDummy([]int{3}),
			removeValue: 3, removeAll: true, nodesLeft: 0, leftValues: []int{},
		},
		{
			title: "delete not existing node", list: getLinkedListDummy([]int{3}),
			removeValue: 1, removeAll: true, nodesLeft: 1, headValue: 3, tailValue: 3, leftValues: []int{3},
		},
		{
			title: "delete head", list: getLinkedListDummy([]int{3, 4}),
			removeValue: 3, removeAll: true, nodesLeft: 1, headValue: 4, tailValue: 4, leftValues: []int{4},
		},
		{
			title: "delete tail", list: getLinkedListDummy([]int{3, 4}),
			removeValue: 4, removeAll: true, nodesLeft: 1, headValue: 3, tailValue: 3, leftValues: []int{3},
		},
		{
			title: "delete middle", list: getLinkedListDummy([]int{3, 4, 5}),
			removeValue: 4, removeAll: true, nodesLeft: 2, headValue: 3, tailValue: 5, leftValues: []int{3, 5},
		},
		{
			title: "delete only 1 matching value", list: getLinkedListDummy([]int{3, 1, 2, 3, 4, 5, 3}),
			removeValue: 3, removeAll: false, nodesLeft: 6, headValue: 1, tailValue: 3, leftValues: []int{1, 2, 3, 4, 5, 3},
		},
		{
			title: "delete all matching values", list: getLinkedListDummy([]int{3, 1, 2, 3, 4, 5, 3}),
			removeValue: 3, removeAll: true, nodesLeft: 4, headValue: 1, tailValue: 5, leftValues: []int{1, 2, 4, 5},
		},
		{
			title: "delete all matching values (left one)", list: getLinkedListDummy([]int{3, 3, 3, 3, 4, 3, 3}),
			removeValue: 3, removeAll: true, nodesLeft: 1, headValue: 4, tailValue: 4, leftValues: []int{4},
		},
		{
			title: "list with all matching values", list: getLinkedListDummy([]int{3, 3, 3}),
			removeValue: 3, removeAll: true, nodesLeft: 0, leftValues: []int{},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.DeleteDummy(test.removeValue, test.removeAll)
			require.Equal(t, test.nodesLeft, test.list.CountDummy())
			if test.list.CountDummy() != 0 {
				require.Equal(t, test.headValue, test.list.head.next.value)
				require.Equal(t, test.tailValue, test.list.tail.prev.value)
			}

			checkForwardRelationsDummy(t, test.list, test.leftValues)
			checkBackwardRelationsDummy(t, test.list, test.leftValues)
		})
	}
}

func TestInsertDummy(t *testing.T) {
	linkedLists := []LinkedList2Dummy{
		getLinkedListDummy([]int{1, 2}),
		getLinkedListDummy([]int{1, 2}),
		getLinkedListDummy([]int{1, 2}),
		getLinkedListDummy([]int{1, 2, 4, 5}),
	}

	testCases := []struct {
		title       string
		list        LinkedList2Dummy
		nodeAfter   *NodeDummy
		nodeInsert  NodeDummy
		headValue   int
		tailValue   int
		countAfter  int
		valuesOrder []int
	}{
		{
			title: "empty list", list: NewLinkedList2Dummy(),
			nodeAfter: &NodeDummy{value: 15}, nodeInsert: NodeDummy{value: 1}, headValue: 1, tailValue: 1, countAfter: 1, valuesOrder: []int{1},
		},
		{
			title: "insert first", list: linkedLists[0],
			nodeAfter: &NodeDummy{value: 15}, nodeInsert: NodeDummy{value: 0}, headValue: 0, tailValue: 2, countAfter: 3, valuesOrder: []int{0, 1, 2},
		},
		{
			title: "insert after head", list: linkedLists[1],
			nodeAfter: linkedLists[1].head.next, nodeInsert: NodeDummy{value: 0}, headValue: 1, tailValue: 2, countAfter: 3, valuesOrder: []int{1, 0, 2},
		},
		{
			title: "insert after tail", list: linkedLists[2],
			nodeAfter: linkedLists[2].tail.prev, nodeInsert: NodeDummy{value: 0}, headValue: 1, tailValue: 0, countAfter: 3, valuesOrder: []int{1, 2, 0},
		},
		{
			title: "insert into middle", list: linkedLists[3],
			nodeAfter: linkedLists[3].head.next.next, nodeInsert: NodeDummy{value: 3}, headValue: 1, tailValue: 5, countAfter: 5, valuesOrder: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.InsertDummy(test.nodeAfter, test.nodeInsert)
			require.Equal(t, test.countAfter, test.list.CountDummy())
			if test.list.CountDummy() != 0 {
				require.Equal(t, test.headValue, test.list.head.next.value)
				require.Equal(t, test.tailValue, test.list.tail.prev.value)
			}

			checkForwardRelationsDummy(t, test.list, test.valuesOrder)
			checkBackwardRelationsDummy(t, test.list, test.valuesOrder)
		})
	}
}

func TestInsertFirstDummy(t *testing.T) {
	testCases := []struct {
		title      string
		list       LinkedList2Dummy
		nodeInsert NodeDummy
		headValue  int
		tailValue  int
		countAfter int
	}{
		{title: "empty list", list: NewLinkedList2Dummy(), nodeInsert: NodeDummy{value: 1}, headValue: 1, tailValue: 1, countAfter: 1},
		{title: "insert at list with 1 value", list: getLinkedListDummy([]int{2}), nodeInsert: NodeDummy{value: 1}, headValue: 1, tailValue: 2, countAfter: 2},
		{title: "insert at list with 2 values", list: getLinkedListDummy([]int{2, 3, 4}), nodeInsert: NodeDummy{value: 1}, headValue: 1, tailValue: 4, countAfter: 4},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.InsertFirstDummy(test.nodeInsert)
			require.Equal(t, test.countAfter, test.list.CountDummy())
			if test.list.CountDummy() != 0 {
				require.Equal(t, test.headValue, test.list.head.next.value)
				require.Equal(t, test.tailValue, test.list.tail.prev.value)
			}
		})
	}
}

func TestCleanDummy(t *testing.T) {
	testCases := []struct {
		title string
		list  LinkedList2Dummy
	}{
		{title: "empty list", list: NewLinkedList2Dummy()},
		{title: "list with values", list: getLinkedListDummy([]int{0, 1, 1, 3, 1, 5, 11})},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.Clean()

			require.True(t, test.list.head.dummy)
			require.True(t, test.list.tail.dummy)
			require.Equal(t, 0, test.list.CountDummy())
		})
	}
}

func getLinkedListDummy(values []int) LinkedList2Dummy {
	list := NewLinkedList2Dummy()
	for _, v := range values {
		list.AddInTailDummy(NodeDummy{value: v})
	}
	return list
}

func checkForwardRelationsDummy(t *testing.T, list LinkedList2Dummy, correctOrder []int) {
	t.Helper()

	if list.CountDummy() == 0 {
		require.Equal(t, len(correctOrder), 0)
		return
	}

	currentNode := list.head
	index := 0
	for currentNode != nil {
		if !currentNode.dummy {
			assert.Equal(t, correctOrder[index], currentNode.value)
			index++
		}
		currentNode = currentNode.next
	}
}

func checkBackwardRelationsDummy(t *testing.T, list LinkedList2Dummy, correctOrder []int) {
	t.Helper()

	if list.CountDummy() == 0 {
		require.Equal(t, len(correctOrder), 0)
		return
	}

	currentNode := list.tail
	index := list.CountDummy() - 1
	for currentNode != nil {
		if !currentNode.dummy {
			assert.Equal(t, correctOrder[index], currentNode.value)
			index--
		}
		currentNode = currentNode.prev
	}
}

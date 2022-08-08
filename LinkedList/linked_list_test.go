package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddInTail(t *testing.T) {
	t.Run("add in tail 3 times", func(t *testing.T) {
		list := LinkedList{}

		list.AddInTail(Node{value: 1})
		require.Equal(t, 1, list.head.value)
		require.Equal(t, 1, list.tail.value)

		list.AddInTail(Node{value: 3})
		require.Equal(t, 1, list.head.value)
		require.Equal(t, 3, list.head.next.value)
		require.Equal(t, 3, list.tail.value)

		list.AddInTail(Node{value: 5})
		require.Equal(t, 1, list.head.value)
		require.Equal(t, 5, list.head.next.next.value)
		require.Equal(t, 5, list.tail.value)
	})
}

func TestCount(t *testing.T) {
	t.Run("count len after insertion", func(t *testing.T) {
		list := LinkedList{}

		require.Equal(t, 0, list.Count())

		list.AddInTail(Node{value: 1})
		require.Equal(t, 1, list.Count())

		list.AddInTail(Node{value: 3})
		require.Equal(t, 2, list.Count())

		list.AddInTail(Node{value: 5})
		require.Equal(t, 3, list.Count())
	})
}

func TestFind(t *testing.T) {
	list := LinkedList{}
	nodeValues := []int{1, 3, 5, 7, 9, 11}
	for _, v := range nodeValues {
		list.AddInTail(Node{value: v})
	}

	t.Run(fmt.Sprintf("finding 4 at %v", nodeValues), func(t *testing.T) {
		node, err := list.Find(4)
		require.Error(t, err)
		require.Equal(t, Node{value: -1, next: nil}, node)
	})

	for _, v := range nodeValues {
		v := v

		t.Run(fmt.Sprintf("finding %v at %v", v, nodeValues), func(t *testing.T) {
			node, err := list.Find(v)
			require.NoError(t, err)
			require.Equal(t, v, node.value)
		})
	}
}

func TestFindAll(t *testing.T) {
	inputValues := []int{0, 1, 1, 3, 1, 5, 11}
	list := getLinkedList(inputValues)

	testCases := []struct {
		value      int
		foundCount int
	}{
		{value: 1, foundCount: 3},
		{value: 5, foundCount: 1},
		{value: 19, foundCount: 0},
	}

	for _, test := range testCases {
		test := test
		t.Run(fmt.Sprintf("finding %v at %v", test.value, inputValues), func(t *testing.T) {
			nodes := list.FindAll(test.value)
			require.Equal(t, test.foundCount, len(nodes))
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		title       string
		list        LinkedList
		removeValue int
		removeAll   bool
		nodesLeft   int
		headValue   int
		tailValue   int
	}{
		{title: "empty list", list: LinkedList{}, removeValue: 3, removeAll: true, nodesLeft: 0},
		{title: "delete the only node", list: getLinkedList([]int{3}), removeValue: 3, removeAll: true, nodesLeft: 0},
		{title: "delete not existing node", list: getLinkedList([]int{3}), removeValue: 1, removeAll: true, nodesLeft: 1, headValue: 3, tailValue: 3},
		{title: "delete head", list: getLinkedList([]int{3, 4}), removeValue: 3, removeAll: true, nodesLeft: 1, headValue: 4, tailValue: 4},
		{title: "delete tail", list: getLinkedList([]int{3, 4}), removeValue: 4, removeAll: true, nodesLeft: 1, headValue: 3, tailValue: 3},
		{title: "delete middle", list: getLinkedList([]int{3, 4, 5}), removeValue: 4, removeAll: true, nodesLeft: 2, headValue: 3, tailValue: 5},
		{title: "delete only 1 matching value", list: getLinkedList([]int{3, 1, 2, 3, 4, 5, 3}), removeValue: 3, removeAll: false, nodesLeft: 6, headValue: 1, tailValue: 3},
		{title: "delete all matching values", list: getLinkedList([]int{3, 1, 2, 3, 4, 5, 3}), removeValue: 3, removeAll: true, nodesLeft: 4, headValue: 1, tailValue: 5},
		{title: "list with all matching values", list: getLinkedList([]int{3, 3, 3}), removeValue: 3, removeAll: true, nodesLeft: 0},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.Delete(test.removeValue, test.removeAll)
			require.Equal(t, test.nodesLeft, test.list.Count())
			if test.list.Count() != 0 {
				require.Equal(t, test.headValue, test.list.head.value)
				require.Equal(t, test.tailValue, test.list.tail.value)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	linkedLists := []LinkedList{
		getLinkedList([]int{1, 2}),
		getLinkedList([]int{1, 2}),
		getLinkedList([]int{1, 2}),
		getLinkedList([]int{1, 2, 4, 5}),
	}

	testCases := []struct {
		title      string
		list       LinkedList
		nodeAfter  *Node
		nodeInsert Node
		headValue  int
		tailValue  int
		countAfter int
	}{
		{title: "empty list", list: LinkedList{}, nodeAfter: &Node{value: 15}, nodeInsert: Node{value: 1}, headValue: 1, tailValue: 1, countAfter: 1},
		{title: "insert first", list: linkedLists[0], nodeAfter: &Node{value: 15}, nodeInsert: Node{value: 0}, headValue: 0, tailValue: 2, countAfter: 3},
		{title: "insert after head", list: linkedLists[1], nodeAfter: linkedLists[1].head, nodeInsert: Node{value: 0}, headValue: 1, tailValue: 2, countAfter: 3},
		{title: "insert after tail", list: linkedLists[2], nodeAfter: linkedLists[2].tail, nodeInsert: Node{value: 0}, headValue: 1, tailValue: 0, countAfter: 3},
		{title: "insert into middle", list: linkedLists[3], nodeAfter: linkedLists[3].head.next, nodeInsert: Node{value: 3}, headValue: 1, tailValue: 5, countAfter: 5},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.Insert(test.nodeAfter, test.nodeInsert)
			require.Equal(t, test.countAfter, test.list.Count())
			if test.list.Count() != 0 {
				require.Equal(t, test.headValue, test.list.head.value)
				require.Equal(t, test.tailValue, test.list.tail.value)
			}
		})
	}
}

func TestInsertFirst(t *testing.T) {
	testCases := []struct {
		title      string
		list       LinkedList
		nodeInsert Node
		headValue  int
		tailValue  int
		countAfter int
	}{
		{title: "empty list", list: LinkedList{}, nodeInsert: Node{value: 1}, headValue: 1, tailValue: 1, countAfter: 1},
		{title: "insert at list with 1 value", list: getLinkedList([]int{2}), nodeInsert: Node{value: 1}, headValue: 1, tailValue: 2, countAfter: 2},
		{title: "insert at list with 2 values", list: getLinkedList([]int{2, 3, 4}), nodeInsert: Node{value: 1}, headValue: 1, tailValue: 4, countAfter: 4},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.InsertFirst(test.nodeInsert)
			require.Equal(t, test.countAfter, test.list.Count())
			if test.list.Count() != 0 {
				require.Equal(t, test.headValue, test.list.head.value)
				require.Equal(t, test.tailValue, test.list.tail.value)
			}
		})
	}
}

func TestClean(t *testing.T) {
	testCases := []struct {
		title string
		list  LinkedList
	}{
		{title: "empty list", list: LinkedList{}},
		{title: "list with values", list: getLinkedList([]int{0, 1, 1, 3, 1, 5, 11})},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.title, func(t *testing.T) {
			test.list.Clean()

			require.Nil(t, test.list.head)
			require.Nil(t, test.list.tail)
			require.Equal(t, 0, test.list.Count())
		})
	}
}

func getLinkedList(values []int) LinkedList {
	list := LinkedList{}
	for _, v := range values {
		list.AddInTail(Node{value: v})
	}
	return list
}

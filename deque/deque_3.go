package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Size(t *testing.T) {
	tests := []struct {
		name     string
		deque    Deque[int]
		expected int
	}{
		{
			name: "3",
			deque: Deque[int]{
				values: []int{1, 2, 3},
			},
			expected: 3,
		},
		{
			name: "0",
			deque: Deque[int]{
				values: []int{},
			},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.deque.Size())
		})
	}
}

func Test_AddFront(t *testing.T) {
	tests := []struct {
		name           string
		deque          Deque[int]
		itm            int
		expectedSize   int
		expectedValues []int
	}{
		{
			name: "+1",
			deque: Deque[int]{
				values: []int{1, 2, 3},
			},
			itm:            4,
			expectedSize:   4,
			expectedValues: []int{4, 1, 2, 3},
		},
		{
			name: "empty",
			deque: Deque[int]{
				values: []int{},
			},
			itm:            4,
			expectedSize:   1,
			expectedValues: []int{4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.deque.AddFront(tc.itm)
			assert.Equal(t, tc.expectedSize, tc.deque.Size())
			assert.Equal(t, tc.expectedValues, tc.deque.values)
		})
	}
}

func Test_AddTail(t *testing.T) {
	tests := []struct {
		name           string
		deque          Deque[int]
		itm            int
		expectedSize   int
		expectedValues []int
	}{
		{
			name: "+1",
			deque: Deque[int]{
				values: []int{1, 2, 3},
			},
			itm:            4,
			expectedSize:   4,
			expectedValues: []int{1, 2, 3, 4},
		},
		{
			name: "empty",
			deque: Deque[int]{
				values: []int{},
			},
			itm:            4,
			expectedSize:   1,
			expectedValues: []int{4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.deque.AddTail(tc.itm)
			assert.Equal(t, tc.expectedSize, tc.deque.Size())
			assert.Equal(t, tc.expectedValues, tc.deque.values)
		})
	}
}

func Test_RemoveFront(t *testing.T) {
	tests := []struct {
		name           string
		deque          Deque[int]
		expectedItm    int
		expectedSize   int
		expectedValues []int
		expectedErr    bool
	}{
		{
			name: "-1",
			deque: Deque[int]{
				values: []int{1, 2, 3},
			},
			expectedItm:    1,
			expectedSize:   2,
			expectedValues: []int{2, 3},
			expectedErr:    false,
		},
		{
			name: "empty",
			deque: Deque[int]{
				values: []int{4},
			},
			expectedItm:    4,
			expectedSize:   0,
			expectedValues: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, err := tc.deque.RemoveFront()

			assert.Equal(t, tc.expectedItm, value)
			assert.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expectedSize, tc.deque.Size())
			assert.Equal(t, tc.expectedValues, tc.deque.values)
		})
	}
}

func Test_RemoveTail(t *testing.T) {
	tests := []struct {
		name           string
		deque          Deque[int]
		expectedItm    int
		expectedSize   int
		expectedValues []int
		expectedErr    bool
	}{
		{
			name: "-1",
			deque: Deque[int]{
				values: []int{1, 2, 3},
			},
			expectedItm:    3,
			expectedSize:   2,
			expectedValues: []int{1, 2},
			expectedErr:    false,
		},
		{
			name: "empty",
			deque: Deque[int]{
				values: []int{4},
			},
			expectedItm:    4,
			expectedSize:   0,
			expectedValues: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value, err := tc.deque.RemoveTail()

			assert.Equal(t, tc.expectedItm, value)
			assert.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expectedSize, tc.deque.Size())
			assert.Equal(t, tc.expectedValues, tc.deque.values)
		})
	}
}

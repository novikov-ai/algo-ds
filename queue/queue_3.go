package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Size(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "ok",
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q := Queue[int]{
				elements: []int{1, 2, 3},
			}

			got := q.Size()

			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Dequeue(t *testing.T) {
	tests := []struct {
		name         string
		queue        Queue[int]
		expectedItem int
		expected     int
		expectedErr  bool
	}{
		{
			name: "ok",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			expectedItem: 1,
			expected:     2,
			expectedErr:  false,
		},
		{
			name: "last",
			queue: Queue[int]{
				elements: []int{1},
			},
			expectedItem: 1,
			expected:     0,
			expectedErr:  false,
		},
		{
			name: "err",
			queue: Queue[int]{
				elements: []int{},
			},
			expectedItem: 0,
			expected:     0,
			expectedErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.queue.Dequeue()
			if tc.expectedErr {
				assert.NotNil(t, err)
			}

			gotSize := tc.queue.Size()

			assert.Equal(t, tc.expected, gotSize)
			assert.Equal(t, tc.expectedItem, got)
		})
	}
}

func Test_Enqueue(t *testing.T) {
	tests := []struct {
		name     string
		queue    Queue[int]
		item     int
		expected int
	}{
		{
			name: "ok",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			item:     4,
			expected: 4,
		},
		{
			name: "one",
			queue: Queue[int]{
				elements: []int{},
			},
			item:     1,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.Enqueue(tc.item)

			gotSize := tc.queue.Size()

			assert.Equal(t, tc.expected, gotSize)
		})
	}
}

func Test_Rotate(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		queue    Queue[int]
		expected Queue[int]
	}{
		{
			name: "empty",
			queue: Queue[int]{
				elements: []int{},
			},
			expected: Queue[int]{
				elements: []int{},
			},
			n: 0,
		},
		{
			name: "zero",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			expected: Queue[int]{
				elements: []int{1, 2, 3},
			},
			n: 0,
		},
		{
			name: "one",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			expected: Queue[int]{
				elements: []int{2, 3, 1},
			},
			n: 1,
		},
		{
			name: "two",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			expected: Queue[int]{
				elements: []int{3, 1, 2},
			},
			n: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.Rotate(tc.n)

			assert.Equal(t, tc.expected, tc.queue)
		})
	}
}

func Test_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		queue    Queue[int]
		expected Queue[int]
	}{
		{
			name: "empty",
			queue: Queue[int]{
				elements: []int{},
			},
			expected: Queue[int]{
				elements: []int{},
			},
		},
		{
			name: "one",
			queue: Queue[int]{
				elements: []int{1},
			},
			expected: Queue[int]{
				elements: []int{1},
			},
		},
		{
			name: "three",
			queue: Queue[int]{
				elements: []int{1, 2, 3},
			},
			expected: Queue[int]{
				elements: []int{3, 2, 1},
			},
		},
		{
			name: "four",
			queue: Queue[int]{
				elements: []int{1, 2, 3, 4},
			},
			expected: Queue[int]{
				elements: []int{4, 3, 2, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.Reverse()

			assert.Equal(t, tc.expected, tc.queue)
		})
	}
}
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Size(t *testing.T) {
	tests := []struct {
		name     string
		stack    Stack[int]
		expected int
	}{
		{
			name: "ok",
			stack: Stack[int]{
				values: []int{1, 2, 3},
				count:  3,
			},
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.stack.Size()
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Peek(t *testing.T) {
	tests := []struct {
		name        string
		stack       Stack[int]
		expected    int
		expectedErr bool
	}{
		{
			name: "ok",
			stack: Stack[int]{
				values: []int{1, 2, 4},
				count:  3,
			},
			expected: 4,
		},
		{
			name: "empty",
			stack: Stack[int]{
				values: []int{},
				count:  0,
			},
			expected:    0,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.stack.Peek()
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expectedErr, err != nil)
		})
	}
}

func Test_Pop(t *testing.T) {
	tests := []struct {
		name          string
		stack         Stack[int]
		expected      int
		expectedCount int
		expectedErr   bool
	}{
		{
			name: "ok",
			stack: Stack[int]{
				values: []int{1, 2, 4},
				count:  3,
			},
			expected:      4,
			expectedCount: 2,
			expectedErr:   false,
		},
		{
			name: "empty",
			stack: Stack[int]{
				values: []int{},
				count:  0,
			},
			expected:    0,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.stack.Pop()
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expectedCount, tc.stack.count)
			assert.Equal(t, tc.expectedErr, err != nil)
		})
	}
}

func Test_Push(t *testing.T) {
	tests := []struct {
		name          string
		stack         Stack[int]
		item          int
		expectedCount int
	}{
		{
			name: "ok",
			stack: Stack[int]{
				values: []int{1, 2, 3},
				count:  3,
			},
			item:          4,
			expectedCount: 4,
		},
		{
			name: "empty",
			stack: Stack[int]{
				values: []int{},
				count:  0,
			},
			item:          5,
			expectedCount: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.Push(tc.item)
			assert.Equal(t, tc.expectedCount, tc.stack.count)
			assert.Equal(t, tc.item, tc.stack.values[len(tc.stack.values)-1])
		})
	}
}

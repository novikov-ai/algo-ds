package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	tests := []struct {
		name      string
		ascending bool
		items     []int
		expected  []int
	}{
		{
			name:      "ascending single",
			ascending: true,
			items:     []int{5},
			expected:  []int{5},
		},
		{
			name:      "ascending ordered",
			ascending: true,
			items:     []int{3, 1, 4, 1, 5},
			expected:  []int{1, 1, 3, 4, 5},
		},
		{
			name:      "ascending reverse input",
			ascending: true,
			items:     []int{5, 4, 3, 2, 1},
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "ascending duplicates",
			ascending: true,
			items:     []int{2, 2, 2},
			expected:  []int{2, 2, 2},
		},
		{
			name:      "descending ordered",
			ascending: false,
			items:     []int{1, 3, 2, 5, 4},
			expected:  []int{5, 4, 3, 2, 1},
		},
		{
			name:      "descending reverse input",
			ascending: false,
			items:     []int{1, 2, 3, 4, 5},
			expected:  []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := New[int](tc.ascending)
			for _, v := range tc.items {
				l.Add(v)
			}

			got := toSlice(l)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Delete(t *testing.T) {
	tests := []struct {
		name      string
		ascending bool
		initial   []int
		delete    int
		expected  []int
	}{
		{
			name:      "ascending delete head",
			ascending: true,
			initial:   []int{1, 2, 3},
			delete:    1,
			expected:  []int{2, 3},
		},
		{
			name:      "ascending delete tail",
			ascending: true,
			initial:   []int{1, 2, 3},
			delete:    3,
			expected:  []int{1, 2},
		},
		{
			name:      "ascending delete middle",
			ascending: true,
			initial:   []int{1, 2, 3},
			delete:    2,
			expected:  []int{1, 3},
		},
		{
			name:      "ascending delete only element",
			ascending: true,
			initial:   []int{1},
			delete:    1,
			expected:  []int{},
		},
		{
			name:      "ascending delete missing",
			ascending: true,
			initial:   []int{1, 2, 3},
			delete:    9,
			expected:  []int{1, 2, 3},
		},
		{
			name:      "ascending delete first duplicate",
			ascending: true,
			initial:   []int{2, 2, 3},
			delete:    2,
			expected:  []int{2, 3},
		},
		{
			name:      "descending delete head",
			ascending: false,
			initial:   []int{3, 2, 1},
			delete:    3,
			expected:  []int{2, 1},
		},
		{
			name:      "descending delete tail",
			ascending: false,
			initial:   []int{3, 2, 1},
			delete:    1,
			expected:  []int{3, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := New[int](tc.ascending)
			for _, v := range tc.initial {
				l.Add(v)
			}

			l.Delete(tc.delete)
			got := toSlice(l)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Find(t *testing.T) {
	tests := []struct {
		name      string
		ascending bool
		initial   []int
		find      int
		wantErr   bool
		expected  int
	}{
		{
			name:      "ascending found in middle",
			ascending: true,
			initial:   []int{1, 2, 3},
			find:      2,
			expected:  2,
		},
		{
			name:      "ascending found head",
			ascending: true,
			initial:   []int{1, 2, 3},
			find:      1,
			expected:  1,
		},
		{
			name:      "ascending found tail",
			ascending: true,
			initial:   []int{1, 2, 3},
			find:      3,
			expected:  3,
		},
		{
			name:      "ascending not found",
			ascending: true,
			initial:   []int{1, 2, 3},
			find:      9,
			wantErr:   true,
		},
		{
			name:      "ascending empty list",
			ascending: true,
			initial:   []int{},
			find:      1,
			wantErr:   true,
		},
		{
			name:      "descending found",
			ascending: false,
			initial:   []int{3, 2, 1},
			find:      2,
			expected:  2,
		},
		{
			name:      "descending not found",
			ascending: false,
			initial:   []int{3, 2, 1},
			find:      9,
			wantErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := New[int](tc.ascending)
			for _, v := range tc.initial {
				l.Add(v)
			}

			got, err := l.Find(tc.find)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got.value)
		})
	}
}

func toSlice(l *OrderedList[int]) []int {
	result := make([]int, 0, l.Count())
	current := l.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}
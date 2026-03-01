package main

import (
	// "os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		capacity int
	}{
		{
			name:     "init16",
			count:    0,
			capacity: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := DynArray[int]{}
			da.Init()

			assert.Equal(t, tt.capacity, da.capacity)
			assert.Equal(t, tt.count, da.count)
		})
	}
}

func TestMakeArray(t *testing.T) {
	tests := []struct {
		name     string
		withInit func(da *DynArray[int])
		expected []int
	}{
		{
			name:     "without expand",
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "expand with copy",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)

				num := 42
				da.count = 3
				for i := 0; i < da.count; i++ {
					if num == 51 {
						continue
					}
					da.array[i] = num
					num += 3
				}
			},
			expected: []int{42, 45, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := DynArray[int]{}

			if tt.withInit != nil {
				tt.withInit(&da)
			}

			da.MakeArray(20)

			assert.Equal(t, tt.expected, da.array)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name             string
		withInit         func(da *DynArray[int])
		item             int
		index            int
		expectedCount    int
		expectedCapacity int
		expected         []int
		expectedErr      bool
	}{
		{
			name: "expand cap",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
				count := 1
				for i := range da.array {
					da.array[i] = count
					count++
				}
				da.count = 16
			},
			item:             42,
			index:            8,
			expectedCount:    17,
			expectedCapacity: 32,
			expected:         []int{1, 2, 3, 4, 5, 6, 7, 8, 42, 9, 10, 11, 12, 13, 14, 15, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "in the middle",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
				da.array[0] = 1
				da.array[1] = 2
				da.array[2] = 3
				da.array[3] = 4
				da.count = 4
			},
			item:             42,
			index:            2,
			expectedCount:    5,
			expectedCapacity: 16,
			expected:         []int{1, 2, 42, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "at the end",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
				da.array[0] = 1
				da.array[1] = 2
				da.count = 2
			},
			item:             42,
			index:            2,
			expectedCount:    3,
			expectedCapacity: 16,
			expected:         []int{1, 2, 42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "at the beginning",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
			},
			item:             42,
			index:            0,
			expectedCount:    1,
			expectedCapacity: 16,
			expected:         []int{42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "out of range",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
			},
			item:             42,
			index:            1,
			expectedCount:    0,
			expectedCapacity: 16,
			expected:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := DynArray[int]{}

			if tt.withInit != nil {
				tt.withInit(&da)
			}

			err := da.Insert(tt.item, tt.index)

			if tt.expectedErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.expected, da.array)
			assert.Equal(t, tt.expectedCapacity, da.capacity)
			assert.Equal(t, tt.expectedCount, da.count)
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name             string
		withInit         func(da *DynArray[int])
		index            int
		expectedCount    int
		expectedCapacity int
		expected         []int
		expectedErr      bool
	}{
		{
			name: "shrinked cap",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(32)
				count := 1
				for i := range da.array {
					da.array[i] = count
					count++
				}
				da.count = 16
			},
			index:            8,
			expectedCount:    15,
			expectedCapacity: 21,
			expected:         []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 16, 17, 18, 19, 20, 21},
			expectedErr:      false,
		},
		{
			name: "NO shrinking",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
				count := 1
				for i := range da.array {
					da.array[i] = count
					count++
				}
				da.count = 10
			},
			index:            8,
			expectedCount:    9,
			expectedCapacity: 16,
			expected:         []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 10, 11, 12, 13, 14, 15, 16},
			expectedErr:      false,
		},
		{
			name: "remove from the end",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
				da.array[0] = 1
				da.array[1] = 2
				da.count = 2
			},
			index:            1,
			expectedCount:    1,
			expectedCapacity: 16,
			expected:         []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "remove from the beginning",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
			},
			index:            0,
			expectedCount:    0,
			expectedCapacity: 16,
			expected:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      false,
		},
		{
			name: "out of range",
			withInit: func(da *DynArray[int]) {
				da.MakeArray(16)
			},
			index:            1,
			expectedCount:    0,
			expectedCapacity: 16,
			expected:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := DynArray[int]{}

			if tt.withInit != nil {
				tt.withInit(&da)
			}

			err := da.Remove(tt.index)

			if tt.expectedErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.expected, da.array)
			assert.Equal(t, tt.expectedCapacity, da.capacity)
			assert.Equal(t, tt.expectedCount, da.count)
		})
	}
}
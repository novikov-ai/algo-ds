package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Put(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		key      int
		nc       NativeCache[int]
		expected NativeCache[int]
	}{
		{
			name:  "put ok",
			value: "key1",
			key:   1,
			nc: func() NativeCache[int] {
				nc := New[int](13)
				return nc
			}(),
			expected: func() NativeCache[int] {
				nc := New[int](13)
				key := "key1"
				index := nc.HashFun(key)
				nc.slots[index] = key
				nc.values[index] = 1
				return nc
			}(),
		},
		{
			name:  "replace by index in full cache",
			value: "key1",
			key:   42,
			nc: func() NativeCache[int] {
				nc := New[int](5)
				nc.values = []int{1, 2, 3, 4, 5}
				nc.slots = []string{"1", "2", "3", "4", "5"}
				nc.hits = []int{3, 4, 7, 1, 22}
				return nc
			}(),
			expected: func() NativeCache[int] {
				nc := New[int](5)
				nc.values = []int{1, 2, 3, 42, 5}
				nc.slots = []string{"1", "2", "3", "key1", "5"}
				nc.hits = []int{3, 4, 7, 0, 22}
				return nc
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.nc.Put(tc.value, tc.key)
			assert.Equal(t, tc.expected, tc.nc)
			assert.Equal(t, tc.expected.hits, tc.nc.hits)
			assert.Equal(t, tc.expected.slots, tc.nc.slots)
			assert.Equal(t, tc.expected.values, tc.nc.values)
		})
	}
}

func Test_Get(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		key          string
		nc           NativeCache[int]
		expected     int
		expectedHits []int
		expectedErr  bool
	}{
		{
			name:  "get ok",
			value: "key1",
			nc: func() NativeCache[int] {
				nc := New[int](5)
				nc.Put("key1", 123)
				return nc
			}(),
			expected:     123,
			expectedHits: []int{0, 0, 1, 0, 0},
		},
		{
			name:  "many hits",
			value: "key1",
			nc: func() NativeCache[int] {
				nc := New[int](5)
				nc.Put("key1", 123)

				for range 5 {
					nc.Get("key1")
				}

				return nc
			}(),
			expected:     123,
			expectedHits: []int{0, 0, 6, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.nc.Get(tc.value)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expectedHits, tc.nc.hits)
		})
	}
}

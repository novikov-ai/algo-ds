package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashFun(t *testing.T) {
	tests := []struct {
		name     string
		nd       NativeDictionary[int]
		value    string
		expected int
	}{
		{
			name: "ok",
			nd: func() NativeDictionary[int] {
				return Init[int](13)
			}(),
			value:    "native",
			expected: 5,
		},
		{
			name: "reversed ok",
			nd: func() NativeDictionary[int] {
				return Init[int](13)
			}(),
			value:    "planer",
			expected: 0,
		},
		{
			name: "one",
			nd: func() NativeDictionary[int] {
				return Init[int](13)
			}(),
			value:    "o",
			expected: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.nd.HashFun(tc.value)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_IsKey(t *testing.T) {
	tests := []struct {
		name     string
		nd       NativeDictionary[int]
		value    string
		expected bool
	}{
		{
			name: "has",
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				nd.Put("key1", 1)
				nd.Put("key2", 2)
				nd.Put("key3", 3)
				return nd
			}(),
			value:    "key2",
			expected: true,
		},
		{
			name: "has not",
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				nd.Put("key1", 1)
				nd.Put("key2", 2)
				nd.Put("key3", 3)
				return nd
			}(),
			value:    "key5",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.nd.IsKey(tc.value)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Put(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		key      int
		nd       NativeDictionary[int]
		expected NativeDictionary[int]
	}{
		{
			name:  "put ok",
			value: "key1",
			key:   1,
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				return nd
			}(),
			expected: func() NativeDictionary[int] {
				nd := Init[int](13)
				key := "key1"
				index := nd.HashFun(key)
				nd.slots[index] = key
				nd.values[index] = 1
				return nd
			}(),
		},
		{
			name:  "put with collision",
			value: "kye1",
			key:   3,
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				key := "key1"
				index := nd.HashFun(key)
				nd.slots[index] = key
				nd.values[index] = 1
				return nd
			}(),
			expected: func() NativeDictionary[int] {
				nd := Init[int](13)
				key := "kye1"
				index := nd.HashFun(key)
				nd.slots[index] = key
				nd.values[index] = 3
				return nd
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.nd.Put(tc.value, tc.key)
			assert.Equal(t, tc.expected, tc.nd)
		})
	}
}

func Test_Get(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		nd       NativeDictionary[int]
		expected int
		hasError bool
	}{
		{
			name: "found",
			key:  "key1",
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				nd.Put("key1", 1)
				return nd
			}(),
			expected: 1,
			hasError: false,
		},
		{
			name: "not found",
			key:  "key2",
			nd: func() NativeDictionary[int] {
				nd := Init[int](13)
				nd.Put("key1", 1)
				return nd
			}(),
			expected: 0,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.nd.Get(tc.key)

			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

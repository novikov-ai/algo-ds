package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashFun(t *testing.T) {
	tests := []struct {
		name     string
		ht       HashTable
		input    string
		expected int
	}{
		{
			name: "max len",
			ht: func() HashTable {
				return Init(19, 3)
			}(),
			input:    "1234567891234567891",
			expected: 0,
		},
		{
			name: "max len + 1",
			ht: func() HashTable {
				return Init(19, 3)
			}(),
			input:    "12345678912345678912",
			expected: 1,
		},
		{
			name: "basic",
			ht: func() HashTable {
				return Init(19, 3)
			}(),
			input:    "value",
			expected: 5,
		},
		{
			name: "equals len",
			ht: func() HashTable {
				return Init(19, 3)
			}(),
			input:    "lague",
			expected: 5,
		},
		{
			name: "empty",
			ht: func() HashTable {
				return Init(19, 3)
			}(),
			input:    "",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ht.HashFun(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_SeekSlot(t *testing.T) {
	tests := []struct {
		name     string
		ht       HashTable
		input    string
		expected int
	}{
		{
			name: "first slot",
			ht: func() HashTable {
				ht := Init(19, 3)
				return ht
			}(),
			input:    "value",
			expected: 5,
		},
		{
			name: "second slot",
			ht: func() HashTable {
				ht := Init(19, 3)
				ht.slots[5] = "occupied"
				return ht
			}(),
			input:    "value",
			expected: 8,
		},
		{
			name: "jump over and start again",
			ht: func() HashTable {
				ht := Init(19, 3)
				ht.slots[5] = "occupied"
				ht.slots[8] = "occupied again"
				ht.slots[11] = "occupied again again"
				ht.slots[14] = "occupied again again again"
				ht.slots[17] = "occupied again again again again"
				return ht
			}(),
			input:    "value",
			expected: 1,
		},
		{
			name: "no empty slots",
			ht: func() HashTable {
				ht := Init(19, 3)
				for i := range ht.size {
					ht.slots[i] = "occupied"
				}
				return ht
			}(),
			input:    "value",
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ht.SeekSlot(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Put(t *testing.T) {
	tests := []struct {
		name     string
		ht       HashTable
		input    string
		expected int
	}{
		{
			name: "success",
			ht: func() HashTable {
				ht := Init(19, 3)
				return ht
			}(),
			input:    "value",
			expected: 5,
		},
		{
			name: "can't put",
			ht: func() HashTable {
				ht := Init(19, 3)
				for i := range ht.size {
					ht.slots[i] = "occupied"
				}
				return ht
			}(),
			input:    "value",
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ht.Put(tc.input)
			assert.Equal(t, tc.expected, got)
			if got != -1 {
				assert.Equal(t, tc.input, tc.ht.slots[tc.expected])
			}
		})
	}
}

func Test_Find(t *testing.T) {
	tests := []struct {
		name     string
		ht       HashTable
		input    string
		expected int
	}{
		{
			name: "not found",
			ht: func() HashTable {
				ht := Init(19, 3)
				return ht
			}(),
			input:    "value",
			expected: -1,
		},
		{
			name: "found",
			ht: func() HashTable {
				ht := Init(19, 3)
				ht.Put("value")
				return ht
			}(),
			input:    "value",
			expected: 5,
		},
		{
			name: "second slot",
			ht: func() HashTable {
				ht := Init(19, 3)
				ht.slots[5] = "value"
				ht.slots[8] = "valeu"
				return ht
			}(),
			input:    "valeu",
			expected: 8,
		},
		{
			name: "jump over and start again",
			ht: func() HashTable {
				ht := Init(19, 3)
				ht.slots[5] = "value"
				ht.slots[8] = "valeu"
				ht.slots[11] = "vaelu"
				ht.slots[14] = "valeu"
				ht.slots[17] = "avlue"
				ht.slots[1] = "aaaaa"
				return ht
			}(),
			input:    "aaaaa",
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ht.Find(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		element  string
		expected map[string]string
	}{
		{
			name: "empty",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				return &ps
			}(),
			element: "test",
			expected: map[string]string{
				"test": "",
			},
		},
		{
			name: "try duplicate",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				return &ps
			}(),
			element: "test",
			expected: map[string]string{
				"test": "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ps.Put(tc.element)
			assert.Equal(t, tc.expected, tc.ps.values)
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name           string
		ps             *PowerSet[string]
		element        string
		expected       map[string]string
		expectedResult bool
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				return &ps
			}(),
			element:        "test",
			expected:       map[string]string{},
			expectedResult: true,
		},
		{
			name: "not ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test123"] = ""
				return &ps
			}(),
			element: "test",
			expected: map[string]string{
				"test123": "",
			},
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.Remove(tc.element)
			assert.Equal(t, tc.expected, tc.ps.values)
			assert.Equal(t, tc.expectedResult, got)
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		ps2      PowerSet[string]
		expected PowerSet[string]
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
		},
		{
			name: "not ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test4"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				return ps
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.Intersection(tc.ps2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		ps2      PowerSet[string]
		expected PowerSet[string]
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test4"] = ""
				ps.values["test5"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				ps.values["test4"] = ""
				ps.values["test5"] = ""
				return ps
			}(),
		},
		{
			name: "one is empty",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
		},
		{
			name: "another is empty",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.Union(tc.ps2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		ps2      PowerSet[string]
		expected PowerSet[string]
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test4"] = ""
				ps.values["test5"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
		},
		{
			name: "second is empty",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
		},
		{
			name: "nothing",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: func() PowerSet[string] {
				ps := Init[string]()
				return ps
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.Difference(tc.ps2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestIsSubset(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		ps2      PowerSet[string]
		expected bool
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: true,
		},
		{
			name: "not ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: false,
		},
		{
			name: "less than param",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				ps.values["test4"] = ""
				ps.values["test5"] = ""
				return ps
			}(),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.IsSubset(tc.ps2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		ps       *PowerSet[string]
		ps2      PowerSet[string]
		expected bool
	}{
		{
			name: "ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: true,
		},
		{
			name: "not ok",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test12"] = ""
				ps.values["test3"] = ""
				return ps
			}(),
			expected: false,
		},
		{
			name: "big diff",
			ps: func() *PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				return &ps
			}(),
			ps2: func() PowerSet[string] {
				ps := Init[string]()
				ps.values["test"] = ""
				ps.values["test2"] = ""
				ps.values["test3"] = ""
				ps.values["test4"] = ""
				return ps
			}(),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.ps.Equals(tc.ps2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

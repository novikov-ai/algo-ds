package main

import "testing"

var testStrings = []string{
	"0123456789",
	"1234567890",
	"2345678901",
	"3456789012",
	"4567890123",
	"5678901234",
	"6789012345",
	"7890123456",
	"8901234567",
	"9012345678",
}

func TestAddAndIsValue(t *testing.T) {
	bf := New(32)

	for _, s := range testStrings {
		bf.Add(s)
	}

	for _, s := range testStrings {
		if !bf.IsValue(s) {
			t.Errorf("expected %q to be in filter, but it was not", s)
		}
	}
}

func TestIsValueNotAdded(t *testing.T) {
	bf := New(32)

	bf.Add("0123456789")

	absent := []string{"hello", "world", "golang", "bloom"}
	for _, s := range absent {
		if bf.IsValue(s) {
			t.Logf("false positive for %q (acceptable for Bloom filter)", s)
		}
	}
}

func TestEmptyFilter(t *testing.T) {
	bf := New(32)

	for _, s := range testStrings {
		if bf.IsValue(s) {
			t.Errorf("empty filter should not contain %q", s)
		}
	}
}

func TestAddOneCheckAll(t *testing.T) {
	for _, added := range testStrings {
		bf := New(32)
		bf.Add(added)

		if !bf.IsValue(added) {
			t.Errorf("added %q but IsValue returned false", added)
		}
	}
}

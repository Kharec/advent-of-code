package main

import "testing"

var input = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func TestPartOne(t *testing.T) {
	expected := 11
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

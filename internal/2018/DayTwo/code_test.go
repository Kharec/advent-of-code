package daytwo

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	expected := 12
	got := PartOne(input)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	expected := "fgij"

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	PartTwo(input)

	_ = w.Close()
	os.Stdout = oldStdout

	var buffer bytes.Buffer
	_, _ = buffer.ReadFrom(r)
	got := strings.TrimSpace(buffer.String())

	if got != expected {
		t.Errorf("PartTwo() printed %q, want %q", got, expected)
	}
}

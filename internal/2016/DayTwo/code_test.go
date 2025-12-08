package daytwo

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var testInput = []string{"ULL", "RRDDD", "LURDL", "UUUUD"}

func TestPartOne(t *testing.T) {
	expected := 1985
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	expected := "5DB3"

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	PartTwo(testInput)

	_ = w.Close()
	os.Stdout = oldStdout

	var buffer bytes.Buffer
	_, _ = buffer.ReadFrom(r)
	got := strings.TrimSpace(buffer.String())

	if got != expected {
		t.Errorf("PartTwo() printed %q, want %q", got, expected)
	}
}

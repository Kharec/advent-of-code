package dayfive

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var testInput = "abc"

func TestPartOne(t *testing.T) {
	expected := "18f47a30"

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	PartOne(testInput)

	_ = w.Close()
	os.Stdout = oldStdout

	var buffer bytes.Buffer
	_, _ = buffer.ReadFrom(r)
	got := strings.TrimSpace(buffer.String())

	if got != expected {
		t.Errorf("PartOne() printed %q, want %q", got, expected)
	}
}

package dayfourteen

import "testing"

var testInput = []string{
	"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
	"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
}

func TestPartOne(t *testing.T) {
	expected := 1120
	got := calculateMaxDistance(testInput, 1000)
	if got != expected {
		t.Errorf("calculateMaxDistance(testInput, 1000) = %d, want %d", got, expected)
	}
}

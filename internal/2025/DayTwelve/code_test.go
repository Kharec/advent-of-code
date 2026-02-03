package daytwelve

import "testing"

var testInput = []string{
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

func TestPartOne(t *testing.T) {
	expected := 2
	got := PartOne(testInput)
	if got != expected {
		t.Errorf("PartOne() = %d, want %d", got, expected)
	}
}

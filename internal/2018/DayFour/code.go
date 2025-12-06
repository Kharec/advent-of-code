package dayfour

import (
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"advent-of-code/internal/registry"
)

var guardRegex = regexp.MustCompile(`Guard #(\d+) begins shift`)

func init() {
	registry.Register("2018D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

type record struct {
	timestamp time.Time
	action    string
	guardID   int
}

func PartOne(data []string) int {
	records := make([]record, 0, len(data))
	timeLayout := "2006-01-02 15:04"

	for _, line := range data {
		timeStr, action, _ := strings.Cut(line[1:], "] ")
		timestamp, _ := time.Parse(timeLayout, timeStr)
		guardID := -1
		if matches := guardRegex.FindStringSubmatch(action); matches != nil {
			guardID, _ = strconv.Atoi(matches[1])
		}
		records = append(records, record{timestamp, action, guardID})
	}

	slices.SortFunc(records, func(a, b record) int {
		return a.timestamp.Compare(b.timestamp)
	})

	guardSleepMinutes := make(map[int]int)
	guardMinuteCount := make(map[int]map[int]int)
	currentGuard := -1

	for i := range records {
		if records[i].guardID != -1 {
			currentGuard = records[i].guardID
			if guardMinuteCount[currentGuard] == nil {
				guardMinuteCount[currentGuard] = make(map[int]int)
			}
		}
		if records[i].action == "falls asleep" && i+1 < len(records) {
			start, end := records[i].timestamp.Minute(), records[i+1].timestamp.Minute()
			guardSleepMinutes[currentGuard] += end - start
			for m := start; m < end; m++ {
				guardMinuteCount[currentGuard][m]++
			}
		}
	}

	maxGuard, maxMinutes := -1, 0
	for id, minutes := range guardSleepMinutes {
		if minutes > maxMinutes {
			maxGuard, maxMinutes = id, minutes
		}
	}

	maxMinute, maxCount := -1, 0
	for m, count := range guardMinuteCount[maxGuard] {
		if count > maxCount {
			maxMinute, maxCount = m, count
		}
	}

	return maxGuard * maxMinute
}

func PartTwo(data []string) int {
	return 0
}

package dayfourteen

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2015D14", ParseInput, PartOne, PartTwo)
}

type Reindeer struct {
	Speed     int
	FlyTime   int
	RestTime  int
	CycleTime int
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func calculateMaxDistance(data []string, time int) int {
	maxDistance := 0
	pattern := regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)

	for _, line := range data {
		matches := pattern.FindStringSubmatch(line)
		var reindeer Reindeer
		reindeer.Speed, _ = strconv.Atoi(matches[1])
		reindeer.FlyTime, _ = strconv.Atoi(matches[2])
		reindeer.RestTime, _ = strconv.Atoi(matches[3])
		reindeer.CycleTime = reindeer.FlyTime + reindeer.RestTime

		fullCycles := time / reindeer.CycleTime
		distance := fullCycles * reindeer.Speed * reindeer.FlyTime

		remainingTime := time % reindeer.CycleTime
		if remainingTime > reindeer.FlyTime {
			distance += reindeer.Speed * reindeer.FlyTime
		} else {
			distance += reindeer.Speed * remainingTime
		}

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func PartOne(data []string) int {
	return calculateMaxDistance(data, 2503)
}

func PartTwo(data []string) int {
	return 0
}

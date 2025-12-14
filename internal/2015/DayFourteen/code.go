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

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func calculateMaxDistance(data []string, time int) int {
	maxDistance := 0
	pattern := regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)

	for _, line := range data {
		matches := pattern.FindStringSubmatch(line)

		speed, _ := strconv.Atoi(matches[1])
		flyTime, _ := strconv.Atoi(matches[2])
		restTime, _ := strconv.Atoi(matches[3])
		cycleTime := flyTime + restTime

		fullCycles := time / cycleTime
		distance := fullCycles * speed * flyTime

		remainingTime := time % cycleTime
		distance += speed * min(remainingTime, flyTime)

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

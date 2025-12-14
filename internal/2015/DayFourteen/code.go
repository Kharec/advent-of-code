package dayfourteen

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reindeerPattern = regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.`)

const raceTime = 2503

type Reindeer struct {
	Speed    int
	FlyTime  int
	RestTime int
	Points   int
}

func init() {
	registry.Register("2015D14", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func parseReindeer(line string) Reindeer {
	matches := reindeerPattern.FindStringSubmatch(line)
	speed, _ := strconv.Atoi(matches[1])
	flyTime, _ := strconv.Atoi(matches[2])
	restTime, _ := strconv.Atoi(matches[3])
	return Reindeer{Speed: speed, FlyTime: flyTime, RestTime: restTime}
}

func parseReindeers(data []string) []Reindeer {
	reindeers := make([]Reindeer, 0, len(data))
	for _, line := range data {
		reindeers = append(reindeers, parseReindeer(line))
	}
	return reindeers
}

func (reindeer Reindeer) distanceAtTime(time int) int {
	cycleTime := reindeer.FlyTime + reindeer.RestTime
	fullCycles := time / cycleTime
	distance := fullCycles * reindeer.Speed * reindeer.FlyTime
	remainingTime := time % cycleTime
	distance += reindeer.Speed * min(remainingTime, reindeer.FlyTime)
	return distance
}

func calculateMaxDistance(data []string, time int) int {
	reindeers := parseReindeers(data)
	maxDistance := 0

	for _, reindeer := range reindeers {
		distance := reindeer.distanceAtTime(time)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func calculateMaxPoints(data []string, time int) int {
	reindeers := parseReindeers(data)

	for second := 1; second <= time; second++ {
		maxDistance := 0
		distances := make([]int, len(reindeers))

		for idx := range reindeers {
			distance := reindeers[idx].distanceAtTime(second)
			distances[idx] = distance
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		for idx := range reindeers {
			if distances[idx] == maxDistance {
				reindeers[idx].Points++
			}
		}
	}

	maxPoints := 0
	for idx := range reindeers {
		if reindeers[idx].Points > maxPoints {
			maxPoints = reindeers[idx].Points
		}
	}

	return maxPoints
}

func PartOne(data []string) int {
	return calculateMaxDistance(data, raceTime)
}

func PartTwo(data []string) int {
	return calculateMaxPoints(data, raceTime)
}

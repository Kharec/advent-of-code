package dayfour

import (
	"advent-of-code/internal/registry"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2016D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

func isValidRoom(encryptedName string, expectedChecksum string) bool {
	letterFrequency := [26]int{}

	for _, character := range encryptedName {
		if character >= 'a' && character <= 'z' {
			letterFrequency[character-'a']++
		}
	}

	type letterFrequencyPair struct {
		letter rune
		count  int
	}

	letterFrequencyPairs := make([]letterFrequencyPair, 0, 26)
	for index, frequency := range letterFrequency {
		if frequency > 0 {
			letterFrequencyPairs = append(letterFrequencyPairs, letterFrequencyPair{
				letter: rune('a' + index),
				count:  frequency,
			})
		}
	}

	sort.Slice(letterFrequencyPairs, func(i, j int) bool {
		if letterFrequencyPairs[i].count != letterFrequencyPairs[j].count {
			return letterFrequencyPairs[i].count > letterFrequencyPairs[j].count
		}
		return letterFrequencyPairs[i].letter < letterFrequencyPairs[j].letter
	})

	if len(letterFrequencyPairs) < 5 {
		return false
	}

	for index := range 5 {
		if letterFrequencyPairs[index].letter != rune(expectedChecksum[index]) {
			return false
		}
	}

	return true
}

func decryptRoomName(encryptedName string, sectorID int) string {
	result := strings.Builder{}
	result.Grow(len(encryptedName))
	shift := sectorID % 26
	for _, char := range encryptedName {
		if char == '-' {
			result.WriteByte(' ')
		} else if char >= 'a' && char <= 'z' {
			shifted := ((int(char-'a') + shift) % 26) + 'a'
			result.WriteByte(byte(shifted))
		}
	}
	return result.String()
}

var roomPattern = regexp.MustCompile(`^(.+)-(\d+)(?:\[([a-z]{5})\])?$`)

func PartOne(data []string) int {
	sum := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		matches := roomPattern.FindStringSubmatch(line)
		if len(matches) < 4 || matches[3] == "" {
			continue
		}
		encryptedName := matches[1]
		sectorIdentifier, _ := strconv.Atoi(matches[2])
		actualChecksum := matches[3]
		if isValidRoom(encryptedName, actualChecksum) {
			sum += sectorIdentifier
		}
	}
	return sum
}

func PartTwo(data []string) int {
	for _, line := range data {
		matches := roomPattern.FindStringSubmatch(line)
		encryptedName := matches[1]
		sectorIdentifier, _ := strconv.Atoi(matches[2])
		checksum := matches[3]

		if !isValidRoom(encryptedName, checksum) {
			continue
		}

		decrypted := decryptRoomName(encryptedName, sectorIdentifier)
		if strings.Contains(decrypted, "northpole") {
			return sectorIdentifier
		}
	}
	return 0
}

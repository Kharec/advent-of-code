package dayfour

import (
	"advent-of-code/internal/registry"
	"bytes"
	"crypto/md5"
	"os"
	"strconv"
)

func init() {
	registry.Register("2015D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []byte {
	content, _ := os.ReadFile(filepath)
	return bytes.TrimSpace(content)
}

func findHashWithZeroes(data []byte, zeroes int) int {
	buffer := make([]byte, len(data), len(data)+10)
	copy(buffer, data)
	for idx := 1; ; idx++ {
		buffer = buffer[:len(data)]
		buffer = strconv.AppendInt(buffer, int64(idx), 10)
		hash := md5.Sum(buffer)
		switch zeroes {
		case 5:
			if hash[0] == 0 && hash[1] == 0 && hash[2]&0xF0 == 0 {
				return idx
			}
		case 6:
			if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
				return idx
			}
		}
	}
}

func PartOne(data []byte) int {
	return findHashWithZeroes(data, 5)
}

func PartTwo(data []byte) int {
	return findHashWithZeroes(data, 6)
}

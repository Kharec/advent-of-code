package dayfive

import (
	"advent-of-code/internal/registry"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func init() {
	registry.Register("2016D5", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func PartOne(data string) int {
	doorIDBytes := []byte(data)
	doorIDLen := len(doorIDBytes)
	input := make([]byte, doorIDLen, doorIDLen+20)
	copy(input, doorIDBytes)
	password := make([]byte, 0, 8)
	index := 0
	hexChars := "0123456789abcdef"

	for len(password) < 8 {
		indexBytes := strconv.AppendInt(input[:doorIDLen], int64(index), 10)
		hash := md5.Sum(indexBytes)

		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			char := hexChars[hash[2]&0x0F]
			password = append(password, char)
		}
		index++
	}

	fmt.Println(string(password))
	return 0
}

func PartTwo(data string) int {
	doorIDBytes := []byte(data)
	doorIDLen := len(doorIDBytes)
	input := make([]byte, doorIDLen, doorIDLen+20)
	copy(input, doorIDBytes)
	password := make([]byte, 8)
	filled := make([]bool, 8)
	filledCount := 0
	index := 0
	hexChars := "0123456789abcdef"

	for filledCount < 8 {
		indexBytes := strconv.AppendInt(input[:doorIDLen], int64(index), 10)
		hash := md5.Sum(indexBytes)

		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			position := int(hash[2] & 0x0F)
			if position < 8 && !filled[position] {
				char := hexChars[hash[3]>>4]
				password[position] = char
				filled[position] = true
				filledCount++
			}
		}
		index++
	}

	fmt.Println(string(password))
	return 0
}

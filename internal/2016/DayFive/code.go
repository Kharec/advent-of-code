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
	password := make([]byte, 0, 8)
	index := 0
	hexChars := "0123456789abcdef"
	var buffer []byte

	for len(password) < 8 {
		buffer = strconv.AppendInt(buffer[:0], int64(index), 10)
		input := append(doorIDBytes, buffer...)
		hash := md5.Sum(input)

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
	return 0
}

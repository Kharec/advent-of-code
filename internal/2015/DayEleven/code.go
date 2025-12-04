package dayeleven

import (
	"os"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2015D11", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) string {
	content, _ := os.ReadFile(filepath)
	return string(content)
}

func PartOne(data string) int {
	bytes := []byte(data)

	increment := func() {
		for idx := len(bytes) - 1; idx >= 0; idx-- {
			if bytes[idx] == 'z' {
				bytes[idx] = 'a'
			} else {
				bytes[idx]++
				if bytes[idx] == 'i' || bytes[idx] == 'o' || bytes[idx] == 'l' {
					bytes[idx]++
				}
				for j := idx + 1; j < len(bytes); j++ {
					bytes[j] = 'a'
				}
				return
			}
		}
	}

	increment()

	for {
		forbiddenPosition := -1
		for idx := range bytes {
			if bytes[idx] == 'i' || bytes[idx] == 'o' || bytes[idx] == 'l' {
				forbiddenPosition = idx
				break
			}
		}

		if forbiddenPosition != -1 {
			bytes[forbiddenPosition]++
			if bytes[forbiddenPosition] == 'i' || bytes[forbiddenPosition] == 'o' || bytes[forbiddenPosition] == 'l' {
				bytes[forbiddenPosition]++
			}
			for j := forbiddenPosition + 1; j < len(bytes); j++ {
				bytes[j] = 'a'
			}
			continue
		}

		hasStraight := false
		for idx := 0; idx < len(bytes)-2; idx++ {
			if bytes[idx+1] == bytes[idx]+1 && bytes[idx+2] == bytes[idx]+2 {
				hasStraight = true
				break
			}
		}

		pairChars := make(map[byte]bool)
		for idx := 0; idx < len(bytes)-1; idx++ {
			if bytes[idx] == bytes[idx+1] {
				pairChars[bytes[idx]] = true
				idx++
			}
		}

		if hasStraight && len(pairChars) >= 2 {
			break
		}

		increment()
	}

	println(string(bytes))
	return 0
}

func PartTwo(data string) int {
	return 0
}

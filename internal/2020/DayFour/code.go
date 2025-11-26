package dayfour

import (
	"advent-of-code/internal/registry"
	"os"
	"strconv"
	"strings"
)

func init() {
	registry.Register("2020D4", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)

	text := string(content)
	passports := strings.Split(text, "\n\n")

	result := make([]string, 0, len(passports))
	for _, passport := range passports {
		passport = strings.TrimSpace(passport)
		if passport == "" {
			continue
		}
		passport = strings.ReplaceAll(passport, "\n", " ")
		passport = strings.Join(strings.Fields(passport), " ")
		result = append(result, passport)
	}
	return result
}

func PartOne(data []string) int {
	count := 0
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, passport := range data {
		fields := strings.Fields(passport)
		present := make(map[string]bool)
		for _, field := range fields {
			if idx := strings.Index(field, ":"); idx != -1 {
				present[field[:idx]] = true
			}
		}
		missing := false
		for _, field := range required {
			if !present[field] {
				missing = true
				break
			}
		}
		if !missing {
			count++
		}
	}
	return count
}

func PartTwo(data []string) int {
	isNumberInRange := func(value string, min, max int, digits int) bool {
		if len(value) != digits {
			return false
		}
		n, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return n >= min && n <= max
	}

	isValidHeight := func(value string) bool {
		if len(value) < 3 {
			return false
		}
		unit := value[len(value)-2:]
		number := value[:len(value)-2]
		n, err := strconv.Atoi(number)
		if err != nil {
			return false
		}
		if unit == "cm" {
			return n >= 150 && n <= 193
		}
		if unit == "in" {
			return n >= 59 && n <= 76
		}
		return false
	}

	isValidHairColor := func(value string) bool {
		if len(value) != 7 || value[0] != '#' {
			return false
		}
		for _, c := range value[1:] {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				return false
			}
		}
		return true
	}

	isValidPID := func(value string) bool {
		if len(value) != 9 {
			return false
		}
		for _, c := range value {
			if c < '0' || c > '9' {
				return false
			}
		}
		return true
	}

	validECL := map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true,
		"grn": true, "hzl": true, "oth": true,
	}

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	count := 0

	for _, passport := range data {
		fields := map[string]string{}
		for field := range strings.FieldsSeq(passport) {
			parts := strings.SplitN(field, ":", 2)
			if len(parts) == 2 {
				fields[parts[0]] = parts[1]
			}
		}

		passportValidators := map[string]func(string) bool{
			"byr": func(v string) bool { return isNumberInRange(v, 1920, 2002, 4) },
			"iyr": func(v string) bool { return isNumberInRange(v, 2010, 2020, 4) },
			"eyr": func(v string) bool { return isNumberInRange(v, 2020, 2030, 4) },
			"hgt": isValidHeight,
			"hcl": isValidHairColor,
			"ecl": func(v string) bool { return validECL[v] },
			"pid": isValidPID,
		}

		valid := true
		for _, key := range required {
			value, ok := fields[key]
			if !ok || !passportValidators[key](value) {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}


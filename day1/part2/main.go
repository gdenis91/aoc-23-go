package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	calibrationSum := 0
	for _, line := range lines {
		v, err := calibrationValue(line)
		if err != nil {
			panic(fmt.Errorf("get calibration value: %w", err))
		}
		calibrationSum += v
	}
	fmt.Println("Calibration sum:", calibrationSum)
}

func calibrationValue(line string) (int, error) {
	fmt.Println("line:", line)
	// Replace with the digit surrounded by the original string on both ends
	// This should handle overlapping words like "twone"
	// "twone" -> two2twone1one
	replacerValues := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
	for k, v := range replacerValues {
		line = strings.ReplaceAll(line, k, v)
	}
	fmt.Println("replaced line:", line)

	digits := make([]string, 0)
	for _, r := range line {
		if unicode.IsDigit(r) {
			digits = append(digits, string(r))
		}
	}
	fmt.Println("digits:", digits)
	if len(digits) == 0 {
		return 0, nil
	}
	first, last := digits[0], digits[len(digits)-1]
	cVal := fmt.Sprint(first, last)
	fmt.Println("calibration value:", cVal)
	return strconv.Atoi(cVal)
}

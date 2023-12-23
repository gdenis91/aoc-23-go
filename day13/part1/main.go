package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	patterns := mustParseInput(input)
	total := 0
outer:
	for _, p := range patterns {
		fmt.Println("Section:")
		fmt.Println(p.raw)

		// Check vertical reflection
		// For each center line check if there is a reflection
		// Start at index one so we have at least one col to reflect on the left
		// End one early so we have at least one col to reflect on the right
		for i := 1; i < len(p.cols); i++ {
			fmt.Println("Checking center:", i)
			if isReflected(p.cols, i) {
				total += i
				fmt.Println("Reflection along vertical line:", i)
				continue outer
			}
		}

		fmt.Println("No reflection along vertical line, checking horizontal")
		for i := 1; i < len(p.rows); i++ {
			fmt.Println("Checking center:", i)
			if isReflected(p.rows, i) {
				total += (i * 100)
				fmt.Println("Reflection along horizontal line:", i)
				continue outer
			}
		}
	}
	fmt.Println("Total:", total)
}

func isReflected(value [][]rune, reflectionPoint int) bool {
	leftOfPoint := reflectionPoint
	rightOfPoint := len(value) - reflectionPoint
	elemsToCheck := int(math.Min(float64(leftOfPoint), float64(rightOfPoint)))
	leftPtr := reflectionPoint - 1
	rightPtr := reflectionPoint
	for i := 0; i < elemsToCheck; i++ {
		if !slices.Equal(value[leftPtr], value[rightPtr]) {
			return false
		}
		leftPtr--
		rightPtr++
	}
	return true
}

func mustParseInput(input string) []pattern {
	patterns := make([]pattern, 0)
	for _, section := range strings.Split(input, "\n\n") {
		lines := strings.Split(section, "\n")
		numRows := len(lines)
		numCols := len([]rune(lines[0]))
		pattern := pattern{
			raw:  section,
			rows: make([][]rune, numRows),
			cols: make([][]rune, numCols),
		}
		for rowNo, line := range lines {
			pattern.rows[rowNo] = []rune(line)
			for colNo, v := range line {
				pattern.cols[colNo] = append(pattern.cols[colNo], v)
			}
		}
		patterns = append(patterns, pattern)
	}
	return patterns
}

type pattern struct {
	raw  string
	rows [][]rune
	cols [][]rune
}

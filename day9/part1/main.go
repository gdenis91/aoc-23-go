package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	var sum int
	for _, l := range lines {
		nums := make([]int, 0)
		for _, v := range strings.Fields(l) {
			nums = append(nums, util.MustAtoi(v))
		}
		sum += extrapolateValue(nums)
	}
	fmt.Println("Sum:", sum)
}

// 0 3 6 9 12 15
func extrapolateValue(values []int) int {
	allDiffs := make([][]int, 0)
	curDiffs := slices.Clone(values)
	allDiffs = append(allDiffs, curDiffs)
	for {
		diffs := make([]int, 0)
		allZero := true
		for i := 0; i < len(curDiffs)-1; i++ {
			diff := curDiffs[i] - curDiffs[i+1]
			diff = diff * -1
			if diff != 0 {
				allZero = false
			}
			diffs = append(diffs, int(diff))
		}
		curDiffs = slices.Clone(diffs)
		allDiffs = append(allDiffs, diffs)
		if allZero {
			break
		}
	}

	slices.Reverse(allDiffs)
	allDiffs[0] = append(allDiffs[0], 0)
	for i := 0; i < len(allDiffs)-1; i++ {
		exVal := allDiffs[i][len(allDiffs[i])-1] + allDiffs[i+1][len(allDiffs[i+1])-1]
		fmt.Println("ExVal:", exVal)
		allDiffs[i+1] = append(allDiffs[i+1], exVal)
	}

	for _, diffs := range allDiffs {
		fmt.Println("Diffs:", diffs)
	}
	return allDiffs[len(allDiffs)-1][len(allDiffs[len(allDiffs)-1])-1]
}

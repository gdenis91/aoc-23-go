package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

func main() {
	time, dist := mustParseInput(input)
	fmt.Println("Time:", time)
	fmt.Println("Distance:", dist)
	winningDistances := findWinningDistances(time, dist)
	fmt.Println("Ways to win:", len(winningDistances))
}

func findWinningDistances(t int, toBeat int) []int {
	distances := make([]int, 0, t)
	for btnPressTime := 0; btnPressTime <= t; btnPressTime++ {
		remainingTime := t - btnPressTime
		dist := remainingTime * btnPressTime
		if dist > toBeat {
			distances = append(distances, dist)
		}
	}
	return distances
}

func mustParseInput(v string) (int, int) {
	lines := strings.Split(v, "\n")
	times := make([]string, 0)
	for i, t := range strings.Fields(lines[0]) {
		if i == 0 {
			continue
		}
		times = append(times, t)
	}

	dist := make([]string, 0)
	for i, d := range strings.Fields(lines[1]) {
		if i == 0 {
			continue
		}
		dist = append(dist, d)
	}
	return util.MustAtoi(strings.Join(times, "")), util.MustAtoi(strings.Join(dist, ""))
}

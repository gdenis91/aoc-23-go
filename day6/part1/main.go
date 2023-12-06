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
	times, dist := mustParseInput(input)
	fmt.Println("Times:", times)
	fmt.Println("Distances:", dist)

	product := 1
	for i := 0; i < len(times); i++ {
		distances := findWinningDistances(times[i], dist[i])
		fmt.Println("Winning Distances:", distances)
		product *= len(distances)
	}
	fmt.Println("Product:", product)
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

func mustParseInput(v string) ([]int, []int) {
	lines := strings.Split(v, "\n")
	times := make([]int, 0)
	for i, t := range strings.Fields(lines[0]) {
		if i == 0 {
			continue
		}
		times = append(times, util.MustAtoi(t))
	}

	dist := make([]int, 0)
	for i, d := range strings.Fields(lines[1]) {
		if i == 0 {
			continue
		}
		dist = append(dist, util.MustAtoi(d))
	}
	return times, dist
}

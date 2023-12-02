package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	var powerSum int
	for _, l := range lines {
		g := mustParseGame(l)
		powerSum += findMax(g.red) * findMax(g.green) * findMax(g.blue)
	}
	fmt.Println("Total:", powerSum)
}

func findMax(values []int) int {
	var max int
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

type game struct {
	id    int
	blue  []int
	red   []int
	green []int
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func mustParseGame(line string) game {
	parts := strings.Split(line, ": ")

	g := game{
		id: mustParseGameID(parts[0]),
	}

	reveals := strings.Split(parts[1], "; ")
	for _, r := range reveals {
		r, gr, b := mustParseReveal(r)
		g.red = append(g.red, r)
		g.green = append(g.green, gr)
		g.blue = append(g.blue, b)
	}

	return g
}

// Game 1
func mustParseGameID(v string) int {
	parts := strings.Split(v, " ")
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Errorf("parse game id: %w", err))
	}
	return id
}

// 1 red, 2 green, 6 blue
func mustParseReveal(v string) (r int, g int, b int) {
	colors := strings.Split(v, ", ")
	for _, c := range colors {
		color := strings.Split(c, " ")
		var err error
		switch color[1] {
		case "red":
			r, err = strconv.Atoi(color[0])
		case "green":
			g, err = strconv.Atoi(color[0])
		case "blue":
			b, err = strconv.Atoi(color[0])
		}
		if err != nil {
			panic(fmt.Errorf("parse %s count: %w", color[1], err))
		}
	}
	return r, g, b
}

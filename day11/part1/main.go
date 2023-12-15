package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	mustParseInput(input)
}

type Coord struct {
	X int
	Y int
}

func mustParseInput(input string) {
	lines := strings.Split(input, "\n")
	emptyRows := make(map[int]struct{}, 0)
	emptyCols := make(map[int]struct{}, 0)

	for y := 0; y < len(lines); y++ {
		emptyCols[y] = struct{}{}
		for x := 0; x < len(lines[y]); x++ {
			emptyRows[x] = struct{}{}
		}
	}

	for y, l := range lines {
		for x, v := range l {
			if v == '.' {
				continue
			}
			delete(emptyRows, y)
			delete(emptyCols, x)
		}
	}
	fmt.Println("Empty rows:", emptyRows)
	fmt.Println("Empty cols:", emptyCols)

	expandedWidth := make([][]rune, 0)
	for _, l := range lines {
		row := make([]rune, 0)
		for x, v := range l {
			if _, ok := emptyCols[x]; ok {
				row = append(row, '.')
			}
			row = append(row, v)
		}
		expandedWidth = append(expandedWidth, row)
	}
	expanded := make([][]rune, 0)
	for y, l := range expandedWidth {
		row := make([]rune, 0)
		if _, ok := emptyRows[y]; ok {
			extraRow := strings.Repeat(".", len(l))
			expanded = append(expanded, []rune(extraRow))
		}
		row = append(row, l...)
		expanded = append(expanded, row)
	}

	fmt.Println("Expanded space:")
	galaxyCoords := make([]Coord, 0)
	for y, l := range expanded {
		for x, v := range l {
			if v == '#' {
				galaxyCoords = append(galaxyCoords, Coord{x, y})
			}
			fmt.Print(string(v))
		}
		fmt.Print("\n")
	}

	coordsChecked := 0
	sum := 0
	for i, galaxy := range galaxyCoords {
		for _, other := range galaxyCoords[i+1:] {
			coordsChecked++
			x := galaxy.X - other.X
			y := galaxy.Y - other.Y
			dist := int(math.Abs(float64(x)) + math.Abs(float64(y)))
			fmt.Println(galaxy, other, "Dist:", dist)
			sum += dist
		}
	}

	fmt.Println("Checked pairs:", coordsChecked)
	fmt.Println("Shortest path sum:", sum)
}

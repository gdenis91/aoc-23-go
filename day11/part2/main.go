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

	galaxyCoords := make([]Coord, 0)
	for y, l := range lines {
		for x, v := range l {
			if v == '.' {
				continue
			}
			galaxyCoords = append(galaxyCoords, Coord{x, y})
			delete(emptyRows, y)
			delete(emptyCols, x)
		}
	}
	fmt.Println("Empty rows:", emptyRows)
	fmt.Println("Empty cols:", emptyCols)

	coordsChecked := 0
	sum := 0
	for i, galaxy := range galaxyCoords {
		for _, other := range galaxyCoords[i+1:] {
			coordsChecked++

			// Find the number of empty rows between X
			// Multiply scale factor by number of empties
			// Add result to the distance between coord
			fmt.Println("Scale X")
			xScale := 0
			for x := range emptyCols {
				if isBetween(galaxy.X, other.X, x) {
					fmt.Printf("%d is between %d and %d\n", x, galaxy.X, other.X)
					xScale++
				}
			}
			fmt.Println("X Scale:", xScale)

			// Find the number of empty rows between Y
			// Multiply scale factor by number of empties
			// Add result to the distance between coord
			fmt.Println("Scale Y")
			yScale := 0
			for y := range emptyRows {
				if isBetween(galaxy.Y, other.Y, y) {
					fmt.Printf("%d is between %d and %d\n", y, galaxy.Y, other.Y)
					yScale++
				}
			}
			fmt.Println("Y Scale:", yScale)

			const scaleFactor = 1000000
			x := galaxy.X - other.X
			x = int(math.Abs(float64(x)))
			fmt.Println("Additional cols:", scaleFactor*xScale)
			x -= xScale
			x += scaleFactor * xScale
			y := galaxy.Y - other.Y
			y = int(math.Abs(float64(y)))
			fmt.Println("Additional rows:", scaleFactor*yScale)
			y -= yScale
			y += scaleFactor * yScale
			dist := x + y
			fmt.Println(galaxy, other, "Dist:", dist)
			sum += dist
		}
	}

	fmt.Println("Checked pairs:", coordsChecked)
	fmt.Println("Shortest path sum:", sum)
}

func isBetween(p1 int, p2 int, v int) bool {
	if p1 > p2 {
		t := p1
		p1 = p2
		p2 = t
	}
	// p1 < p2
	return p1 < v && v < p2
}

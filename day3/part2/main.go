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

type coord struct {
	x int
	y int
}

func (c coord) borders(v coord) bool {
	return c == coord{v.x, v.y + 1} ||
		c == coord{v.x + 1, v.y + 1} ||
		c == coord{v.x + 1, v.y} ||
		c == coord{v.x + 1, v.y - 1} ||
		c == coord{v.x, v.y - 1} ||
		c == coord{v.x - 1, v.y - 1} ||
		c == coord{v.x - 1, v.y} ||
		c == coord{v.x - 1, v.y + 1}
}

type partNumber struct {
	v     int
	len   int
	start coord
}

func main() {
	g := parseGrid(input)
	gearCoords := make([]coord, 0)
	partNumbers := make([]partNumber, 0)
	for y := range g {
		for x := 0; x < len(g[y]); x++ {
			// If anything but a number, carry on
			// Otherwise we're at a number and need to build the full number
			// by scanning ahead in the row. Iteration will continue from the next
			// non-number position
			v := g[y][x]
			if v == '*' {
				gearCoords = append(gearCoords, coord{x, y})
			}
			if !unicode.IsDigit(v) {
				continue
			}

			num, endX, isPartNum := mustScanNumber(g, x, y)
			fmt.Printf("Found number %d; isPartNum=%t\n", num, isPartNum)
			if isPartNum {
				partNumbers = append(partNumbers, partNumber{
					v:     num,
					len:   endX - x + 1,
					start: coord{x, y},
				})
			}
			x = endX
		}
	}
	sum := 0
	for _, c := range gearCoords {
		fmt.Println("Checking gear:", c)
		var partCount int
		gearRatio := 1
		for _, pn := range partNumbers {
			for i := 0; i < pn.len; i++ {
				if c.borders(coord{pn.start.x + i, pn.start.y}) {
					fmt.Println("Borders pn:", pn)
					gearRatio = gearRatio * pn.v
					partCount++
					break
				}
			}
		}
		if partCount == 2 {
			fmt.Println("Adding ratio:", gearRatio)
			sum += gearRatio
		}
	}

	fmt.Println("Sum:", sum)
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(lines))
	for _, l := range lines {
		row := make([]rune, 0, len(l))
		for _, r := range l {
			row = append(row, r)
		}
		grid = append(grid, row)
	}
	return grid
}

func mustScanNumber(g [][]rune, x int, y int) (num int, endX int, isPartNum bool) {
	digits := make([]rune, 0)
	row := g[y]
	for ; x < len(row); x++ {
		if !unicode.IsDigit(row[x]) {
			break
		}
		if checkForSymbol(g, x, y) {
			isPartNum = true
		}
		digits = append(digits, row[x])
		endX = x
	}
	num, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(fmt.Errorf("parse digits: %w", err))
	}

	return num, endX, isPartNum
}

func checkForSymbol(g [][]rune, x int, y int) bool {
	check := func(x int, y int) bool {
		if x < 0 || y < 0 || y >= len(g) || x >= len(g[y]) {
			return false
		}
		if g[y][x] == '.' || unicode.IsDigit(g[y][x]) {
			return false
		}
		return true
	}
	return check(x, y+1) ||
		check(x+1, y+1) ||
		check(x+1, y) ||
		check(x+1, y-1) ||
		check(x, y-1) ||
		check(x-1, y-1) ||
		check(x-1, y) ||
		check(x-1, y+1)
}

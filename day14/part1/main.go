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
	ctlPanel := mustParseInput(input)
	fmt.Println("Panel:")
	fmt.Println(ctlPanel.raw)
	fmt.Println("Load:", ctlPanel.load())
	ctlPanel.tilt()
	fmt.Println()
	fmt.Println("Tilted Panel:")
	for y := 0; y < len(ctlPanel.table); y++ {
		row := make([]string, 0)
		for x := 0; x < len(ctlPanel.table[y]); x++ {
			if _, ok := ctlPanel.empty[coord{x, y}]; ok {
				row = append(row, ".")
			} else if _, ok := ctlPanel.roundedRocks[coord{x, y}]; ok {
				row = append(row, "O")
			} else if _, ok := ctlPanel.cubedRocks[coord{x, y}]; ok {
				row = append(row, "#")
			}
		}
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("Load:", ctlPanel.load())
}

type coord struct {
	x int
	y int
}

func (c coord) north() coord {
	return coord{c.x, c.y - 1}
}

type controlPanel struct {
	raw          string
	table        [][]rune
	roundedRocks map[coord]struct{}
	cubedRocks   map[coord]struct{}
	empty        map[coord]struct{}
}

func (c controlPanel) load() int {
	load := 0
	for pos := range c.roundedRocks {
		load += int(math.Abs(float64(pos.y - len(c.table))))
	}

	return load
}

func (c *controlPanel) tilt() {
	for c.step() > 0 {
	}
}

func (c *controlPanel) step() int {
	numMoved := 0
	for y := 0; y < len(c.table); y++ {
		// At top of panel, can't go more
		if y == 0 {
			continue
		}
		for x := 0; x < len(c.table[y]); x++ {
			if c.tryMoveNorth(coord{x, y}) {
				fmt.Println("Moved rock north")
				numMoved++
			}
		}
	}
	return numMoved
}

func (c *controlPanel) tryMoveNorth(pos coord) bool {
	if _, ok := c.cubedRocks[pos]; ok {
		return false
	} else if _, ok := c.empty[pos]; ok {
		return false
	}
	// If we are at a round rock O we check if we can move north
	n := pos.north()

	// If a cubed rock or another round rock is to the north we can't move
	if _, ok := c.cubedRocks[n]; ok {
		return false
	} else if _, ok := c.roundedRocks[n]; ok {
		return false
	}

	// We can move north!
	c.roundedRocks[n] = struct{}{}
	delete(c.roundedRocks, pos)
	c.empty[pos] = struct{}{}
	delete(c.empty, n)

	return true
}

// O....#....
// O.OO#....#
// .....##...
// OO.#O....O
// .O.....O#.
// O.#..O.#.#
// ..O..#O..O
// .......O..
// #....###..
// #OO..#....
func mustParseInput(input string) controlPanel {
	ctlPanel := controlPanel{
		raw:          input,
		roundedRocks: make(map[coord]struct{}),
		cubedRocks:   make(map[coord]struct{}),
		empty:        make(map[coord]struct{}),
	}
	for y, l := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for x, v := range l {
			if v == 'O' {
				ctlPanel.roundedRocks[coord{x, y}] = struct{}{}
			} else if v == '#' {
				ctlPanel.cubedRocks[coord{x, y}] = struct{}{}
			} else {
				ctlPanel.empty[coord{x, y}] = struct{}{}
			}
			row = append(row, v)
		}
		ctlPanel.table = append(ctlPanel.table, row)
	}
	return ctlPanel
}

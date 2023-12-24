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
	fmt.Println("Start Panel:")
	fmt.Println(ctlPanel.raw)
	fmt.Println("Load:", ctlPanel.load())

	seenPanels := make(map[string]int)
	numTilts := 1000000000 * 4
	for i := 0; i < numTilts; i++ {
		startHash := ctlPanel.hash(dir(i))
		// If we've seen this hash before just follow the loop for the rest of the iterations
		if _, seen := seenPanels[startHash]; seen {
			fmt.Println("Hit a previously seen state, in loop")
			i = numTilts - (numTilts-i)%(i-seenPanels[startHash])
			// ctlPanel = mustParseInput(curHash)
			// break
		}
		ctlPanel.tilt(dir(i % 4))
		// endHash := ctlPanel.hash(dir(i % 4))
		seenPanels[startHash] = i

		// if i%1 == 0 {
		// 	fmt.Println("Iteration:", i)
		// 	fmt.Println(endHash)
		// }

		// if i%1000 == 0 {
		// 	fmt.Println(i)
		// }
	}

	fmt.Println()
	fmt.Println("Final Panel:")
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

type dir int

const (
	North = iota
	West
	South
	East
)

func (c coord) neighbor(d dir) coord {
	switch d {
	case North:
		return coord{c.x, c.y - 1}
	case East:
		return coord{c.x + 1, c.y}
	case South:
		return coord{c.x, c.y + 1}
	case West:
		return coord{c.x - 1, c.y}
	}
	return c
}

type controlPanel struct {
	raw          string
	table        [][]rune
	roundedRocks map[coord]struct{}
	cubedRocks   map[coord]struct{}
	empty        map[coord]struct{}
}

func (c *controlPanel) hash(d dir) string {
	rows := make([]string, 0, 1+len(c.table)*len(c.table[0]))
	// rows = append(rows, fmt.Sprint(d))
	for y := 0; y < len(c.table); y++ {
		row := make([]string, 0, len(c.table[y]))
		for x := 0; x < len(c.table[y]); x++ {
			if _, ok := c.empty[coord{x, y}]; ok {
				row = append(row, ".")
			} else if _, ok := c.roundedRocks[coord{x, y}]; ok {
				row = append(row, "O")
			} else if _, ok := c.cubedRocks[coord{x, y}]; ok {
				row = append(row, "#")
			}
		}
		rows = append(rows, strings.Join(row, ""))
	}
	return strings.Join(rows, "\n")
}

func (c *controlPanel) load() int {
	load := 0
	for pos := range c.roundedRocks {
		load += int(math.Abs(float64(pos.y - len(c.table))))
	}

	return load
}

func (c *controlPanel) tilt(d dir) {
	for c.step(d) > 0 {
	}
}

func (c *controlPanel) step(d dir) int {
	numMoved := 0
	for y := 0; y < len(c.table); y++ {
		// At top of panel, can't go more
		if y == 0 && d == North {
			continue
		} else if y == len(c.table)-1 && d == South {
			continue
		}
		for x := 0; x < len(c.table[y]); x++ {
			// At left of panel, can't go more
			if x == 0 && d == West {
				continue
			} else if x == len(c.table[y])-1 && d == East {
				continue
			}

			if c.tryMove(coord{x, y}, d) {
				// fmt.Println("Moved rock:", d)
				numMoved++
			}
		}
	}
	return numMoved
}

func (c *controlPanel) tryMove(pos coord, d dir) bool {
	if _, ok := c.cubedRocks[pos]; ok {
		return false
	} else if _, ok := c.empty[pos]; ok {
		return false
	}
	// If we are at a round rock O we check if we can move in dir
	n := pos.neighbor(d)

	// If a cubed rock or another round rock is to the dir we can't move
	if _, ok := c.cubedRocks[n]; ok {
		return false
	} else if _, ok := c.roundedRocks[n]; ok {
		return false
	}

	// We can move in the dir!
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

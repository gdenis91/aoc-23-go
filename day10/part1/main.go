package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pipes, start := mustParseInput(input)
	// fmt.Println("Pipes:")
	// for _, row := range pipes {
	// 	for _, p := range row {
	// 		fmt.Printf("%s", string(p.Shape))
	// 	}
	// 	fmt.Print("\n")
	// }
	pipeLen := findLen(start, pipes)
	fmt.Println("Loop length:", pipeLen)
	fmt.Println("Farthest point:", pipeLen/2)
}

func neighborCoords(c Coord) []Coord {
	coords := make([]Coord, 0, 4)
	coords = append(coords, Coord{c.X, c.Y + 1})
	coords = append(coords, Coord{c.X, c.Y - 1})
	coords = append(coords, Coord{c.X + 1, c.Y})
	coords = append(coords, Coord{c.X - 1, c.Y})
	return coords
}

func findLen(start Pipe, pipes [][]Pipe) int {
	currentPipe := start
	var loopLen int
	visited := make(map[Coord]struct{})

	getPipe := func(coord Coord) (Pipe, bool) {
		if coord.X < 0 || coord.Y < 0 || coord.Y >= len(pipes) || coord.X >= len(pipes[coord.Y]) {
			return Pipe{}, false
		}
		return pipes[coord.Y][coord.X], true
	}

	for {
		fmt.Println("Current pipe:", currentPipe)
		var nextPipe *Pipe
		for _, c := range neighborCoords(currentPipe.Coord) {
			p, ok := getPipe(c)
			if !ok {
				continue
			}
			fmt.Println("Checking pipe:", p)
			if _, visited := visited[c]; !visited && p.AllowedIncomingConnection(currentPipe) {
				nextPipe = &p
				break
			}
		}
		loopLen++
		if nextPipe == nil {
			fmt.Println("No new neigbors")
			break
		}
		fmt.Println("Next pipe:", nextPipe)
		visited[currentPipe.Coord] = struct{}{}
		currentPipe = *nextPipe
	}

	fmt.Println("Visited:", visited)
	fmt.Println("Walked Pipes:")
	for y, row := range pipes {
		for x, p := range row {
			if _, ok := visited[Coord{x, y}]; ok {
				fmt.Print("👟")
				continue
			}
			fmt.Printf("%s", string(p.Shape))
		}
		fmt.Print("\n")
	}

	return loopLen
}

func mustParseInput(input string) ([][]Pipe, Pipe) {
	lines := strings.Split(input, "\n")
	pipes := make([][]Pipe, 0, len(lines))
	var startingPipe Pipe
	for y, l := range lines {
		row := make([]Pipe, 0, len(l))
		for x, p := range l {
			pipe := Pipe{
				Shape: p,
				Coord: Coord{
					X: x,
					Y: y,
				},
			}
			if p == 'S' {
				startingPipe = pipe
			}
			row = append(row, pipe)
		}
		pipes = append(pipes, row)
	}
	return pipes, startingPipe
}

type Coord struct {
	X int
	Y int
}

func (c Coord) Up() Coord {
	return Coord{
		X: c.X,
		Y: c.Y - 1,
	}
}

func (c Coord) Down() Coord {
	return Coord{
		X: c.X,
		Y: c.Y + 1,
	}
}

func (c Coord) Left() Coord {
	return Coord{
		X: c.X - 1,
		Y: c.Y,
	}
}

func (c Coord) Right() Coord {
	return Coord{
		X: c.X + 1,
		Y: c.Y,
	}
}

type Pipe struct {
	Shape rune
	Coord Coord
}

func (p Pipe) String() string {
	return fmt.Sprintf("%s, (%d,%d)", string(p.Shape), p.Coord.X, p.Coord.Y)
}

func (p Pipe) AllowedIncomingConnection(incomingPipe Pipe) bool {
	// fmt.Println("Incoming pipe:", incomingPipe)
	// fmt.Println("Next pipe:", p)
	switch p.Shape {
	case 'S':
		switch incomingPipe.Coord {
		case p.Coord.Up():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'F'
		case p.Coord.Down():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'L'
		case p.Coord.Left():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'F'
		case p.Coord.Right():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'L'
		}
	case '-':
		switch incomingPipe.Coord {
		case p.Coord.Left():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'F' || incomingPipe.Shape == 'S'
		case p.Coord.Right():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'S'
		}
	case '|':
		switch incomingPipe.Coord {
		case p.Coord.Up():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'F' || incomingPipe.Shape == 'S'
		case p.Coord.Down():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'S'
		}
	case '7':
		switch incomingPipe.Coord {
		case p.Coord.Left():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == 'F' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'S'
		case p.Coord.Down():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'S'
		}
	case 'F': // Check pipe F (106,19)
		switch incomingPipe.Coord {
		case p.Coord.Right(): // Coming from Right (107, 19)
			return incomingPipe.Shape == '-' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'S'
		case p.Coord.Down(): // Coming from Down (106, 18)
			return incomingPipe.Shape == '|' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'J' || incomingPipe.Shape == 'S'
		}
	case 'L':
		switch incomingPipe.Coord {
		case p.Coord.Right():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == 'J' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'S'
		case p.Coord.Up():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'F' || incomingPipe.Shape == 'S'
		}
	case 'J':
		switch incomingPipe.Coord {
		case p.Coord.Left():
			return incomingPipe.Shape == '-' || incomingPipe.Shape == 'L' || incomingPipe.Shape == 'F' || incomingPipe.Shape == 'S'
		case p.Coord.Up():
			return incomingPipe.Shape == '|' || incomingPipe.Shape == 'F' || incomingPipe.Shape == '7' || incomingPipe.Shape == 'S'
		}
	}
	return false
}

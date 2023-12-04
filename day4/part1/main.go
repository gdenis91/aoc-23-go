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
	var points int
	for i, l := range lines {
		c := mustParseCard(l)
		fmt.Printf("Card %d: %v\n", i+1, c)
		points += c.points()
	}
	fmt.Println("Points:", points)
}

type card struct {
	winningNumbers map[int]struct{}
	numbers        []int
}

func (c card) points() int {
	var points int
	for _, n := range c.numbers {
		_, ok := c.winningNumbers[n]
		if !ok {
			continue
		} else if points == 0 {
			points += 1
		} else {
			points = points * 2
		}
	}
	return points
}

func mustParseCard(line string) card {
	numberParts := strings.Split(strings.Split(line, ": ")[1], "|")
	c := card{
		winningNumbers: make(map[int]struct{}),
	}
	for _, v := range strings.Fields(numberParts[0]) {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("parse winning number %s: %w", v, err))
		}
		c.winningNumbers[n] = struct{}{}
	}
	for _, v := range strings.Fields(numberParts[1]) {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("parse number %s: %w", v, err))
		}
		c.numbers = append(c.numbers, n)
	}
	return c
}

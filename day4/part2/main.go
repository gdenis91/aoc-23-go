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
	copyCounts := make(map[int]int)
	var cards int
	for i, l := range lines {
		c := mustParseCard(l)
		fmt.Printf("Card %d: %v\n", i, c)
		for j := 1; j <= c.matches(); j++ {
			copyCounts[i+j]++
		}
		cards++
		copyCount := copyCounts[i]
		if copyCount > 0 {
			fmt.Printf("Scratching %d copies\n", copyCount)
		}
		for j := 0; j < copyCount; j++ {
			cards++
			for k := 1; k <= c.matches(); k++ {
				copyCounts[i+k]++
			}
		}
	}
	fmt.Println("Cards:", cards)
}

type card struct {
	winningNumbers map[int]struct{}
	numbers        []int
}

func (c card) matches() int {
	var matches int
	for _, n := range c.numbers {
		_, ok := c.winningNumbers[n]
		if !ok {
			continue
		}
		matches++
	}
	return matches
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

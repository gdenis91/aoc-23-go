package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lr, nodes := mustParseInput(input)
	lrPos := 0
	curNode := nodes["AAA"]
	var steps int
	for {
		if curNode.value == "ZZZ" {
			break
		}

		if lr[lrPos] == 'L' {
			curNode = nodes[curNode.left]
		} else {
			curNode = nodes[curNode.right]
		}
		steps++
		lrPos++
		if lrPos >= len(lr) {
			lrPos = 0
		}
	}
	fmt.Println("Steps:", steps)
}

func mustParseInput(input string) (lrInput []rune, nodes map[string]node) {
	fmt.Println("Input:")
	lines := strings.Split(input, "\n")
	lrInput = []rune(lines[0])
	fmt.Println(lines[0])

	// AAA = (BBB, CCC)
	nodes = make(map[string]node, len(lines[2:]))
	for _, r := range lines[2:] {
		parts := strings.FieldsFunc(r, func(r rune) bool {
			return r == ' ' || r == '=' || r == '(' || r == ',' || r == ')'
		})
		nodes[parts[0]] = node{
			value: parts[0],
			left:  parts[1],
			right: parts[2],
		}
		fmt.Println(nodes[parts[0]])
	}
	return lrInput, nodes
}

type node struct {
	value string
	left  string
	right string
}

package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lr, navigators := mustParseInput(input)
	fmt.Println("Num starting nodes:", len(navigators))
	cycleLens := make([]int, 0)
	for _, n := range navigators {
		cycleLens = append(cycleLens, n.computeCycleLen(lr))
	}
	lcm := LCM(cycleLens[0], cycleLens[1], cycleLens[2:]...)
	fmt.Println("lcm:", lcm)
}

func mustParseInput(input string) (lrInput []rune, navigators []*navigator) {
	fmt.Println("Input:")
	lines := strings.Split(input, "\n")
	lrInput = []rune(lines[0])
	fmt.Println(lines[0])

	// AAA = (BBB, CCC)
	nodes := make(map[string]node, len(lines[2:]))
	navigators = make([]*navigator, 0)
	for _, r := range lines[2:] {
		parts := strings.FieldsFunc(r, func(r rune) bool {
			return r == ' ' || r == '=' || r == '(' || r == ',' || r == ')'
		})
		nodes[parts[0]] = node{
			value: parts[0],
			left:  parts[1],
			right: parts[2],
		}
		if strings.HasSuffix(nodes[parts[0]].value, "A") {
			nav := &navigator{
				curNode: nodes[parts[0]],
				nodes:   nodes,
			}
			navigators = append(navigators, nav)
		}
		fmt.Println(nodes[parts[0]])
	}
	return lrInput, navigators
}

type node struct {
	value string
	left  string
	right string
}

type navigator struct {
	curNode node
	nodes   map[string]node
}

func (n *navigator) computeCycleLen(lr []rune) int {
	lrPos := 0
	var steps int
	curNode := n.curNode
	for {
		if lr[lrPos] == 'L' {
			curNode = n.nodes[curNode.left]
		} else {
			curNode = n.nodes[curNode.right]
		}

		steps++
		lrPos++
		if lrPos >= len(lr) {
			lrPos = 0
		}

		if strings.HasSuffix(curNode.value, "Z") {
			fmt.Printf("Found ending node after %d steps: %s\n", steps, curNode.value)
			return steps
		}
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

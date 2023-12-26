package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

type opertion int

const (
	Equals opertion = iota
	Dash
)

type lens struct {
	focalLen int
	label    string
}

func main() {
	steps := strings.Split(input, ",")
	boxes := make(map[int][]lens, 256)
	for _, v := range steps {
		fmt.Println(boxes)
		var parts []string
		var op opertion
		if strings.Contains(v, "=") {
			parts = strings.Split(v, "=")
			op = Equals
		} else if strings.Contains(v, "-") {
			parts = strings.Split(v, "-")
			op = Dash
		}
		label := parts[0]
		boxID := holidyHash(label)
		if op == Dash {
			for _, l := range boxes[boxID] {
				if label == l.label {
					fmt.Println("Match!")
					boxes[boxID] = slices.DeleteFunc(boxes[boxID], func(v lens) bool {
						if v.label == l.label {
							fmt.Println("Deleting:", l.label)
						}
						return label == v.label
					})
					break
				}
			}
			continue
		}
		focalLen := util.MustAtoi(parts[1])
		replaced := false
		for i, l := range boxes[boxID] {
			if label == l.label {
				boxes[boxID][i] = lens{focalLen, label}
				replaced = true
				break
			}
		}
		if replaced {
			continue
		}
		boxes[boxID] = append(boxes[boxID], lens{focalLen, label})

	}
	fmt.Println(boxes)

	result := 0
	for i := 0; i < 256; i++ {
		box := boxes[i]
		for j, l := range box {
			v := (i + 1) * (j + 1) * l.focalLen
			fmt.Printf("Box %d,%s:\n", i, l.label)
			fmt.Printf("%d * %d * %d = %d\n", (i + 1), (j + 1), l.focalLen, v)
			result += v
		}
	}
	fmt.Println(result)
}

func holidyHash(v string) int {
	cv := 0
	for _, r := range v {
		cv += int(r)
		cv *= 17
		cv = int(cv) % 256
	}
	return cv
}

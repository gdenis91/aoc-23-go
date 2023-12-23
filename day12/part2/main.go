package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

const Operational = '.'
const Damaged = '#'
const Unknown = '?'

func main() {
	rows := mustParseInput(input)
	sum := 0
	for i, r := range rows {
		// fmt.Println(r.raw)
		arrangements := r.numOptions(0, r.arrangements)
		// fmt.Println("Arrangements:", arrangements)
		fmt.Println(i+1, arrangements)
		sum += arrangements
	}
	fmt.Println("Sum:", sum)
}

func (r row) numOptions(pos int, contiguousGroups []int) int {
	// fmt.Printf("numOptions: pos=%d, contiguousGroups=%v\n", pos, contiguousGroups)
	// If we just left behind a broken spring this is not a valid position
	if pos >= len(r.conditionRecord) {
		return 0
	}

	if pos > 0 && r.conditionRecord[pos-1] == '#' {
		return 0
	}

	totalOptions := 0
	for i := pos; i < len(r.conditionRecord); i++ {
		// If we can't insert the current contiguous group at the current pos
		// advance by one and continue to try inserting until we find a success
		if len(contiguousGroups) == 0 {
			break
		}
		if i > 0 && r.conditionRecord[i-1] == '#' {
			break
		}
		if !r.canInsert(i, contiguousGroups[0]) {
			continue
		}

		// If this is the last contiguous group see if we've covered all the broken springs
		if len(contiguousGroups) == 1 {
			// fmt.Println("Remaining:", r.remainingConditionRecord(i+contiguousGroups[0]))
			hasBrokenSprings := false
			for _, v := range r.conditionRecord[i+contiguousGroups[0]:] {
				if v == '#' {
					hasBrokenSprings = true
					break
				}
			}
			if hasBrokenSprings {
				// fmt.Println("Missing required broken spring at:", pos)
			} else {
				// fmt.Println("No more required broken springs, valid")
				totalOptions += 1
			}
		}

		// If we can insert move the pos over by two and move to the next contiguous group
		var numOptions int
		lookup := fmt.Sprint(i+contiguousGroups[0]+1, contiguousGroups[1:])
		if v, ok := r.prevResults[lookup]; ok {
			numOptions = v
		} else {
			numOptions = r.numOptions(i+contiguousGroups[0]+1, contiguousGroups[1:])
			r.prevResults[lookup] = numOptions
		}
		totalOptions += numOptions
	}
	return totalOptions
}

func (r row) canInsert(pos int, size int) bool {
	if pos+size > len(r.conditionRecord) {
		return false
	}
	for i := 0; i < size; i++ {
		if r.conditionRecord[pos+i] == Operational {
			// fmt.Println("Can't insert at opertional pos")
			return false
		}
	}
	// fmt.Printf("canInsert: pos=%d size=%d\n", pos, size)
	return true
}

type row struct {
	raw             string
	conditionRecord []rune
	arrangements    []int
	// Store a map of previous runs of the condition record and number of found options
	// If we have a match we can short circuit the checks and just return this
	prevResults map[string]int
}

// ???.### 1,1,3
func mustParseInput(input string) []row {
	rows := make([]row, 0)

	for _, l := range strings.Split(input, "\n") {
		fields := strings.Fields(l)
		arrangements := make([]int, 0)
		for _, v := range strings.Split(fields[1], ",") {
			arrangements = append(arrangements, util.MustAtoi(v))
		}
		repeatedArrangements := make([]int, 0, len(arrangements)*5)
		repeatedRecords := make([]string, 0, 5)
		for i := 0; i < 5; i++ {
			repeatedArrangements = append(repeatedArrangements, arrangements...)
			repeatedRecords = append(repeatedRecords, fields[0])
		}
		rows = append(rows, row{
			raw:             l,
			conditionRecord: []rune(strings.Join(repeatedRecords, "?")),
			arrangements:    repeatedArrangements,
			prevResults:     make(map[string]int),
		})
	}
	return rows
}

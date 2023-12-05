package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/gdenis91/aoc-23-go/util"
)

//go:embed input.txt
var input string

func main() {
	seeds, plantingMaps := mustParseInput(input)
	fmt.Println("Seeds:", seeds)
	for _, pm := range plantingMaps {
		fmt.Println("Planting map:", pm)
	}

	lowestLocation := findLocation(seeds[0], plantingMaps)
	for _, s := range seeds[1:] {
		l := findLocation(s, plantingMaps)
		if l < lowestLocation {
			lowestLocation = l
		}
	}
	fmt.Println("Location:", lowestLocation)
}

func findLocation(seed int, maps map[string]plantingMap) int {
	currentMap := maps["seed"]
	currentLookup := seed
	for {
		fmt.Printf("Looking up mapping %s -> %s\n", currentMap.sourceKey, currentMap.destinationKey)
		for _, v := range currentMap.ranges {
			if !v.inSourceRange(currentLookup) {
				continue
			}
			currentLookup = v.getDestination(currentLookup)
			fmt.Println("Found destination:", currentLookup)
			break
		}
		if currentMap.destinationKey == "location" {
			break
		}
		currentMap = maps[currentMap.destinationKey]
	}
	fmt.Println("Seed:", seed)
	fmt.Println("Loc:", currentLookup)
	return currentLookup
}

func mustParseInput(input string) (seeds []int, plantingMaps map[string]plantingMap) {
	lines := strings.Split(input, "\n")
	plantingMaps = make(map[string]plantingMap, 0)
	for i, l := range lines {
		if i == 0 {
			seeds = mustParseSeeds(l)
			continue
		}
		if len(l) == 0 {
			continue
		}
		if unicode.IsLetter([]rune(l)[0]) {
			pm := mustParsePlantingMap(lines[i:])
			plantingMaps[pm.sourceKey] = pm
		}
	}
	return
}

// seeds: 79 14 55 13
func mustParseSeeds(line string) []int {
	fields := strings.Fields(strings.Split(line, ": ")[1])
	seeds := make([]int, 0, len(fields))
	for _, v := range fields {
		seeds = append(seeds, util.MustAtoi(v))
	}
	return seeds
}

// seed-to-soil map:
// 50 98 2
// 52 50 48
func mustParsePlantingMap(lines []string) plantingMap {
	var pm plantingMap
	for i, l := range lines {
		if len(l) == 0 {
			break
		}

		if i == 0 {
			keys := strings.Split(strings.Split(l, " ")[0], "-to-")
			pm.sourceKey = keys[0]
			pm.destinationKey = keys[1]
			continue
		}

		fields := strings.Fields(l)
		pm.ranges = append(pm.ranges, plantingRange{
			destinationStart: util.MustAtoi(fields[0]),
			sourceStart:      util.MustAtoi(fields[1]),
			length:           util.MustAtoi(fields[2]),
		})
	}
	return pm
}

type plantingMap struct {
	sourceKey      string
	destinationKey string
	ranges         []plantingRange
}

type plantingRange struct {
	sourceStart      int
	destinationStart int
	length           int
}

func (r plantingRange) inSourceRange(v int) bool {
	return v > r.sourceStart && v <= r.sourceStart+r.length
}

func (r plantingRange) getDestination(v int) int {
	if !r.inSourceRange(v) {
		return v
	}
	offset := v - r.sourceStart
	return r.destinationStart + offset
}

package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	steps := strings.Split(input, ",")
	result := 0
	for _, v := range steps {
		result += holidyHash(v)
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
